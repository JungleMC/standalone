package nbt

import (
	"bytes"
	"reflect"
)

func Unmarshal(data []byte, value reflect.Value) (tagName string, err error) {
	buf := bytes.NewBuffer(data)
	tagType, err := readTagType(buf)
	if err != nil {
		return
	}

	if tagType == tagEnd {
		return
	}

	tagName, err = readString(buf)
	if err != nil {
		return
	}

	err = readValue(buf, tagType, value)
	if err != nil {
		return
	}
	return
}

func readValue(buf *bytes.Buffer, tagType namedTagType, v reflect.Value) error {
	switch tagType {
	case tagByte:
		return readTagByte(buf, v)
	case tagShort:
		return readTagShort(buf, v)
	case tagInt:
		return readTagInt(buf, v)
	case tagLong:
		return readTagLong(buf, v)
	case tagFloat:
		return readTagFloat(buf, v)
	case tagDouble:
		return readTagDouble(buf, v)
	case tagString:
		return readTagString(buf, v)
	case tagList:
		return readTagList(buf, v)
	case tagCompound:
		switch v.Kind() {
		case reflect.Struct:
			return readTagCompoundStruct(buf, v)
		case reflect.Map:
			return readTagCompoundMap(buf, v)
		}
	case tagByteArray:
		return readTagByteArray(buf, v)
	case tagIntArray:
		return readTagIntArray(buf, v)
	case tagLongArray:
		return readTagLongArray(buf, v)
	}
	return nil
}
