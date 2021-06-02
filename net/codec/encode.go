package codec

import (
	"bytes"
	"errors"
	"github.com/junglemc/nbt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func Marshal(v interface{}) []byte {
	val := reflect.ValueOf(v)
	return typeEncoder(val.Type(), true)(val)
}

type Marshaler interface {
	MarshalMinecraft() ([]byte, error)
}

type codecOpts struct {
	name                string
	nbt                 bool
	varint              bool
	inferSize           bool
	optional            bool
	optionalField       string
	optionalFieldValues []int64
}

var marshalType = reflect.TypeOf((*Marshaler)(nil)).Elem()

type encoderFunc func(v reflect.Value) []byte

func typeEncoder(t reflect.Type, allowAddr bool) encoderFunc {
	if t.Kind() == reflect.Ptr && allowAddr && reflect.PtrTo(t).Implements(marshalType) {
		newCondAddrEncoder(addrMarshalEncoder, typeEncoder(t, false))
	}

	if t.Implements(marshalType) {
		return marshalEncoder
	}

	switch t.Kind() {
	case reflect.Bool:
		return boolEncoder
	case reflect.Uint8:
		return uint8Encoder
	case reflect.Uint16:
		return uint16Encoder
	case reflect.Uint32:
		return uint32Encoder
	case reflect.Uint64, reflect.Uintptr:
		return uint64Encoder
	case reflect.Int8:
		return int8Encoder
	case reflect.Int16:
		return int16Encoder
	case reflect.Int32:
		return int32Encoder
	case reflect.Int64:
		return int64Encoder
	case reflect.Float32:
		return float32Encoder
	case reflect.Float64:
		return float64Encoder
	case reflect.String:
		return stringEncoder
	case reflect.Interface:
		return interfaceEncoder
	case reflect.Struct:
		return structEncoder
	case reflect.Map:
		return mapEncoder
	case reflect.Slice:
		return sliceEncoder
	case reflect.Array:
		return arrayEncoder
	case reflect.Ptr:
		return ptrEncoder
	default:
		panic(errors.New("unsupported type: " + t.String()))
	}
}

type nbtEncoder struct {
	tagName string
}

func (s nbtEncoder) encode(v reflect.Value) []byte {
	return nbt.Marshal(s.tagName, v.Interface())
}

func structTypeEncoder(t reflect.StructField, allowAddr bool, opts codecOpts) encoderFunc {
	if opts.nbt {
		return nbtEncoder{tagName: opts.name}.encode
	}

	switch t.Type.Kind() {
	case reflect.Slice:
		if opts.inferSize {
			return sliceEncoderSizeInference
		}
		break
	case reflect.Array:
		if opts.inferSize {
			return arrayEncoderSizeInference
		}
		break
	case reflect.Int32:
		if opts.varint {
			return varInt32Encoder
		}
		break
	case reflect.Int64:
		if opts.varint {
			return varInt64Encoder
		}
		break
	}

	return typeEncoder(t.Type, allowAddr)
}

type condAddrEncoder struct {
	canAddrEnc, elseEnc encoderFunc
}

func newCondAddrEncoder(canAddrEnc, elseEnc encoderFunc) encoderFunc {
	enc := condAddrEncoder{canAddrEnc: canAddrEnc, elseEnc: elseEnc}
	return enc.encode
}

func (ce condAddrEncoder) encode(v reflect.Value) []byte {
	if v.CanAddr() {
		return ce.canAddrEnc(v)
	} else {
		return ce.elseEnc(v)
	}
}

func marshalEncoder(v reflect.Value) []byte {
	if v.Kind() == reflect.Ptr && v.IsNil() {
		panic(errors.New("nil value"))
	}
	m, ok := v.Interface().(Marshaler)
	if !ok {
		panic(errors.New("no marshal function"))
	}

	buf := bytes.Buffer{}
	b, err := m.MarshalMinecraft()
	if err != nil {
		panic(err)
	}
	buf.Write(b)

	return buf.Bytes()
}

func addrMarshalEncoder(v reflect.Value) []byte {
	va := v.Addr()
	if va.IsNil() {
		panic(errors.New("nil value"))
	}
	m := va.Interface().(Marshaler)

	buf := bytes.Buffer{}
	b, err := m.MarshalMinecraft()
	if err != nil {
		panic(err)
	}
	buf.Write(b)

	return buf.Bytes()
}

func interfaceEncoder(v reflect.Value) []byte {
	if v.IsNil() {
		log.Println(errors.New("nil value"))
		return []byte{}
	}
	return typeEncoder(v.Elem().Type(), true)(v.Elem())
}

func boolEncoder(v reflect.Value) []byte {
	if v.Bool() {
		return []byte{byte(0x01)}
	}
	return WriteBool(v.Bool())
}

func uint8Encoder(v reflect.Value) []byte {
	return WriteUint8(uint8(v.Uint()))
}

func uint16Encoder(v reflect.Value) []byte {
	return WriteUint16(uint16(v.Uint()))
}

func uint32Encoder(v reflect.Value) []byte {
	return WriteUint32(uint32(v.Uint()))
}

func uint64Encoder(v reflect.Value) []byte {
	return WriteUint64(v.Uint())
}

func int8Encoder(v reflect.Value) []byte {
	return WriteInt8(int8(v.Int()))
}

