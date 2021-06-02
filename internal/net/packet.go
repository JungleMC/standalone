package net

import (
	"bufio"
	"bytes"
	"compress/zlib"
	. "github.com/junglemc/JungleTree/internal/net/codec"
	. "github.com/junglemc/JungleTree/internal/net/protocol"
	. "github.com/junglemc/JungleTree/pkg/codec"
	"reflect"
)

type Packet interface{}

func ReadPacket(buf *bytes.Buffer, proto Protocol, compressed bool) (Packet, error) {
	payloadCheck, err := buf.ReadByte()
	if payloadCheck == 0xFE {
		legacyPing(buf)
		return nil, nil
	} else {
		_ = buf.UnreadByte()
	}

	var uncompressedLength int32
	var reader *bufio.Reader
	if compressed {
		compressedLength := ReadVarInt32(buf)
		uncompressedLength = ReadVarInt32(buf)

		if uncompressedLength > 0 {
			data := make([]byte, compressedLength)
			_, _ = buf.Read(data)
			zlibReader, err := zlib.NewReader(bytes.NewBuffer(data))
			if err != nil {
				return nil, err
			}
			reader = bufio.NewReader(zlibReader)
		} else {
			reader = bufio.NewReader(buf)
		}
	} else {
		uncompressedLength = ReadVarInt32(buf)
		if err != nil {
			return nil, err
		}
		reader = bufio.NewReader(buf)
	}

	data := make([]byte, uncompressedLength)
	_, err = reader.Read(data)
	if err != nil {
		return nil, err
	}

	// Redefine the bytes reader here
	buf = bytes.NewBuffer(data)
	id := ReadVarInt32(buf)

	pktType := Registry.ServerboundType(id, proto)
	if pktType == nil {
		panic("nil type")
	}

	pkt := reflect.New(pktType).Elem()
	err = Unmarshal(buf.Bytes(), pkt)
	if err != nil {
		return nil, err
	}
	return pkt.Interface().(Packet), err
}

func WritePacket(buf *bytes.Buffer, v reflect.Value, proto Protocol, compressed bool, compressionThreshold int) {
	if v.Kind() == reflect.Interface {
		v = reflect.ValueOf(v.Interface())
	}

	id := Registry.ClientboundID(v.Type(), proto)

	packet := encodePacket(id, v.Interface())

	if compressed {
		if len(packet) >= compressionThreshold {
			packet = compress(packet)
		} else {
			packet = append(WriteVarInt32(0), packet...)
		}
	}

	buf.Write(WriteVarInt32(int32(len(packet))))
	buf.Write(packet)
}

func encodePacket(id int32, pkt interface{}) []byte {
	return append(WriteVarInt32(id), Marshal(pkt)...)
}

func compress(data []byte) []byte {
	buf := &bytes.Buffer{}
	writer := zlib.NewWriter(buf)
	writer.Write(data)
	writer.Flush()
	return append(WriteVarInt32(int32(len(data))), buf.Bytes()...)
}

func legacyPing(buf *bytes.Buffer) {
	buf.ReadByte()
	buf.ReadByte()

	buf.ReadByte() // packet identifier for a plugin message

	mcPingHostLength := ReadUint16(buf)
	mcPingHost := make([]byte, mcPingHostLength)
	buf.Read(mcPingHost)

	ReadInt16(buf) // Remaining
	buf.ReadByte()

	hostnameLength := ReadInt16(buf)
	hostname := make([]byte, hostnameLength)
	buf.Read(hostname)
	ReadInt16(buf)
}
