package nbt

import (
	"bytes"
	"errors"
	"math"
	"reflect"
)

type namedTagType byte

const (
	tagEnd namedTagType = iota
	tagByte
	tagShort
	tagInt
	tagLong
	tagFloat
	tagDouble
	tagByteArray
	tagString
	tagList
	tagCompound
	tagIntArray
	tagLongArray
	tagNone = 0xFF
)

func readTagType(buf *bytes.Buffer) (t namedTagType, err error) {
	tb, err := buf.ReadByte()
	return namedTagType(tb), err
}

func readUInt16(buf *bytes.Buffer) (uint16, error) {
	b := make([]byte, 2)
	n, err := buf.Read(b)
	if err != nil {
		return 0, err
	}
	if n > 2 {
		return 0, errors.New("too much data")
	}
	if n < 2 {
		return 0, errors.New("not enough data")
	}
	return uint16(b[0])<<8 | uint16(b[1]), nil
}

func readInt16(buf *bytes.Buffer) (int16, error) {
	uv, err := readUInt16(buf)
	if err != nil {
		return 0, err
	}
	return int16(uv), nil
}

func readUInt32(buf *bytes.Buffer) (uint32, error) {
	b := make([]byte, 4)
	n, err := buf.Read(b)
	if err != nil {
		return 0, err
	}
	if n > 4 {
		return 0, errors.New("too much data")
	}
	if n < 4 {
		return 0, errors.New("not enough data")
	}
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3]), nil
}

func readInt32(buf *bytes.Buffer) (int32, error) {
	v, err := readUInt32(buf)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func readUInt64(buf *bytes.Buffer) (uint64, error) {
	b := make([]byte, 8)
	n, err := buf.Read(b)
	if err != nil {
		return 0, err
	}
	if n > 8 {
		return 0, errors.New("too much data")
	}
	if n < 8 {
		return 0, errors.New("not enough data")
	}
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 | uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7]), nil
}

