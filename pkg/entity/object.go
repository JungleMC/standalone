package entity

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type ObjectEntity struct {
	entityType    *Type
	entityId      int32
	uuid          uuid.UUID
	X             float64
	Y             float64
	Z             float64
	Pitch         byte
	Yaw           byte
	VelocityX     int16
	VelocityY     int16
	VelocityZ     int16
	ObjectData    int32
	tickProcessor func(e *ObjectEntity, time time.Time) error
	shouldRemove  bool
	wait          *sync.WaitGroup
}

func NewObjectEntity(entityType *Type, entityId int32, uuid uuid.UUID, processor func(e *ObjectEntity, time time.Time) error) *ObjectEntity {
	return &ObjectEntity{
		entityType:    entityType,
		entityId:      entityId,
		uuid:          uuid,
		tickProcessor: processor,
		wait:          &sync.WaitGroup{},
	}
}

func (e *ObjectEntity) Type() *Type {
	return e.entityType
}

func (e *ObjectEntity) ID() int32 {
	return e.entityId
}

func (e *ObjectEntity) UUID() uuid.UUID {
	return e.uuid
}

func (e *ObjectEntity) id(id int32) {
	e.entityId = id
}

func (e *ObjectEntity) tick(time time.Time) {
	if e.tickProcessor != nil {
		e.tickProcessor(e, time)
	}
}

func (e *ObjectEntity) waitGroup() *sync.WaitGroup {
	return e.wait
}

func (e *ObjectEntity) remove() bool {
	return e.shouldRemove
}

func (e *ObjectEntity) Remove() {
	e.shouldRemove = true
}
