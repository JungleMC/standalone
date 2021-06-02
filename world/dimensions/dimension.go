package dimensions

import (
	"github.com/junglemc/JungleTree"
)

type DimensionStorage struct {
	Type    string           `json:"type" nbt:"type"`
	Entries []DimensionEntry `json:"value" nbt:"value"`
}

type DimensionEntry struct {
	Name      JungleTree.Identifier `json:"name" nbt:"name"`
	ID        int32                 `json:"id" nbt:"id"`
	Dimension Dimension             `json:"element" nbt:"element"`
}

type Dimension struct {
	PiglinSafe         bool                  `json:"piglin_safe" nbt:"piglin_safe"`
	Natural            bool                  `json:"natural" nbt:"natural"`
	AmbientLight       float32               `json:"ambient_light" nbt:"ambient_light"`
	IsFixedTime        bool                  `json:"is_fixed_time" nbt:"-"`
	FixedTime          int64                 `json:"fixed_time,omitempty" nbt:"fixed_time" optional:"IsFixedTime""`
	Infiburn           JungleTree.Identifier `json:"infiburn" nbt:"infiniburn"`
	RespawnAnchorWorks bool                  `json:"respawn_anchor_works" nbt:"respawn_anchor_works"`
	HasSkylight        bool                  `json:"has_skylight" nbt:"has_skylight"`
	BedWorks           bool                  `json:"bed_works" nbt:"bed_works"`
	Effects            JungleTree.Identifier `json:"effects" nbt:"effects"`
	HasRaids           bool                  `json:"has_raids" nbt:"has_raids"`
	LogicalHeight      int32                 `json:"logical_height" nbt:"logical_height"`
	CoodinateScale     float32               `json:"coordinate_scale" nbt:"coordinate_scale"`
	Ultrawarm          bool                  `json:"ultrawarm" nbt:"ultrawarm"`
	HasCeiling         bool                  `json:"has_ceiling" nbt:"has_ceiling"`
}