func int16Encoder(v reflect.Value) []byte {
	return WriteInt16(int16(v.Int()))
}

func int32Encoder(v reflect.Value) []byte {
	return WriteInt32(int32(v.Int()))
}

func int64Encoder(v reflect.Value) []byte {
	return WriteInt64(v.Int())
}

func float32Encoder(v reflect.Value) []byte {
	return WriteFloat32(float32(v.Float()))
}

func float64Encoder(v reflect.Value) []byte {
	return WriteFloat64(v.Float())
}

func stringEncoder(v reflect.Value) []byte {
	return WriteString(v.String())
}

func varInt32Encoder(v reflect.Value) []byte {
	return WriteVarInt32(int32(v.Int()))
}

func varInt64Encoder(v reflect.Value) []byte {
	return WriteVarInt64(v.Int())
}

func structEncoder(v reflect.Value) []byte {
	buf := bytes.Buffer{}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)
		fieldOpts := findOpts(fieldType)

		// Optional not present, continuing
		if fieldOpts.optional && !fieldOpts.optionalPresent(fieldType.Tag, v) {
			continue
		}

		data := structTypeEncoder(fieldType, true, fieldOpts)(field)
		buf.Write(data)
	}

	return buf.Bytes()
}

func mapEncoder(v reflect.Value) []byte {
	buf := bytes.Buffer{}

	mapRange := v.MapRange()
	for mapRange.Next() {
		buf.Write(typeEncoder(mapRange.Value().Type(), true)(mapRange.Value()))
	}

	return buf.Bytes()
}

func sliceEncoder(v reflect.Value) []byte {
	buf := bytes.Buffer{}

	buf.Write(varInt32Encoder(reflect.ValueOf(int32(v.Len()))))

	if v.Len() > 0 {
		sliceType := v.Type().Elem()
		encoder := typeEncoder(sliceType, true)

		for i := 0; i < v.Len(); i++ {
			buf.Write(encoder(v.Index(i)))
		}
	}

	return buf.Bytes()
}

func sliceEncoderSizeInference(v reflect.Value) []byte {
	buf := bytes.Buffer{}

	if v.Len() > 0 {
		sliceType := v.Type().Elem()
		encoder := typeEncoder(sliceType, true)

		for i := 0; i < v.Len(); i++ {
			buf.Write(encoder(v.Index(i)))
		}
	}

	return buf.Bytes()
}

func arrayEncoder(v reflect.Value) []byte {
	buf := bytes.Buffer{}

	buf.Write(varInt32Encoder(reflect.ValueOf(int32(v.Len()))))

	if v.Len() > 0 {
		arrayType := v.Type().Elem()
		encoder := typeEncoder(arrayType, true)

		for i := 0; i < v.Len(); i++ {
			data := encoder(v.Index(i))
			buf.Write(data)
		}
	}

	return buf.Bytes()
}

func arrayEncoderSizeInference(v reflect.Value) []byte {
	buf := bytes.Buffer{}

	if v.Len() > 0 {
		arrayType := v.Type().Elem()
		encoder := typeEncoder(arrayType, true)

		for i := 0; i < v.Len(); i++ {
			data := encoder(v.Index(i))
			buf.Write(data)
		}
	}

	return buf.Bytes()
}

func (o *codecOpts) optionalPresent(tag reflect.StructTag, val reflect.Value) (present bool) {
	optionalTag := tag.Get("optional")
	if optionalTag != "" {
		optionalSplit := strings.Split(optionalTag, "=")

		optionalField := val.FieldByName(optionalSplit[0])
		switch optionalField.Type().Kind() {
		case reflect.Bool:
			return optionalField.Bool()
		case reflect.Int32:
			for i := 0; i < len(optionalSplit[1:]); {
				equalValue, err := strconv.ParseInt(optionalSplit[i], 10, 32)
				if err != nil {
					return false
				}
				return optionalField.Int() == equalValue
			}
			break
		}
		return false
	}
	return true
}

func findOpts(fieldType reflect.StructField) codecOpts {
	result := codecOpts{}

	netTag := fieldType.Tag.Get("net")
	if netTag == "" {
		result.name = fieldType.Name
	} else {
		result.name = netTag
	}

	typeTag := fieldType.Tag.Get("type")
	if typeTag == "nbt" {
		result.nbt = true
	} else if typeTag == "varint" {
		result.varint = true
	}

	optionalTag := fieldType.Tag.Get("optional")
	if optionalTag != "" {
		result.optional = true
		optionalSplit := strings.Split(optionalTag, "=")
		result.optionalField = optionalSplit[0]

		for i := 0; i < len(optionalSplit[1:]); {
			optionalRequiredValue, err := strconv.ParseInt(optionalSplit[i], 10, 32)
			if err != nil {
				panic(err)
			}
			result.optionalFieldValues = append(result.optionalFieldValues, optionalRequiredValue)
		}
	}

	sizeTag := fieldType.Tag.Get("size")
	if sizeTag == "infer" {
		result.inferSize = true
	}
	return result
}

func ptrEncoder(v reflect.Value) []byte {
	if v.IsNil() {
		return []byte{}
	}
	return typeEncoder(v.Elem().Type(), false)(v.Elem())
}
