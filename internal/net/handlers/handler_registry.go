package handlers

import (
	. "reflect"

	. "github.com/junglemc/JungleTree/internal/net"
	. "github.com/junglemc/JungleTree/internal/pkg/net/packets"
)

var Handshake = map[Type]func(c *Client, pkt Packet) error{
	TypeOf(ServerboundHandshakePacket{}): handshakeSetProtocol,
}

var Status = map[Type]func(c *Client, pkt Packet) error{
	TypeOf(ServerboundStatusRequestPacket{}): statusRequest,
	TypeOf(ServerboundStatusPingPacket{}):    statusPing,
}

var Login = map[Type]func(c *Client, pkt Packet) error{
	TypeOf(ServerboundLoginStartPacket{}):              loginStart,
	TypeOf(ServerboundLoginEncryptionResponsePacket{}): loginEncryptionResponse,
}

var Play = map[Type]func(c *Client, pkt Packet) error{
	TypeOf(ServerboundClientSettingsPacket{}): playClientSettings,
	TypeOf(ServerboundPluginMessagePacket{}):  playPluginMessage,
}
