package net

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "fmt"
    "github.com/junglemc/JungleTree/internal/net/protocol"
	"github.com/junglemc/JungleTree/pkg/event"
	"log"
    "net"
    "reflect"
)

func init() {
	event.Register(event.ServerStartupEvent{}, event.ServerStartupListener{})
	event.Register(event.ServerLoadedEvent{}, event.ServerLoadedListener{})
}

type Server struct {
    Address string
    Port    uint16

    privateKey      *rsa.PrivateKey
    privateKeyBytes []byte
    publicKeyBytes  []byte

    OnlineMode           bool
    CompressionThreshold int
    Debug                bool
    Verbose              bool
    MaxOnlinePlayers     int32

    HandshakeHandlers map[reflect.Type]func(c *Client, pkt Packet) error
    StatusHandlers    map[reflect.Type]func(c *Client, pkt Packet) error
    LoginHandlers     map[reflect.Type]func(c *Client, pkt Packet) error
    PlayHandlers      map[reflect.Type]func(c *Client, pkt Packet) error

    DisconnectHandler func(c *Client, reason string)
}

func (s *Server) GenerateKeys() {
    if s.privateKey != nil {
        panic("keys already generated")
    }

    privKey, err := rsa.GenerateKey(rand.Reader, 1024)
    if err != nil {
        panic(err)
    }

    privKey.Precompute()
    if err := privKey.Validate(); err != nil {
        panic(err)
    }
    s.privateKey = privKey
    s.privateKeyBytes = x509.MarshalPKCS1PrivateKey(privKey)
    s.publicKeyBytes, _ = x509.MarshalPKIXPublicKey(privKey.Public())
}

func (s *Server) PrivateKey() *rsa.PrivateKey {
    return s.privateKey
}

func (s *Server) PublicKey() []byte {
    return s.publicKeyBytes
}

func (s *Server) Listen() error {
    if s.OnlineMode {
        s.GenerateKeys()
    }

    listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Address, s.Port))
    if err != nil {
        return err
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("client::listenerAccept: %s\n", err)
            conn.Close()
            continue
        }
        client := &Client{
            Connection: conn,
            Server:     s,
            Protocol:   protocol.Handshake,
        }
        go client.listen()
    }
}

func NewServer(address string, port uint16, onlineMode bool, compressionThreshold int, debug bool, verbose bool, handshake, status, login, play map[reflect.Type]func(c *Client, pkt Packet) error, disconnect func(c *Client, reason string)) *Server {
    return &Server{
        Address: address,
        Port:    port,

        OnlineMode:           onlineMode,
        CompressionThreshold: compressionThreshold,
        Debug:                debug,
        Verbose:              verbose,

        HandshakeHandlers: handshake,
        StatusHandlers:    status,
        LoginHandlers:     login,
        PlayHandlers:      play,
        DisconnectHandler: disconnect,
    }
}
