package protocol

import (
	"fmt"
	"github.com/junglemc/JungleTree/packet"
	"reflect"
)

var Reg = Registry{
	Handshake: func() (clientbound map[int32]reflect.Type, serverbound map[int32]reflect.Type) {
		return packet.HandshakeClientboundIds, packet.HandshakeServerboundIds
	},
	Status: func() (clientbound map[int32]reflect.Type, serverbound map[int32]reflect.Type) {
		return packet.StatusClientboundIds, packet.StatusServerboundIds
	},
	Login: func() (clientbound map[int32]reflect.Type, serverbound map[int32]reflect.Type) {
		return packet.LoginClientboundIds, packet.LoginServerboundIds
	},
	Play: func() (clientbound map[int32]reflect.Type, serverbound map[int32]reflect.Type) {
		return packet.PlayClientboundIds, packet.PlayServerboundIds
	},
}

type Registry struct {
	Handshake RegistryMap
	Status    RegistryMap
	Login     RegistryMap
	Play      RegistryMap
}

type RegistryMap func() (clientbound map[int32]reflect.Type, serverbound map[int32]reflect.Type)

func (r *Registry) ClientboundID(t reflect.Type, p Protocol) int32 {
	var clientbound map[int32]reflect.Type

	switch p {
	case ProtocolHandshake:
		clientbound, _ = r.Handshake()
		break
	case ProtocolStatus:
		clientbound, _ = r.Status()
		break
	case ProtocolLogin:
		clientbound, _ = r.Login()
		break
	case ProtocolPlay:
		clientbound, _ = r.Play()
		break
	}

	for id, pkt := range clientbound {
		if pkt.Name() == t.Name() {
			return id
		}
	}

	panic("not found") // TODO: Cleanup
}

func (r *Registry) ServerboundType(id int32, p Protocol) reflect.Type {
	var serverbound map[int32]reflect.Type

	switch p {
	case ProtocolHandshake:
		_, serverbound = r.Handshake()
		break
	case ProtocolStatus:
		_, serverbound = r.Status()
		break
	case ProtocolLogin:
		_, serverbound = r.Login()
		break
	case ProtocolPlay:
		_, serverbound = r.Play()
		break
	}

	result := serverbound[id]
	if result == nil {
		panic(fmt.Sprintf("not found: packetID=0x%02X", id)) // TODO: Cleanup
	}
	return result
}
