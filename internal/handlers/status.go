package handlers

import (
    "github.com/junglemc/net"
    "log"
)

func statusPingStart(c *net.Client, p net.Packet) {
    log.Println("Hello!")
}
