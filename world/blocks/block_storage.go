package blocks

import (
	_ "embed"
	"encoding/json"
	"errors"
)

//go:embed "blocks.json"
var blockData []byte

var blockStorage []Block

func Load() (err error) {
	if blockStorage != nil {
		return errors.New("block data already loaded")
	}

	if err != nil {
		return
	}

	blockStorage = make([]Block, 0)
	err = json.Unmarshal(blockData, &blockStorage)
	return
}

func Empty() *Block {
	return ByName("air")
}

func ByName(name string) *Block {
	for i, block := range blockStorage {
		if block.Name == name {
			return &blockStorage[i]
		}
	}
	if name == "air" {
		panic(errors.New("default could not be set: air not present in block data"))
	}
	return Empty()
}
