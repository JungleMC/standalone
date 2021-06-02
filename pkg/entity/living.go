package entity

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type LivingEntity struct {
	entityType    *Type
	entityId      int32
	uuid          uuid.UUID
	X             float64
	Y             float64
	Z             float64
	HeadPitch     byte
	HeadYaw       byte
	Pitch         byte
	Yaw           byte
	tickProcessor func(e *LivingEntity, time time.Time) error
	shouldRemove  bool
	wait          *sync.WaitGroup
}

func NewLivingEntity(entityType *Type, uuid uuid.UUID, processor func(e *LivingEntity, time time.Time) error) *LivingEntity {
	return &LivingEntity{
		entityType:    entityType,
		uuid:          uuid,
		tickProcessor: processor,
		wait:          &sync.WaitGroup{},
	}
}

func (e *LivingEntity) Type() *Type {
	return e.entityType
}

func (e *LivingEntity) ID() int32 {
	return e.entityId
}

func (e *LivingEntity) UUID() uuid.UUID {
	return e.uuid
}

func (e *LivingEntity) id(id int32) {
	e.entityId = id
}

func (e *LivingEntity) tick(time time.Time) {
	if e.tickProcessor != nil {
		e.tickProcessor(e, time)
	}
}

func (e *LivingEntity) waitGroup() *sync.WaitGroup {
	return e.wait
}

func (e *LivingEntity) remove() bool {
	return e.shouldRemove
}

func (e *LivingEntity) IsDead() bool {
	return e.shouldRemove
}

func (e *LivingEntity) Kill() {
	e.shouldRemove = true
}
