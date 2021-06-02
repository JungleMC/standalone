package command

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/junglemc/JungleTree/pkg/codec"
	. "github.com/junglemc/JungleTree/pkg/util"
)

//go:embed "tags.json"
var data []byte

var tags *Tags

func Load() error {
	return json.Unmarshal(data, tags)
}

type Tags struct {
	BlockTags  map[string]*Tag `json:"block"`
	ItemTags   map[string]*Tag `json:"item"`
	FluidTags  map[string]*Tag `json:"fluid"`
	EntityTags map[string]*Tag `json:"entity_types"`
}

type Tag struct {
	Replace bool          `json:"replace,omitempty"`
	Values  []*Identifier `json:"values"`
}

type TagType int

const (
	TagTypeBlock TagType = iota
	TagTypeItem
	TagTypeFluid
	TagTypeEntity
)

func (t *Tags) MarshalMinecraft() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write(writeTags(t.BlockTags, TagTypeBlock))
	buf.Write(writeTags(t.ItemTags, TagTypeItem))
	buf.Write(writeTags(t.FluidTags, TagTypeFluid))
	buf.Write(writeTags(t.EntityTags, TagTypeEntity))
	return buf.Bytes(), nil
}

func writeTags(tags map[string]*Tag, tagType TagType) []byte {
	buf := &bytes.Buffer{}
	buf.Write(codec.WriteVarInt32(int32(len(tags))))
	for k, v := range tags {
		buf.Write(codec.WriteString(k))
		buf.Write(codec.WriteVarInt32(int32(len(v.Values))))
		// TODO: Write data
	}
	return buf.Bytes()
}
