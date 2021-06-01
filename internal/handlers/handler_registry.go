package handlers

import (
	"github.com/junglemc/net"
	packet2 "github.com/junglemc/packet"
	"reflect"
)

func t(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}

var Handshake = map[reflect.Type]func(c *net.Client, pkt net.Packet) error{
	t(packet2.ServerboundHandshakePacket{}): handshakeSetProtocol,
}

var Status = map[reflect.Type]func(c *net.Client, pkt net.Packet) error{
	t(packet2.ServerboundStatusRequestPacket{}): statusRequest,
	t(packet2.ServerboundStatusPingPacket{}):    statusPing,
}

var Login = map[reflect.Type]func(c *net.Client, pkt net.Packet) error{
	t(packet2.ServerboundLoginStartPacket{}):              loginStart,
	t(packet2.ServerboundLoginEncryptionResponsePacket{}): loginEncryptionResponse,
}

var Play = map[reflect.Type]func(c *net.Client, pkt net.Packet) error{
	t(packet2.ServerboundClientSettingsPacket{}): playClientSettings,
	t(packet2.ServerboundPluginMessagePacket{}):  playPluginMessage,
}
