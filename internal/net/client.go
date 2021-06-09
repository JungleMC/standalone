package net

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"net"
	"reflect"

	"github.com/junglemc/JungleTree/internal/net/auth"
	. "github.com/junglemc/JungleTree/internal/net/protocol"
	. "github.com/junglemc/JungleTree/internal/pkg/net/packets"
	"github.com/junglemc/JungleTree/pkg/util"
)

type Client struct {
	Connection          net.Conn
	Server              *Server
	Protocol            Protocol
	GameProtocolVersion int32
	disconnect          bool
	disconnectReason    string
	Profile             auth.Profile
	ExpectedVerifyToken []byte
	EncryptionEnabled   bool
	sharedSecret        []byte
	encryptStream       cipher.Stream
	decryptStream       cipher.Stream
	CompressionEnabled  bool
	Gamemode            util.GameMode
}

func (c *Client) listen() {
	c.onClientConnect()

	for {
		if c.disconnect {
			c.Server.DisconnectHandler(c, c.disconnectReason)
			c.onClientDisconnect()
			_ = c.Connection.Close()
			return
		}

		buf := make([]byte, 1024)
		numRead, err := c.Connection.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("client::codec::ReadPacket: %s\n", err)
			}
			c.disconnect = true
			continue
		}

		buf = buf[:numRead]
		if c.EncryptionEnabled {
			c.decryptStream.XORKeyStream(buf, buf)
		}

		reader := bytes.NewBuffer(buf)
		pkt, err := ReadPacket(reader, c.Protocol, c.CompressionEnabled)
		if err != nil {
			if err != io.EOF {
				log.Printf("client::codec::ReadPacket: %s\n", err)
				c.disconnectReason = err.Error()
			}
			c.disconnect = true
			continue
		}

		if pkt == nil {
			continue
		}

		c.onClientPacket(pkt)
	}
}

func (c *Client) Send(pkt Packet) (err error) {
	if c.Server.Debug {
		log.Printf("tx -> %s\n", reflect.TypeOf(pkt).Elem().Name())
		if c.Server.Verbose {
			log.Printf("%+v\n", pkt)
		}
	}

	buf := &bytes.Buffer{}
	WritePacket(buf, reflect.ValueOf(pkt).Elem(), c.Protocol, c.CompressionEnabled, c.Server.CompressionThreshold)

	data := buf.Bytes()
	if c.EncryptionEnabled {
		c.encryptStream.XORKeyStream(data, data)
	}

	_, err = c.Connection.Write(data)
	if err != nil {
		log.Printf("client::Send::Write: %s\n", err)
	}
	return
}

func (c *Client) onClientConnect() {
	c.Protocol = Handshake
	c.ExpectedVerifyToken = make([]byte, 4)
	c.Profile = auth.Profile{}
	_, _ = rand.Read(c.ExpectedVerifyToken)
}

func (c *Client) onClientDisconnect() {
	// Do nothing (yet)
}

func (c *Client) onClientPacket(pkt Packet) {
	if c.Server.Debug {
		log.Printf("rx -> %s\n", reflect.TypeOf(pkt).Name())
		if c.Server.Verbose {
			log.Printf("%+v\n", pkt)
		}
	}

	var find map[reflect.Type]func(c *Client, pkt Packet) error

	s := c.Server

	// TODO: Merge with registry
	switch c.Protocol {
	case Handshake:
		find = s.HandshakeHandlers
		break
	case Status:
		find = s.StatusHandlers
		break
	case Login:
		find = s.LoginHandlers
		break
	case Play:
		find = s.PlayHandlers
		break
	}

	if find == nil {
		log.Printf("Handler registry for %s protocol not defined\n", c.Protocol.String())
		return
	}

	funcCall := find[reflect.TypeOf(pkt)]
	if funcCall == nil {
		log.Printf("Handler for packet \"%s\" not defined\n", reflect.TypeOf(pkt).Name())
		return
	}
	err := funcCall(c, pkt)
	if err != nil {
		c.Disconnect(err.Error())
		return
	}
}

func (c *Client) Disconnect(reason string) {
	c.disconnect = true
	c.disconnectReason = reason
}

func (c *Client) EnableEncryption(sharedSecret []byte) (err error) {
	block, err := aes.NewCipher(sharedSecret)
	if err != nil {
		return
	}

	c.sharedSecret = sharedSecret
	c.encryptStream, err = auth.NewCFB8Encrypter(block, sharedSecret)
	if err != nil {
		return
	}

	c.decryptStream, err = auth.NewCFB8Decrypter(block, sharedSecret)
	if err != nil {
		return
	}

	c.EncryptionEnabled = true
	return
}

func (c *Client) EnableCompression() (err error) {
	err = c.Send(&ClientboundLoginCompressionPacket{Threshold: int32(c.Server.CompressionThreshold)})
	if err != nil {
		return
	}
	c.CompressionEnabled = true
	return
}
