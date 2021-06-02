package item

import (
	_ "embed"
	"encoding/json"
	"errors"
)

//go:embed "items.json"
var itemData []byte

//go:embed "tags.json"
var tagData []byte

var items []Type
var tags []Tag

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