func readInt64(buf *bytes.Buffer) (int64, error) {
	v, err := readUInt64(buf)
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

func readFloat32(buf *bytes.Buffer) (float32, error) {
	v, err := readUInt32(buf)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(v), nil
}

func readFloat64(buf *bytes.Buffer) (float64, error) {
	v, err := readUInt64(buf)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(v), nil
}

func readByteSlice(buf *bytes.Buffer) ([]byte, error) {
	length, err := readInt32(buf)
	if err != nil {
		return nil, err
	}
	v := make([]byte, length, length)
	readBytes, err := buf.Read(v)
	if readBytes > int(length) {
		return v, errors.New("read too many bytes")
	}
	if readBytes < int(length) {
		return v, errors.New("read too few bytes")
	}
	return v, nil
}

func readInt32Slice(buf *bytes.Buffer) ([]int32, error) {
	length, err := readInt32(buf)
	if err != nil {
		return nil, err
	}
	v := make([]int32, length, length)
	for i := 0; i < int(length); i++ {
		v[i], err = readInt32(buf)
		if err != nil {
			return v, err
		}
	}
	return v, nil
}

func readInt64Slice(buf *bytes.Buffer) ([]int64, error) {
	length, err := readInt32(buf)
	if err != nil {
		return nil, err
	}
	v := make([]int64, length, length)
	for i := 0; i < int(length); i++ {
		v[i], err = readInt64(buf)
		if err != nil {
			return v, err
		}
	}
	return v, nil
}

func readString(buf *bytes.Buffer) (string, error) {
	length, err := readUInt16(buf)
	if err != nil {
		return "", err
	}
	if length == 0 {
		return "", nil
	}

	v := make([]byte, length, length)
	readBytes, err := buf.Read(v)
	if err != nil {
		return "", err
	}
	if readBytes > int(length) {
		return "", errors.New("read too many bytes")
	}
	if readBytes < int(length) {
		return "", errors.New("read too few bytes")
	}
	return string(v), nil
}

func writeTagType(t namedTagType) []byte {
	return []byte{byte(t)}
}

func writeByte(v byte) []byte {
	return []byte{v}
}

func writeUInt16(v uint16) []byte {
	return []byte{byte(v >> 8), byte(v)}
}

func writeInt16(v int16) []byte {
	return writeUInt16(uint16(v))
}

func writeUInt32(v uint32) []byte {
	return []byte{byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
}

func writeInt32(v int32) []byte {
	return writeUInt32(uint32(v))
}

func writeUInt64(v uint64) []byte {
	return []byte{byte(v >> 56), byte(v >> 48), byte(v >> 40), byte(v >> 32), byte(v >> 24), byte(v >> 16), byte(v >> 8), byte(v)}
}

func writeInt64(v int64) []byte {
	return writeUInt64(uint64(v))
}

func writeFloat32(v float32) []byte {
	return writeUInt32(math.Float32bits(v))
}

func writeFloat64(v float64) []byte {
	return writeUInt64(math.Float64bits(v))
}

func writeByteSlice(v []byte) []byte {
	return append(writeInt32(int32(len(v))), v...)
}

func writeInt32Slice(v reflect.Value) []byte {
	buf := &bytes.Buffer{}
	buf.Write(writeInt32(int32(v.Len())))
	for i := 0; i < v.Len(); i++ {
		buf.Write(writeInt32(int32(v.Index(i).Int())))
	}
	return buf.Bytes()
}

func writeInt64Slice(v reflect.Value) []byte {
	buf := &bytes.Buffer{}
	buf.Write(writeInt32(int32(v.Len())))
	for i := 0; i < v.Len(); i++ {
		buf.Write(writeInt64(v.Index(i).Int()))
	}
	return buf.Bytes()
}

func writeString(v string) []byte {
	return append(writeUInt16(uint16(len(v))), []byte(v)...)
}

func writeList(v reflect.Value) []byte {
	buf := &bytes.Buffer{}

	nestedTagType := typeOf(v.Type().Elem())
	if v.Len() <= 0 {
		nestedTagType = tagEnd // Mimic notchian behavior
	}
	buf.Write(writeTagType(nestedTagType))

	buf.Write(writeInt32(int32(v.Len())))
	for i := 0; i < v.Len(); i++ {
		buf.Write(writeValue(nestedTagType, v.Index(i).Interface()))
	}
	return buf.Bytes()
}

func writeCompound(value interface{}) []byte {
	if value == nil {
		return writeTagType(tagEnd)
	}

	buf := &bytes.Buffer{}

	v := reflect.ValueOf(value)
	if v.Type().Kind() == reflect.Map {
		mapRange := v.MapRange()
		for mapRange.Next() {
			nestedTagType := typeOf(reflect.TypeOf(mapRange.Value().Interface()))

			buf.Write(writeTagType(nestedTagType))
			buf.Write(writeString(mapRange.Key().String()))
			buf.Write(writeValue(nestedTagType, mapRange.Value().Interface()))
		}
	} else {
		numFields := v.NumField()

		for i := 0; i < numFields; i++ {
			f := v.Type().Field(i)

			nestedTagName := f.Tag.Get("nbt")

			// Take the field name if unspecified
			if nestedTagName == "" {
				nestedTagName = f.Name
			}

			// Ignore unwanted tags
			if nestedTagName == "-" {
				continue
			}

			nestedTagType := typeOf(f.Type)
			if f.Tag.Get("nbt_type") == "list" {
				nestedTagType = tagList
			}

			isPresent := true

			tag := f.Tag
			optional := tag.Get("optional")
			if optional != "" {
				optionalField := v.FieldByName(optional)
				if optionalField.Type().Kind() == reflect.Bool {
					isPresent = optionalField.Bool()
				} else {
					panic(errors.New("optional field type not handled: " + optionalField.Kind().String()))
				}
			}

			// Ignore if the present field is not true
			if !isPresent {
				continue
			}

			buf.Write(writeTagType(nestedTagType))
			buf.Write(writeString(nestedTagName))
			buf.Write(writeValue(nestedTagType, v.Field(i).Interface()))
		}
	}

	buf.Write(writeTagType(tagEnd))
	return buf.Bytes()
}
