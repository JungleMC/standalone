package item

import (
	. "github.com/junglemc/JungleTree/pkg/util"
)

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
