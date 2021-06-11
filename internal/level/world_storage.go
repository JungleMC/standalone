package level

import (
	"math/rand"
	"os"

	"github.com/junglemc/JungleTree/internal/storage"
	world2 "github.com/junglemc/JungleTree/pkg/level"
)

func Load() error {
	ok, err := storage.Has("jungletree:worlds", nil)
	if err != nil {
		panic(err)
	}

	if !ok {
		// TODO: Environment variable wrap default values - 12 factor apps style
		name := os.Getenv("DEFAULT_WORLD_NAME")
		if name == "" {
			name = "world"
		}

		world := world2.NewWorld(name, rand.Uint64(), 256)
		err = storage.Put("jungletree:default_world", world.Name, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
