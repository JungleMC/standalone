package biome

import (
	_ "embed"
	"encoding/json"

	"github.com/junglemc/JungleTree/pkg/level/dimensions"
	. "github.com/junglemc/JungleTree/pkg/util"
)

//go:embed "biomes.json"
var data []byte

var biomes Biomes

func Load() (err error) {
	return json.Unmarshal(data, &biomes)
}

func Store() Biomes {
	return biomes
}

func ByName(name string) *Biome {
	for i, biome := range biomes.Entries {
		if biome.Name == name {
			return &biomes.Entries[i].Element
		}
	}
	return nil
}

type Biomes struct {
	Type    string  `json:"type" nbt:"type"`
	Entries []Entry `json:"value" nbt:"value"`
}

type Entry struct {
	Name    string `json:"name" nbt:"name"`
	ID      int32  `json:"id" nbt:"id"`
	Element Biome  `json:"element" nbt:"element"`
}

type Biome struct {
	ID                     int32      `json:"id" nbt:"-"`
	Name                   string     `json:"name" nbt:"-"`
	Category               string     `json:"category" nbt:"category"`
	Temperature            float32    `json:"temperature" nbt:"temperature"`
	Precipitation          string     `json:"precipitation" nbt:"precipitation"`
	Depth                  float32    `json:"depth" nbt:"depth"`
	Scale                  float32    `json:"scale" nbt:"scale"`
	Dimension              Identifier `json:"dimension" nbt:"-"`
	DisplayName            string     `json:"displayName" nbt:"-"`
	Color                  int32      `json:"color" nbt:"-"`
	Rainfall               float32    `json:"rainfall" nbt:"downfall"`
	Parent                 string     `json:"parent,omitempty" nbt:"-"`
	ChildID                int32      `json:"child,omitempty" nbt:"-"`
	Climates               []Climate  `json:"climates,omitempty" nbt:"-"`
	HasTemperatureModifier bool       `json:"has_temperature_modifier" nbt:"-"`
	TemperatureModifier    string     `json:"temperature_modifier,omitempty" nbt:"temperature_modifier" optional:"HasTemperatureModifier"`
	Effects                Effect     `json:"effects,omitempty" nbt:"effects"`
	HasParticle            bool       `json:"has_particle" nbt:"-"`
	Particle               Particle   `json:"particle,omitempty" nbt:"particle" optional:"HasParticle"`
}

type Climate struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Altitude    float32 `json:"altitude"`
	Weirdness   float32 `json:"weirdness"`
	Offset      float32 `json:"offset"`
}

type Effect struct {
	SkyColor              int32          `json:"sky_color" nbt:"sky_color"`
	WaterFogColor         int32          `json:"water_fog_color" nbt:"water_fog_color"`
	FogColor              int32          `json:"fog_color" nbt:"fog_color"`
	WaterColor            int32          `json:"water_color" nbt:"water_color"`
	HasFoliageColor       bool           `json:"has_foliage_color,omitempty" nbt:"-"`
	FoliageColor          int32          `json:"foliage_color,omitempty" nbt:"foliage_color" optional:"HasFoliageColor"`
	HasGrassColor         bool           `json:"has_grass_color,omitempty" nbt:"-"`
	GrassColor            int32          `json:"grass_color,omitempty" nbt:"grass_color" optional:"HasGrassColor"`
	HasGrassColorModifier bool           `json:"has_grass_color_modifier,omitempty" nbt:"-"`
	GrassColorModifier    string         `json:"grass_color_modifier" nbt:"grass_color_modifier" optional:"HasGrassColorModifier"`
	HasMusic              bool           `json:"has_music,omitempty" nbt:"-"`
	Music                 Music          `json:"music,omitempty" nbt:"music" optional:"HasMusic"`
	HasAmbientSound       bool           `json:"has_ambient_sound,omitempty" nbt:"-"`
	AmbientSound          string         `json:"ambient_sound,omitempty" nbt:"ambient_sound" optional:"HasAmbientSound"`
	HasAdditionsSound     bool           `json:"has_additions_sound,omitempty" nbt:"-"`
	AdditionsSound        AdditionsSound `json:"additions_sound,omitempty" nbt:"additions_sound" optional:"HasAdditionsSound"`
	HasMoodSound          bool           `json:"has_mood_sound" nbt:"-"`
	MoodSound             MoodSound      `json:"mood_sound" nbt:"mood_sound" optional:"HasMoodSound"`
}

type Music struct {
	ReplaceCurrentMusic bool       `json:"replace_current_music" nbt:"replace_current_music"`
	Sound               Identifier `json:"sound" nbt:"sound"`
	MaxDelay            int32      `json:"max_delay" nbt:"max_delay"`
	MinDelay            int32      `json:"min_delay" nbt:"min_delay"`
}

type AdditionsSound struct {
	Sound      string  `json:"sound" nbt:"sound"`
	TickChance float64 `json:"tick_chance" nbt:"tick_chance"`
}

type MoodSound struct {
	Sound             string  `json:"sound" nbt:"sound"`
	TickDelay         int32   `json:"tick_delay" nbt:"tick_delay"`
	Offset            float64 `json:"offset" nbt:"offset"`
	BlockSearchExtent int32   `json:"block_search_extent" nbt:"block_search_extent"`
}

type Particle struct {
	Probability float32         `json:"probability" nbt:"probability"`
	Options     ParticleOptions `json:"options" nbt:"options"`
}

type ParticleOptions struct {
	Type string `json:"type" nbt:"type"`
}

func (b *Biome) DimensionType() *dimensions.Dimension {
	dimension, _ := dimensions.ByName(b.Dimension)
	return dimension
}
