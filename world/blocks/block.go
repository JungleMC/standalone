package blocks

type Block struct {
	Id          uint   `json:"id"`
	Material    string `json:"material,omitempty"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`

	MinStateId     uint         `json:"minStateId,omitempty"`
	MaxStateId     uint         `json:"maxStateId,omitempty"`
	States         []BlockState `json:"states"`
	DefaultStateId uint         `json:"defaultState,omitempty"`

	StackSize   uint `json:"stackSize"`
	CanDig      bool `json:"diggable"`
	Transparent bool `json:"transparent"`
	EmitLight   uint `json:"emitLight"`
	FilterLight uint `json:"filterLight"`

	Resistance   float32         `json:"resistance,omitempty"`
	HarvestTools map[string]bool `json:"harvestTools"`
	Drops        []uint          `json:"drops"`
	BoundingBox  string          `json:"boundingBox"`
}
