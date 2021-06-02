package block

import (
	_ "embed"
	"encoding/json"
	"errors"
)

//go:embed "blocks.json"
var data []byte

var blocks []Block

func Load() (err error) {
	if blocks != nil {
		return errors.New("block data already loaded")
	}

	blocks = make([]Block, 0)
	err = json.Unmarshal(data, &blocks)
	return
}
