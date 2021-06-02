package entity

import (
	_ "embed"
	"encoding/json"
	"errors"
)

//go:embed "data/entities.json"
var entityData []byte

var entityStorage []EntityType

func Load() (err error) {
	if entityStorage != nil {
		return errors.New("entity data already loaded")
	}

	entityStorage = make([]EntityType, 0)
	err = json.Unmarshal(entityData, &entityStorage)
	return
}

func ByName(name string) *EntityType {
	for i, e := range entityStorage {
		if e.Name == name {
			return &entityStorage[i]
		}
	}
	panic(errors.New("unknown entity type " + name))
}
