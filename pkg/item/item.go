package item

import (
	_ "embed"
	"encoding/json"
	"errors"

	. "github.com/junglemc/JungleTree/pkg/util"
)

//go:embed "items.json"
var itemData []byte

//go:embed "tags.json"
var tagData []byte

var (
	items []Type
	tags  []Tag
)

func Load() (err error) {
	if items != nil {
		return errors.New("item data already loaded")
	}

	items = make([]Type, 0)
	err = json.Unmarshal(itemData, &items)
	if err != nil {
		return
	}

	if tags != nil {
		return errors.New("item tags already loaded")
	}

	tags = make([]Tag, 0)
	return json.Unmarshal(tagData, &tags)
}

func (i *Type) Empty() bool {
	return i.Name.Equals("minecraft:air")
}

func (i *Type) tag() (t *Tag, ok bool) {
	for x, tag := range tags {
		for _, v := range tag.Values {
			if v == i.Name {
				return &tags[x], true
			}
		}
	}
	return nil, false
}

type Type struct {
	ID          int32      `json:"id"`
	Name        Identifier `json:"name"`
	DisplayName string     `json:"displayName"`
	StackSize   byte       `json:"stackSize"`
}

type Tag struct {
	Name    Identifier
	Replace bool
	Values  []Identifier
}

func ByName(ident Identifier) (itemType *Type, ok bool) {
	for i, item := range items {
		if item.Name.Equals(ident) {
			return &items[i], true
		}
	}
	return nil, false
}

func ByTag(tag Identifier) []*Type {
	result := make([]*Type, 0, 0)
	for i, itm := range items {
		t, ok := itm.tag()
		if ok && t.Name.Equals(tag) {
			result = append(result, &items[i])
		}
	}
	return result
}
