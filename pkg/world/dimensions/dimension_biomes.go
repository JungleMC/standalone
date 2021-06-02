package dimensions

import (
	"github.com/junglemc/JungleTree/pkg/world/biome"
)

type DimensionBiome struct {
	Dimensions *Dimensions `json:"minecraft:dimension_type" nbt:"minecraft:dimension_type"`
	Biomes     *biome.Biomes          `json:"minecraft:worldgen/biome" nbt:"minecraft:worldgen/biome"`
}

func DimensionBiomes() DimensionBiome {
	return DimensionBiome{
		Dimensions: Store(),
		Biomes:     biome.Store(),
	}
}
