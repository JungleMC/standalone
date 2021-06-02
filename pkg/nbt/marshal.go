package nbt

import (
	"bytes"
	"reflect"
)

func Marshal(tagName string, value interface{}) []byte {
	var tagType namedTagType
	if value == nil {
		tagType = tagCompound
	} else {
		tagType = typeOf(reflect.TypeOf(value))
	}

	buf := &bytes.Buffer{}
	buf.Write(writeTagType(tagType))
	buf.Write(writeString(tagName))
	buf.Write(writeValue(tagType, value))
	return buf.Bytes()
}

func writeValue(tagType namedTagType, value interface{}) []byte {
	v := reflect.ValueOf(value)

	switch tagType {
	case tagByte:
		if reflect.TypeOf(value).Kind() == reflect.Bool {
			if v.Bool() {
				return writeByte(1)
			} else {
				return writeByte(0)
			}
		}
		return writeByte(byte(v.Uint()))
	case tagShort:
		return writeInt16(int16(v.Int()))
	case tagInt:
		return writeInt32(int32(v.Int()))
	case tagLong:
		return writeInt64(v.Int())
	case tagFloat:
		return writeFloat32(float32(v.Float()))
	case tagDouble:
		return writeFloat64(v.Float())
	case tagString:
		return writeString(v.String())
	case tagList:
		return writeList(v)
	case tagCompound:
		return writeCompound(value)
	case tagByteArray:
		return writeByteSlice(v.Bytes())
	case tagIntArray:
		return writeInt32Slice(v)
	case tagLongArray:
		return writeInt64Slice(v)
	}
	return nil
}

func typeOf(t reflect.Type) namedTagType {
	switch t.Kind() {
	case reflect.Uint8, reflect.Bool:
		return tagByte
	case reflect.Int16, reflect.Uint16:
		return tagShort
	case reflect.Int32, reflect.Uint32:
		return tagInt
	case reflect.Float32:
		return tagFloat
	case reflect.Int64, reflect.Uint64:
		return tagLong
	case reflect.Float64:
		return tagDouble
	case reflect.String:
		return tagString
	case reflect.Struct, reflect.Interface, reflect.Map:
		return tagCompound
	case reflect.Array, reflect.Slice:
		switch t.Elem().Kind() {
		case reflect.Uint8:
			return tagByteArray
		case reflect.Int32:
			return tagIntArray
		case reflect.Int64:
			return tagLongArray
		default:
			return tagList
		}
	default:
		return tagNone
	}
}
