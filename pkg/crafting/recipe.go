package crafting

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	codec2 "github.com/junglemc/JungleTree/internal/net/codec"
	"github.com/junglemc/JungleTree/item"
	"github.com/junglemc/JungleTree/pkg/inventory"
	. "github.com/junglemc/JungleTree/pkg/codec"
	. "github.com/junglemc/JungleTree/pkg/util"
	"reflect"
)

//go:embed "recipes.json"
var recipeData []byte

var recipes []*Recipe

func Load() (err error) {
	if recipes != nil {
		return errors.New("recipe data already loaded")
	}

	recipes = make([]*Recipe, 0)
	err = json.Unmarshal(recipeData, &recipes)
	if err != nil {
		return
	}

	for x, v := range recipes {
		if v.Ingredients == nil {
			continue
		}

		for y, ingredient := range v.Ingredients {
			if ingredient.Count == 0 {
				recipes[x].Ingredients[y].Count = 1
			}
		}
	}
	return
}

type Recipe struct {
	ID          Identifier
	Group       Identifier
	Ingredients []Ingredient           `json:"ingredients,omitempty"`
	Ingredient  Ingredient             `json:"ingredient,omitempty"`
	Base        Ingredient             `json:"base,omitempty"`
	Addition    Ingredient             `json:"addition,omitempty"`
	Pattern     []string               `json:"pattern,omitempty"`
	Key         map[string]interface{} `json:"key,omitempty"`
	Result      interface{}            `json:"result,omitempty"`
	Type        Identifier             `json:"type"`
	Experience  float32                `json:"experience,omitempty"`
	CookingTime int32                  `json:"cookingtime,omitempty"`
}

type Ingredient struct {
	Count int32      `type:"varint"`
	Item  Identifier `json:"item,omitempty"`
	Tag   Identifier `json:"tag,omitempty"`
}

func Recipes() []*Recipe {
	return recipes
}

func (v *Recipe) MarshalMinecraft() ([]byte, error) {
	buf := &bytes.Buffer{}

	buf.Write(WriteString(string(v.Type)))
	buf.Write(WriteString(string(v.ID)))

	switch v.Type.Name() {
	case "crafting_shapeless":
		buf.Write(v.writeShapeless())
		break
	case "crafting_shaped":
		buf.Write(v.writeShaped())
		break
	case "smelting", "blasting", "smoking", "campfire_cooking":
		buf.Write(v.writeCooking())
		break
	case "stonecutting":
		buf.Write(v.writeStonecutting())
		break
	case "smithing":
		buf.Write(v.writeSmithing())
		break
	}
	return buf.Bytes(), nil
}

// TODO: Unmarshal

func (v *Recipe) writeShapeless() []byte {
	buf := &bytes.Buffer{}
	buf.Write(WriteString(string(v.Group)))
	buf.Write(writeShapelessIngredients(v.Ingredients...))
	buf.Write(writeResult(v.Result))
	return buf.Bytes()
}

func (v *Recipe) writeShaped() []byte {
	buf := &bytes.Buffer{}
	buf.Write(WriteVarInt32(v.shapedWidth()))
	buf.Write(WriteVarInt32(v.shapedHeight()))
	buf.Write(WriteString(string(v.Group)))
	buf.Write(v.writeShapedIngredients())
	buf.Write(writeResult(v.Result))
	return buf.Bytes()
}

func (v *Recipe) writeCooking() []byte {
	buf := &bytes.Buffer{}
	buf.Write(WriteString(string(v.Group)))

	if v.Ingredients != nil {
		buf.Write(WriteVarInt32(int32(len(v.Ingredients))))

		for _, ingredient := range v.Ingredients {
			itemType, ok := item.ByName(ingredient.Item)
			if !ok {
				panic("failed to find item type")
			}
			buf.Write(writeSlot(itemType, 1))
		}
	} else {
		buf.Write(writeIngredient(v.Ingredient))
	}

	buf.Write(writeResult(v.Result))
	buf.Write(WriteFloat32(v.Experience))
	buf.Write(WriteVarInt32(v.CookingTime))

	return buf.Bytes()
}

func (v *Recipe) writeStonecutting() []byte {
	buf := &bytes.Buffer{}
	buf.Write(WriteString(string(v.Group)))
	if v.Ingredients != nil {
		buf.Write(WriteVarInt32(int32(len(v.Ingredients))))

		for _, ingredient := range v.Ingredients {
			itemType, ok := item.ByName(ingredient.Item)
			if !ok {
				panic("failed to find item type")
			}
			buf.Write(writeSlot(itemType, 1))
		}
	} else {
		buf.Write(writeIngredient(v.Ingredient))
	}
	buf.Write(writeResult(v.Result))
	return buf.Bytes()
}

