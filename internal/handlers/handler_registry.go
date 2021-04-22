package handlers

import (
    "github.com/junglemc/net"
    "github.com/junglemc/net/codec"
    "github.com/junglemc/net/packet"
    "reflect"
)

func t(i interface{}) reflect.Type {
    return reflect.TypeOf(i)
}

var Handshake = map[reflect.Type]func(c *net.Client, pkt codec.Packet) error{
    t(packet.ServerboundHandshakePacket{}): handshakeSetProtocol,
}

var Status = map[reflect.Type]func(c *net.Client, pkt codec.Packet) error{
    t(packet.ServerboundStatusRequestPacket{}): statusRequest,
    t(packet.ServerboundStatusPingPacket{}):    statusPing,
}

var Login = map[reflect.Type]func(c *net.Client, pkt codec.Packet) error{
    t(packet.ServerboundLoginStartPacket{}):              loginStart,
    t(packet.ServerboundLoginEncryptionResponsePacket{}): loginEncryptionResponse,
}

var Play = map[reflect.Type]func(c *net.Client, pkt codec.Packet) error{
    t(packet.ServerboundClientSettingsPacket{}): playClientSettings,
    t(packet.ServerboundPluginMessagePacket{}):  playPluginMessage,
}
