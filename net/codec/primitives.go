package codec

import (
	"bytes"
	"errors"
	"math"
)

const stringMaxLength = 32767

func ReadBool(buf *bytes.Buffer) bool {
	b, _ := buf.ReadByte()
	return b == 0x01
}

func ReadUint8(buf *bytes.Buffer) uint8 {
	b, _ := buf.ReadByte()
	return b
}

func ReadUint16(buf *bytes.Buffer) uint16 {
	b := make([]byte, 2)
	_, _ = buf.Read(b)
	return uint16(b[0])<<8 | uint16(b[1])
}

func ReadUint32(buf *bytes.Buffer) uint32 {
	b := make([]byte, 4)
	_, _ = buf.Read(b)
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

func ReadUint64(buf *bytes.Buffer) uint64 {
	b := make([]byte, 8)
	_, _ = buf.Read(b)
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
}

func ReadInt8(buf *bytes.Buffer) int8 {
	return int8(ReadUint8(buf))
}

func ReadInt16(buf *bytes.Buffer) int16 {
	return int16(ReadUint16(buf))
}

func ReadInt32(buf *bytes.Buffer) int32 {
	return int32(ReadUint32(buf))
}

func ReadInt64(buf *bytes.Buffer) int64 {
	return int64(ReadUint64(buf))
}

func ReadFloat32(buf *bytes.Buffer) float32 {
	return math.Float32frombits(ReadUint32(buf))
}

func ReadFloat64(buf *bytes.Buffer) float64 {
	return math.Float64frombits(ReadUint64(buf))
}

func ReadString(buf *bytes.Buffer) string {
	length := ReadVarInt32(buf)
	data := make([]byte, length, length)
	_, _ = buf.Read(data)
	return string(data)
}

func ReadVarInt32(buf *bytes.Buffer) int32 {
	numRead := int32(0)
	result := int32(0)

	for {
		read, _ := buf.ReadByte()
		value := read & 0b01111111
		result |= int32(value) << (7 * numRead)
		numRead++
		if numRead > 5 {
			panic(errors.New("varint32 is too big"))
		}

		if (read & 0b10000000) == 0 {
			break
		}
	}
	return result
}

func ReadVarInt64(buf *bytes.Buffer) int64 {
	numRead := int32(0)
	result := int64(0)

	for {
		read, _ := buf.ReadByte()
		value := read & 0b01111111
		result |= int64(value) << (7 * numRead)
		numRead++
		if numRead > 10 {
			panic(errors.New("varint64 is too big"))
		}

		if (read & 0b10000000) == 0 {
			break
		}
	}
	return result
}

func WriteBool(v bool) []byte {
	if v {
		return []byte{0x01}
	}
	return []byte{0x00}
}

func WriteUint8(v uint8) []byte {
	return []byte{v}
}

func WriteUint16(v uint16) []byte {
	return []byte{byte(v >> 8), byte(v)}
}

func WriteUint32(v uint32) []byte {
	return []byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
}

func WriteUint64(v uint64) []byte {
	return []byte{byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32), byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
}

func WriteInt8(v int8) []byte {
	return WriteUint8(uint8(v))
}

func WriteInt16(v int16) []byte {
	return WriteUint16(uint16(v))
}

func WriteInt32(v int32) []byte {
	return WriteUint32(uint32(v))
}

func WriteInt64(v int64) []byte {
	return WriteUint64(uint64(v))
}

func WriteFloat32(v float32) []byte {
	return WriteUint32(math.Float32bits(v))
}

func WriteFloat64(v float64) []byte {
	return WriteUint64(math.Float64bits(v))
}

func WriteString(v string) []byte {
	if len(v) > stringMaxLength {
		panic(errors.New("string exceeds max length"))
	}
	data := WriteVarInt32(int32(len(v)))
	return append(data, []byte(v)...)
}

func WriteVarInt32(v int32) []byte {
	buf := bytes.Buffer{}
	uvalue := uint32(v)

	for {
		temp := byte(uvalue & 0b01111111)
		uvalue >>= 7
		if uvalue != 0 {
			temp |= 0b10000000
		}
		_ = buf.WriteByte(temp)

		if uvalue == 0 {
			break
		}
	}
	return buf.Bytes()
}

func WriteVarInt64(v int64) []byte {
	buf := bytes.Buffer{}
	uvalue := uint64(v)
	for {
		temp := byte(uvalue & 0b01111111)
		uvalue >>= 7
		if uvalue != 0 {
			temp |= 0b10000000
		}
		buf.WriteByte(temp)
		if uvalue == 0 {
			break
		}
	}
	return buf.Bytes()
}
