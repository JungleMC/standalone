package block

import "errors"

type Block struct {
	Id          uint   `json:"id"`
	Material    string `json:"material,omitempty"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`

	MinStateId     uint    `json:"minStateId,omitempty"`
	MaxStateId     uint    `json:"maxStateId,omitempty"`
	States         []State `json:"states"`
	DefaultStateId uint    `json:"defaultState,omitempty"`

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

type State struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	NumValues uint     `json:"num_values"`
	Values    []string `json:"values"`
}

func Empty() *Block {
	return ByName("air")
}

func ByName(name string) *Block {
	for i, block := range blocks {
		if block.Name == name {
			return &blocks[i]
		}
	}
	if name == "air" {
		panic(errors.New("default could not be set: air not present in block data"))
	}
	return Empty()
}
