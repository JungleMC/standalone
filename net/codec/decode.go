package codec

import (
	"bytes"
	"errors"
	"github.com/junglemc/JungleTree/nbt"
	"reflect"
)

func Unmarshal(data []byte, val reflect.Value) error {
	if val.Kind() == reflect.Ptr || (val.Kind() == reflect.Interface && val.IsNil()) {
		return errors.New("invalid interface, cannot be ptr or nil interface")
	}

	if !val.CanSet() {
		return errors.New("unable to set")
	}

	buf := bytes.NewBuffer(data)
	decodedVal := typeDecoder(val.Type())(buf)
	val.Set(decodedVal)
	return nil
}

type Unmarshaler interface {
	UnmarshalMinecraft(buf *bytes.Buffer) error
}

var unmarshalType = reflect.TypeOf((*Unmarshaler)(nil)).Elem()

type decoderFunc func(buf *bytes.Buffer) reflect.Value

func typeDecoder(t reflect.Type) decoderFunc {
	if t.Kind() != reflect.Ptr && reflect.PtrTo(t).Implements(unmarshalType) {
		return unmarshalDecoder{t: t}.decode
	}

	if t.Implements(unmarshalType) {
		return unmarshalDecoder{t: t.Elem()}.decode
	}

	switch t.Kind() {
	case reflect.Bool:
		return boolDecoder
	case reflect.Uint8:
		return uint8Decoder
	case reflect.Uint16:
		return uint16Decoder
	case reflect.Uint32:
		return uint32Decoder
	case reflect.Uint64, reflect.Uintptr:
		return uint64Decoder
	case reflect.Int8:
		return int8Decoder
	case reflect.Int16:
		return int16Decoder
	case reflect.Int32:
		return int32Decoder
	case reflect.Int64:
		return int64Decoder
	case reflect.Float32:
		return float32Decoder
	case reflect.Float64:
		return float64Decoder
	case reflect.String:
		return stringDecoder
	case reflect.Interface:
		return interfaceDecoder{ifaceType: t}.decode
	case reflect.Struct:
		return structDecoder{structType: t}.decode
	case reflect.Map:
		return mapDecoder
	case reflect.Slice:
		return sliceDecoder{t: t}.decode
	case reflect.Array:
		return arrayDecoder{arrayType: t}.decode
	default:
		panic(errors.New("unsupported type: " + t.String()))
	}
}

func structTypeDecoder(t reflect.StructField, opts codecOpts) decoderFunc {
	switch t.Type.Kind() {
	case reflect.Struct:
		return structDecoder{structType: t.Type, nbt: opts.nbt}.decode
	case reflect.Slice:
		return sliceDecoder{t: t.Type, inferSize: opts.inferSize}.decode
	case reflect.Array:
		return arrayDecoder{arrayType: t.Type, inferSize: opts.inferSize}.decode
	case reflect.Int32:
		if opts.varint {
			return varInt32Decoder
		}
		break
	case reflect.Int64:
		if opts.varint {
			return varInt64Decoder
		}
		break
	}

	return typeDecoder(t.Type)
}

type unmarshalDecoder struct {
	t reflect.Type
}

func (u unmarshalDecoder) decode(buf *bytes.Buffer) reflect.Value {
	v := reflect.New(u.t)
	m, ok := v.Interface().(Unmarshaler)
	if !ok {
		panic(errors.New("no unmarshal function"))
	}
	err := m.UnmarshalMinecraft(buf)
	if err != nil {
		panic(err)
	}
	return v
}

type interfaceDecoder struct {
	ifaceType reflect.Type
}

func (i interfaceDecoder) decode(buf *bytes.Buffer) reflect.Value {
	return typeDecoder(i.ifaceType.Elem())(buf)
}

func boolDecoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadBool(buf))
}

func uint8Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadUint8(buf))
}

func uint16Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadUint16(buf))
}

func uint32Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadUint32(buf))
}

func uint64Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadUint64(buf))
}

func int8Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadInt8(buf))
}

func int16Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadInt16(buf))
}

func int32Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadInt32(buf))
}

func int64Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadInt64(buf))
}

func float32Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadFloat32(buf))
}

func float64Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadFloat64(buf))
}

func stringDecoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadString(buf))
}

func varInt32Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadVarInt32(buf))
}

func varInt64Decoder(buf *bytes.Buffer) reflect.Value {
	return reflect.ValueOf(ReadVarInt64(buf))
}

type structDecoder struct {
	structType reflect.Type
	nbt        bool
	nbtTagName string
}

func (s structDecoder) decode(buf *bytes.Buffer) reflect.Value {
	v := reflect.New(s.structType).Elem()

	if s.nbt {
		var err error
		s.nbtTagName, err = nbt.Unmarshal(buf.Bytes(), v)
		if err != nil {
			panic(err)
		}
		return v
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		fieldOpts := findOpts(fieldType)

		if fieldOpts.optional && !fieldOpts.optionalPresent(fieldType.Tag, field) {
			continue
		}

		data := structTypeDecoder(fieldType, fieldOpts)(buf)

		if !field.CanSet() {
			panic(errors.New("cannot set field: " + fieldType.Name))
		}

		if field.Kind() != reflect.Ptr && data.Kind() == reflect.Ptr {
			field.Set(data.Elem())
		} else {
			field.Set(data)
		}
	}
	return v
}

func mapDecoder(_ *bytes.Buffer) reflect.Value {
	panic(errors.New("not implemented"))
}

type sliceDecoder struct {
	t reflect.Type
	inferSize bool
}

func (s sliceDecoder) decode(buf *bytes.Buffer) reflect.Value {
	var length int

	var result reflect.Value

	if s.inferSize {
		length = buf.Len()
	} else {
		lengthValue := varInt32Decoder(buf)
		length = int(lengthValue.Int())
	}

	result = reflect.MakeSlice(s.t, length, length)

	for i := 0; i < length; i++ {
		val := typeDecoder(s.t.Elem())(buf)
		result.Index(i).Set(val)
	}
	return result
}

type arrayDecoder struct {
	arrayType reflect.Type
	inferSize bool
}

func (a arrayDecoder) decode(buf *bytes.Buffer) reflect.Value {
	result := reflect.New(a.arrayType)

	if !a.inferSize {
		_ = varInt32Decoder(buf)
	}

	for i := 0; i < result.Len(); i++ {
		val := typeDecoder(a.arrayType.Elem())(buf)
		result.Index(i).Set(val)
	}

	return result
}
