package handlers

import (
	"github.com/junglemc/net"
	"github.com/junglemc/net/packet"
	"reflect"
)

var Handshake = map[reflect.Type]func(c *net.Client, pkt net.Packet){
	reflect.TypeOf(packet.ServerboundHandshakeHelloPacket{}): handshakeHandle,
}
