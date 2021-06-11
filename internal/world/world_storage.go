package world

import (
	"math/rand"

	"github.com/junglemc/JungleTree/internal/storage"
	"github.com/junglemc/JungleTree/pkg/world/level"
)

func Load() error {
	ok, err := storage.Has("jungletree:worlds", nil)
	if err != nil {
		panic(err)
	}

	if !ok {
		level.NewWorld("minecraft:world", rand.Uint64(), 256)
	}
	return nil
}
