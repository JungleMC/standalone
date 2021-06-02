package dimensions

import (
	_ "embed"
	"encoding/json"
	"errors"
	"github.com/junglemc/JungleTree"
)

//go:embed "dimensions.json"
var dimensionData []byte

var dimensionStorage DimensionStorage

func Load() (err error) {
	if dimensionStorage.Entries != nil {
		return errors.New("dimension data already loaded")
	}

	dimensionStorage = DimensionStorage{}
	err = json.Unmarshal(dimensionData, &dimensionStorage)
	return
}

func Store() DimensionStorage {
	return dimensionStorage
}

func ByName(name JungleTree.Identifier) (dimension *Dimension, ok bool) {
	for i, dimension := range dimensionStorage.Entries {
		if dimension.Name.Equals(name) {
			return &dimensionStorage.Entries[i].Dimension, true
		}
	}
	return nil, false
}
