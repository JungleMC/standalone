package dimensions

import (
	_ "embed"
	"encoding/json"
	. "github.com/junglemc/JungleTree/pkg/util"
)

//go:embed "dimensions.json"
var data []byte

var dimensions Dimensions

func Load() (err error) {
	return json.Unmarshal(data, &dimensions)
}

func Store() Dimensions {
	return dimensions
}

func ByName(name Identifier) (dimension *Dimension, ok bool) {
	for i, d := range dimensions.Entries {
		if d.Name.Equals(name) {
			return &dimensions.Entries[i].Dimension, true
		}
	}
	return nil, false
}

type Dimensions struct {
	Type    string           `json:"type" nbt:"type"`
	Entries []DimensionEntry `json:"value" nbt:"value"`
}

type DimensionEntry struct {
	Name      Identifier `json:"name" nbt:"name"`
	ID        int32      `json:"id" nbt:"id"`
	Dimension Dimension  `json:"element" nbt:"element"`
}

type Dimension struct {
	PiglinSafe         bool       `json:"piglin_safe" nbt:"piglin_safe"`
	Natural            bool       `json:"natural" nbt:"natural"`
	AmbientLight       float32    `json:"ambient_light" nbt:"ambient_light"`
	IsFixedTime        bool       `json:"is_fixed_time" nbt:"-"`
	FixedTime          int64      `json:"fixed_time,omitempty" nbt:"fixed_time" optional:"IsFixedTime""`
	Infiburn           Identifier `json:"infiburn" nbt:"infiniburn"`
	RespawnAnchorWorks bool       `json:"respawn_anchor_works" nbt:"respawn_anchor_works"`
	HasSkylight        bool       `json:"has_skylight" nbt:"has_skylight"`
	BedWorks           bool       `json:"bed_works" nbt:"bed_works"`
	Effects            Identifier `json:"effects" nbt:"effects"`
	HasRaids           bool       `json:"has_raids" nbt:"has_raids"`
	LogicalHeight      int32      `json:"logical_height" nbt:"logical_height"`
	CoordinateScale    float32    `json:"coordinate_scale" nbt:"coordinate_scale"`
	Ultrawarm          bool       `json:"ultrawarm" nbt:"ultrawarm"`
	HasCeiling         bool       `json:"has_ceiling" nbt:"has_ceiling"`
}
