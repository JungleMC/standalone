package biomes

import (
	_ "embed"
	"encoding/json"
	"errors"
)

//go:embed "biomes.json"
var biomeData []byte

var biomeStorage BiomeStorage

func Load() (err error) {
	if biomeStorage.Entries != nil {
		return errors.New("biome data already loaded")
	}

	biomeStorage = BiomeStorage{}
	err = json.Unmarshal(biomeData, &biomeStorage)
	return
}

func Store() BiomeStorage {
	return biomeStorage
}

func ByName(name string) *Biome {
	for i, biome := range biomeStorage.Entries {
		if biome.Name == name {
			return &biomeStorage.Entries[i].Element
		}
	}
	if name == "ocean" {
		panic(errors.New("default could not be set: ocean not present in biome data"))
	}
	return ByName("ocean")
}
