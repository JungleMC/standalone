package world

import (
	"bytes"
	"github.com/junglemc/JungleTree/net/codec"
)

type BlockPosition struct {
	X int32
	Y int32
	Z int32
}

func (b *BlockPosition) MarshalMinecraft() ([]byte, error) {
	var result [8]byte
	position := uint64(b.X&0x3FFFFFF)<<38 | uint64((b.Z&0x3FFFFFF)<<12) | uint64(b.Y&0xFFF)
	for i := 7; i >= 0; i-- {
		result[i] = byte(position)
		position >>= 8
	}
	return result[:], nil
}

func (b *BlockPosition) UnmarshalMinecraft(buf *bytes.Buffer) error {
	val := codec.ReadInt64(buf)

	b.X = int32(val >> 38)
	b.Y = int32(val & 0xFFF)
	b.Z = int32(val << 26 >> 38)

	// Negative numbers
	if b.X >= 1<<25 {
		b.X -= 1 << 26
	}
	if b.Y >= 1<<11 {
		b.Y -= 1 << 12
	}
	if b.Z >= 1<<25 {
		b.Z -= 1 << 26
	}
	return nil
}