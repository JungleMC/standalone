package world

import (
	"github.com/junglemc/JungleTree/pkg/world/biome"
	"github.com/junglemc/JungleTree/pkg/world/dimensions"
)

type DimensionBiome struct {
	Dimensions dimensions.Dimensions `json:"minecraft:dimension_type" nbt:"minecraft:dimension_type"`
	Biomes     biome.Biomes          `json:"minecraft:worldgen/biome" nbt:"minecraft:worldgen/biome"`
}

func DimensionBiomes() DimensionBiome {
	return DimensionBiome{
		Dimensions: dimensions.Store(),
		Biomes:     biome.Store(),
	}
}
