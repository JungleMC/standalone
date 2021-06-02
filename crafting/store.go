package crafting

import (
	_ "embed"
	"encoding/json"
	"errors"
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
