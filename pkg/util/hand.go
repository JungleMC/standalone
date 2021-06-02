package util

import (
	"bytes"
	"github.com/junglemc/JungleTree/pkg/codec"
)

type Hand int32

const (
	HandLeft Hand = iota
	HandRight
)

func (m *Hand) MarshalMinecraft() ([]byte, error) {
	return codec.WriteVarInt32(int32(*m)), nil
}

func (m *Hand) UnmarshalMinecraft(buf *bytes.Buffer) error {
	*m = Hand(codec.ReadVarInt32(buf))
	return nil
}
