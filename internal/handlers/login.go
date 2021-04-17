package handlers

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
    "github.com/junglemc/mc"
    "github.com/junglemc/net"
    "github.com/junglemc/net/codec"
    "github.com/junglemc/net/packet"
    "github.com/junglemc/net/session"
    "io/ioutil"
    "log"
    "net/http"
    "reflect"
)

const sessionServerUri = "https://sessionserver.mojang.com/session/minecraft/hasJoined?username=%s&serverId=%s"

// TODO: Refactor function bodies
func loginStart(c *net.Client, p codec.Packet) {
    pkt := p.(packet.ServerboundLoginStartPacket)

    c.Profile = mc.Profile{
        ID:         uuid.UUID{},
        Name:       pkt.Username,
        Properties: nil,
    }

    if c.Server.OnlineMode {
        loginEncryptionRequest(c)
    } else {
        loginCompression(c)

        c.Profile.ID, _ = uuid.NewRandom()
        loginSuccess(c)
    }
}

func loginCompression(c *net.Client) {
    err := c.Send(&packet.ClientboundLoginCompressionPacket{Threshold: c.Server.CompressionThreshold})
    if err != nil {
        log.Println(err)
        return
    }
    c.CompressionEnabled = true
}

func loginEncryptionRequest(c *net.Client) {
    pubBytes, err := x509.MarshalPKIXPublicKey(c.Server.PrivateKey.Public())

    pkt := &packet.ClientboundLoginEncryptionRequest{
        ServerId:    "",
        PublicKey:   pubBytes,
        VerifyToken: c.VerifyToken,
    }

    err = c.Send(pkt)
    if err != nil {
        log.Println(err)
    }
}

func loginEncryptionResponse(c *net.Client, p codec.Packet) {
    pkt := p.(packet.ServerboundLoginEncryptionResponsePacket)

    verifyToken, err := rsa.DecryptPKCS1v15(rand.Reader, c.Server.PrivateKey, pkt.VerifyToken)
    if err != nil {
        log.Println("Verify: " + err.Error())
        return
    }

    if !reflect.DeepEqual(c.VerifyToken, verifyToken) {
        log.Println("VerifyToken mismatch")
        return
    }

    sharedSecret, err := rsa.DecryptPKCS1v15(rand.Reader, c.Server.PrivateKey, pkt.SharedSecret)
    if err != nil {
        log.Println("Shared: " + err.Error())
        return
    }

    c.EnableEncryption(sharedSecret)
    loginVerify(c, sharedSecret)
}

func loginVerify(c *net.Client, sharedSecret []byte) {
    pubBytes, err := x509.MarshalPKIXPublicKey(c.Server.PrivateKey.Public())
    if err != nil {
        log.Println(err)
        return
    }

    authDigest := session.AuthDigest(sharedSecret, pubBytes)

    getUri := fmt.Sprintf(sessionServerUri, c.Profile.Name+"", authDigest)

    response, err := http.Get(getUri)
    if err != nil {
        log.Println(err)
        return
    }

    defer response.Body.Close()

    if response.StatusCode == 204 {
        log.Println("Verify failed")
        c.Send(&packet.ClientboundLoginDisconnectPacket{Reason: mc.Chat{Text: "Invalid session"}})
        c.Disconnect = true
        return
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Println(err)
        return
    }

    c.VerifyToken = nil

    c.Profile = mc.Profile{}
    err = json.Unmarshal(body, &c.Profile)
    if err != nil {
        log.Println(err)
        return
    }
    loginCompression(c)
    loginSuccess(c)
}

func loginSuccess(c *net.Client) {
    response := &packet.ClientboundLoginSuccess{
        Uuid:     c.Profile.ID,
        Username: c.Profile.Name,
    }
    err := c.Send(response)
    if err != nil {
        log.Println(err)
    }
    c.Protocol = codec.ProtocolPlay
}
