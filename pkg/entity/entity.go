package entity

import (
	_ "embed"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

//go:embed "entities.json"
var data []byte

var entityTypes []Type

func Load() (err error) {
	if entityTypes != nil {
		return errors.New("entity data already loaded")
	}

	entityTypes = make([]Type, 0)
	err = json.Unmarshal(data, &entityTypes)
	return
}

type Type struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	DisplayName string  `json:"displayName"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
}

type Entity interface {
	Type() *Type
	ID() int32
	UUID() uuid.UUID
	id(id int32)
	tick(time time.Time)
	waitGroup() *sync.WaitGroup
	remove() bool
}

func ByName(name string) *Type {
	for i, e := range entityTypes {
		if e.Name == name {
			return &entityTypes[i]
		}
	}
	panic(errors.New("unknown entity type " + name))
}
