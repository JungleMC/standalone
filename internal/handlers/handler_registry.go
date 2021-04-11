package handlers

import (
    "github.com/junglemc/net"
    "github.com/junglemc/net/packet"
    "reflect"
)

var Handshake = map[reflect.Type]func(c *net.Client, pkt net.Packet){
    reflect.TypeOf(packet.ServerboundHandshakePacket{}): handshakeSetProtocol,
}

var Status = map[reflect.Type]func(c *net.Client, pkt net.Packet){
    reflect.TypeOf(packet.ServerboundStatusRequestPacket{}): statusRequest,
    reflect.TypeOf(packet.ServerboundStatusPingPacket{}):    statusPing,
}

var Login = map[reflect.Type]func(c *net.Client, pkt net.Packet){
    reflect.TypeOf(packet.ServerboundLoginStartPacket{}): loginStart,
}
