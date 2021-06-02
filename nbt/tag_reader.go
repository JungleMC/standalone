package nbt

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
)

func readTagByte(buf *bytes.Buffer, v reflect.Value) (err error) {
	value, err := buf.ReadByte()
	if err != nil {
		return
	}

	switch kind := v.Kind(); kind {
	case reflect.Bool:
		if byte(value) == 1 {
			v.SetBool(true)
		} else {
			v.SetBool(false)
		}
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(value))
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(value))
		break
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
		break
	default:
		return errors.New("cannot parse tagByte as " + kind.String())
	}
	return
}

func readTagShort(buf *bytes.Buffer, v reflect.Value) (err error) {
	value, err := readInt16(buf)
	if err != nil {
		return
	}

	switch kind := v.Kind(); kind {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(int64(value))
	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(value))
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
	default:
		return errors.New("cannot parse tagShort as " + kind.String())
	}
	return
}

func readTagInt(buf *bytes.Buffer, v reflect.Value) (err error) {
	value, err := readInt32(buf)
	if err != nil {
		return
	}

	switch kind := v.Kind(); kind {
	case reflect.Int, reflect.Int32, reflect.Int64:
		v.SetInt(int64(value))
	case reflect.Uint, reflect.Uint32, reflect.Uint64:
		v.SetUint(uint64(value))
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
	default:
		return errors.New("cannot parse tagInt as " + kind.String())
	}
	return
}

func readTagLong(buf *bytes.Buffer, v reflect.Value) (err error) {
	value, err := readInt64(buf)
	if err != nil {
		return
	}

	switch kind := v.Kind(); kind {
	case reflect.Int, reflect.Int64:
		v.SetInt(value)
	case reflect.Uint, reflect.Uint64:
		v.SetUint(uint64(value))
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
	default:
		return errors.New("cannot parse tagLong as " + kind.String())
	}
	return
}

func readTagFloat(buf *bytes.Buffer, v reflect.Value) (err error) {
	value, err := readFloat32(buf)
	if err != nil {
		return
	}

	switch kind := v.Kind(); kind {
	case reflect.Float32:
		v.SetFloat(float64(value))
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
	default:
		return errors.New("cannot parse tagFloat as " + kind.String())
	}
	return
}

func readTagDouble(buf *bytes.Buffer, v reflect.Value) (err error) {
	value, err := readFloat64(buf)
	if err != nil {
		return
	}

	switch kind := v.Kind(); kind {
	case reflect.Float32, reflect.Float64:
		v.SetFloat(value)
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
	default:
		return errors.New("cannot parse tagDouble as " + kind.String())
	}
	return
}

func readTagString(buf *bytes.Buffer, v reflect.Value) (err error) {
	value, err := readString(buf)
	if err != nil {
		return
	}

	switch kind := v.Kind(); kind {
	case reflect.String:
		v.SetString(value)
		break
	case reflect.Interface:
		v.Set(reflect.ValueOf(value))
		break
	default:
		return errors.New("cannot parse tagString as " + kind.String())
	}
	return
}

func readTagList(buf *bytes.Buffer, v reflect.Value) (err error) {
	listType, err := readTagType(buf)
	if err != nil {
		return
	}

	length, err := readInt32(buf)
	if err != nil {
		return
	}

	if length < 0 {
		length = 0
	}

	switch v.Kind() {
	case reflect.Interface:
		v.Set(reflect.ValueOf(make([]interface{}, length)))
		break
	case reflect.Slice:
		v.Set(reflect.MakeSlice(v.Type(), int(length), int(length)))
		break
	case reflect.Array:
		if arrayLength := v.Len(); arrayLength < int(length) {
			return fmt.Errorf("size mismatch in tagList: want=%d, available=%d", arrayLength, length)
		}
	}

	for i := 0; i < int(length); i++ {
		err = readValue(buf, listType, v.Index(i))
		if err != nil {
			return
		}
	}
	return
}

func readTagCompoundStruct(buf *bytes.Buffer, v reflect.Value) (err error) {
	for {
		var cmpTagType namedTagType
		var cmpTagName string

		cmpTagType, err = readTagType(buf)
		if err != nil {
			return
		}

		if cmpTagType == tagEnd {
			break
		}

		cmpTagName, err = readString(buf)
		if err != nil {
			return
		}

		for i := 0; i < v.NumField(); i++ {
			f := v.Type().Field(i)
			tagName := f.Tag.Get("nbt")
			if tagName == "-" {
				continue
			}

			if tagName == "" {
				tagName = f.Name
			}

			if tagName == cmpTagName {
				err = readValue(buf, cmpTagType, v.Field(i))
				if err != nil {
					return
				}
				break
			}
		}
	}
	return
}

func readTagCompoundMap(buf *bytes.Buffer, v reflect.Value) (err error) {
	if v.Type().Key().Kind() != reflect.String {
		return errors.New("map key should be of type string")
	}

	if v.IsNil() {
		v.Set(reflect.MakeMap(v.Type()))
	}

	for {
		var cmpTagType namedTagType
		var cmpTagName string

		cmpTagType, err = readTagType(buf)
		if err != nil {
			return
		}

		if cmpTagType == tagEnd {
			break
		}

		cmpTagName, err = readString(buf)
		if err != nil {
			return
		}

		var val interface{}

		switch cmpTagType {
		case tagByte:
			val = byte(0)
			break
		case tagShort:
			val = int16(0)
			break
		case tagInt:
			val = int32(0)
			break
		case tagLong:
			val = int64(0)
			break
		case tagFloat:
			val = float32(0)
			break
		case tagDouble:
			val = float64(0)
			break
		case tagString:
			val = ""
			break
		case tagList:
			val = make([]interface{}, 0)
			break
		case tagCompound:
			val = make(map[string]interface{})
			break
		case tagByteArray:
			val = make([]byte, 0)
			break
		case tagIntArray:
			val = make([]int32, 0)
			break
		case tagLongArray:
			val = make([]int64, 0)
			break
		}

		err = readValue(buf, cmpTagType, reflect.ValueOf(&val).Elem())
		if err != nil {
			return err
		}
		v.SetMapIndex(reflect.ValueOf(cmpTagName), reflect.ValueOf(val))
	}
	return
}

func readTagByteArray(buf *bytes.Buffer, v reflect.Value) (err error) {
	b, err := readByteSlice(buf)
	if err != nil {
		return
	}

	if v.Type() == reflect.TypeOf(b) {
		v.SetBytes(b)
	} else {
		v.Set(reflect.ValueOf(b))
	}
	return
}

func readTagIntArray(buf *bytes.Buffer, v reflect.Value) (err error) {
	b, err := readInt32Slice(buf)
	if err != nil {
		return
	}
	v.Set(reflect.ValueOf(b))
	return
}

func readTagLongArray(buf *bytes.Buffer, v reflect.Value) (err error) {
	b, err := readInt64Slice(buf)
	if err != nil {
		return
	}
	v.Set(reflect.ValueOf(b))
	return
}
