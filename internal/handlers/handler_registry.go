package handlers

import (
	"github.com/junglemc/JungleTree/net"
	. "github.com/junglemc/JungleTree/packet"
	"reflect"
)

func t(i interface{}) reflect.Type {
	return reflect.TypeOf(i)
}

var Handshake = map[reflect.Type]func(c *net.Client, pkt net.Packet) error{
	t(ServerboundHandshakePacket{}): handshakeSetProtocol,
}

var Status = map[reflect.Type]func(c *net.Client, pkt net.Packet) error{
	t(ServerboundStatusRequestPacket{}): statusRequest,
	t(ServerboundStatusPingPacket{}):    statusPing,
}

var Login = map[reflect.Type]func(c *net.Client, pkt net.Packet) error{
	t(ServerboundLoginStartPacket{}):              loginStart,
	t(ServerboundLoginEncryptionResponsePacket{}): loginEncryptionResponse,
}

var Play = map[reflect.Type]func(c *net.Client, pkt net.Packet) error{
	t(ServerboundClientSettingsPacket{}): playClientSettings,
	t(ServerboundPluginMessagePacket{}):  playPluginMessage,
}
