package util

import (
	"bytes"
	"strings"

	. "github.com/junglemc/JungleTree/pkg/codec"
)

type Identifier string

func (i Identifier) Equals(other Identifier) bool {
	return i == other
}

func (i Identifier) Prefix() string {
	dataSplit := strings.Split(string(i), ":")
	if len(dataSplit) > 1 {
		return dataSplit[0]
	}
	return ""
}

func (i Identifier) Name() string {
	dataSplit := strings.Split(string(i), ":")
	if len(dataSplit) == 2 {
		return dataSplit[1]
	}
	return ""
}

func (i Identifier) Empty() bool {
	return i == ""
}

func (i Identifier) String() string {
	return string(i)
}
func (i *Identifier) MarshalMinecraft() ([]byte, error) {
	return WriteString(string(*i)), nil
}

func (i *Identifier) UnmarshalMinecraft(buf *bytes.Buffer) error {
	*i = Identifier(ReadString(buf))
	return nil
}
