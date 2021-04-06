package handlers

import (
    "github.com/junglemc/mc/packet"
    "github.com/junglemc/net"
    "reflect"
)

var Handshake = map[reflect.Type]func(c *net.Client, pkt net.Packet){
    reflect.TypeOf(packet.ServerboundHandshakeSetProtocol{}): handshakeSetProtocol,
}

var Status = map[reflect.Type]func(c *net.Client, pkt net.Packet){
    reflect.TypeOf(packet.ServerboundStatusPingStart{}): statusRequest,
    reflect.TypeOf(packet.ServerboundStatusPing{}): statusPing,
}
