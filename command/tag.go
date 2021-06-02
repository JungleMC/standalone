package command

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"github.com/junglemc/JungleTree"
	. "github.com/junglemc/JungleTree/net/codec"
)

type Tags struct {
	BlockTags  map[string]*Tag `json:"block"`
	ItemTags   map[string]*Tag `json:"item"`
	FluidTags  map[string]*Tag `json:"fluid"`
	EntityTags map[string]*Tag `json:"entity_types"`
}

type Tag struct {
	Replace bool                     `json:"replace,omitempty"`
	Values  []*JungleTree.Identifier `json:"values"`
}

type TagType int

const (
	TagTypeBlock TagType = iota
	TagTypeItem
	TagTypeFluid
	TagTypeEntity
)

//go:embed "tags.json"
var tagData []byte

var tags Tags

func Load() error {
	if tags.BlockTags != nil {
		return errors.New("tag data already loaded")
	}
	return json.Unmarshal(tagData, &tags)
}

func Store() Tags {
	return tags
}

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
	buf.Write(WriteVarInt32(int32(len(tags))))
	for k, v := range tags {
		buf.Write(WriteString(k))
		buf.Write(WriteVarInt32(int32(len(v.Values))))
		// TODO: Write data
	}
	return buf.Bytes()
}
