package world

import (
	"github.com/junglemc/JungleTree/world/biomes"
	"github.com/junglemc/JungleTree/world/dimensions"
)

type DimensionCodec struct {
	Dimensions dimensions.DimensionStorage `json:"minecraft:dimension_type" nbt:"minecraft:dimension_type"`
	Biomes     biomes.BiomeStorage         `json:"minecraft:worldgen/biome" nbt:"minecraft:worldgen/biome"`
}

func Codec() DimensionCodec {
	return DimensionCodec{
		Dimensions: dimensions.Store(),
		Biomes:     biomes.Store(),
	}
}
