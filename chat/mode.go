package chat

import (
	"bytes"
	. "github.com/junglemc/JungleTree/net/codec"
)

type Mode int32

const (
	Enabled Mode = iota
	CommandsOnly
	Hidden
)

func (m *Mode) MarshalMinecraft() ([]byte, error) {
	return WriteVarInt32(int32(*m)), nil
}

func (m *Mode) UnmarshalMinecraft(buf *bytes.Buffer) error {
	*m = Mode(ReadVarInt32(buf))
	return nil
}