func (v *Recipe) writeSmithing() []byte {
	buf := &bytes.Buffer{}
	buf.Write(writeIngredient(v.Base))
	buf.Write(writeIngredient(v.Addition))
	buf.Write(writeResult(v.Result))
	return buf.Bytes()
}

func writeShapelessIngredients(ingredients ...Ingredient) []byte {
	buf := &bytes.Buffer{}
	buf.Write(WriteVarInt32(int32(len(ingredients))))
	for _, ingredient := range ingredients {
		buf.Write(writeIngredient(ingredient))
	}
	return buf.Bytes()
}

func (v *Recipe) writeShapedIngredients() []byte {
	buf := &bytes.Buffer{}
	for _, line := range v.Pattern {
		for _, key := range line {
			var ingredient Ingredient
			if string(key) == " " {
				ingredient = Ingredient{
					Count: 1,
					Item:  "minecraft:air",
				}
			} else {
				ingredient = v.shapedKeyLookup(string(key))
			}
			buf.Write(writeIngredient(ingredient))
		}
	}
	return buf.Bytes()
}

func writeIngredient(ingredient Ingredient) []byte {
	buf := &bytes.Buffer{}

	if ingredient.Tag.Empty() {
		buf.Write(WriteVarInt32(1))

		itemType, ok := item.ByName(ingredient.Item)
		if !ok {
			panic("failed to find item type")
		}
		buf.Write(writeSlot(itemType, 1))
	} else {
		items := item.ByTag(ingredient.Tag)
		buf.Write(WriteVarInt32(int32(len(items))))

		for _, itm := range items {
			buf.Write(writeSlot(itm, 1))
		}
	}
	return buf.Bytes()
}

func writeSlot(i *item.Type, count byte) []byte {
	slot := inventory.Slot{
		Present: !i.Empty(),
		ID:      i.ID,
		Count:   count,
		Data:    make(map[string]interface{}),
	}
	return codec2.Marshal(slot)
}

func writeResult(v interface{}) []byte {
	switch reflect.TypeOf(v).Kind() {
	case reflect.String:
		itemType, ok := item.ByName(Identifier(v.(string)))
		if !ok {
			panic("failed to find item type")
		}
		return writeSlot(itemType, 1)
	case reflect.Map:
		vMap := v.(map[string]interface{})

		itemType, ok := item.ByName(Identifier(vMap["item"].(string)))
		if !ok {
			panic("failed to find item type")
		}

		count := 1
		if vMap["count"] != nil {
			count = int(vMap["count"].(float64))
		}

		return writeSlot(itemType, byte(count))
	}
	panic("result type unknown: " + reflect.TypeOf(v).Kind().String())
}

func (v *Recipe) shapedWidth() int32 {
	largest := 0
	for _, v := range v.Pattern {
		if len(v) > largest {
			largest = len(v)
		}
	}
	return int32(largest)
}

func (v *Recipe) shapedHeight() int32 {
	return int32(len(v.Pattern))
}

func (v *Recipe) shapedKeyLookup(key string) Ingredient {
	for k, v := range v.Key {
		if k != key {
			continue
		}

		if reflect.TypeOf(v).Kind() == reflect.Slice {
			val := v.([]interface{})

			for _, valInner := range val {
				value := valInner.(map[string]interface{})
				result := Ingredient{Count: 1}

				if value["count"] != nil {
					result.Count = value["count"].(int32)
				}

				if value["item"] != nil {
					result.Item = Identifier(value["item"].(string))
				}

				if value["tag"] != nil {
					result.Tag = Identifier(value["tag"].(string))
				}
				return result
			}
			panic("invalid ingredient slice")
		}

		if reflect.TypeOf(v).Kind() == reflect.Map {
			value := v.(map[string]interface{})
			result := Ingredient{Count: 1}

			if value["count"] != nil {
				result.Count = value["count"].(int32)
			}

			if value["item"] != nil {
				result.Item = Identifier(value["item"].(string))
			}

			if value["tag"] != nil {
				result.Tag = Identifier(value["tag"].(string))
			}

			return result
		}

		if reflect.TypeOf(v).Kind() == reflect.String {
			return Ingredient{
				Count: 1,
				Item:  Identifier(v.(string)),
			}
		}
	}
	panic("shaped key lookup not found: " + key)
}
