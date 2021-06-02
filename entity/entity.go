package entity

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type EntityType struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	DisplayName string  `json:"displayName"`
	Width       float32 `json:"width"`
	Height      float32 `json:"height"`
}

type Entity interface {
	Type() *EntityType
	ID() int32
	UUID() uuid.UUID
	id(id int32)
	tick(time time.Time)
	waitGroup() *sync.WaitGroup
	remove() bool
}
