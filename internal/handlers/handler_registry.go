package handlers

import (
    "github.com/junglemc/net"
    "github.com/junglemc/net/packet"
    "reflect"
)

var Handshake = map[reflect.Type]func(c *net.Client, pkt net.Packet){
    reflect.TypeOf(packet.ServerboundHandshakeSetProtocol{}): handshakeSetProtocol,
}

var Status = map[reflect.Type]func(c *net.Client, pkt net.Packet){
    reflect.TypeOf(packet.ServerboundStatusPingStart{}): statusRequest,
    reflect.TypeOf(packet.ServerboundStatusPing{}):      statusPing,
}

var Login = map[reflect.Type]func(c *net.Client, pkt net.Packet){
    reflect.TypeOf(packet.ServerboundLoginStart{}): loginStart,
}
