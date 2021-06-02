package entity

import (
	"log"
	"sync"
	"sync/atomic"
	"time"
)

type EntityRunner struct {
	running      bool
	TPS          int
	ticker       *time.Ticker
	stop         chan bool
	entities     []Entity
	spawnQueue   chan Entity
	wait         *sync.WaitGroup
	nextEntityId int32
}

func (t *EntityRunner) Run() {
	if t.running {
		return
	}

	t.ticker = time.NewTicker(time.Second / time.Duration(t.TPS))
	t.stop = make(chan bool)
	t.spawnQueue = make(chan Entity)
	t.wait = &sync.WaitGroup{}
	t.running = true

	go func() {
		defer close(t.spawnQueue)

		t.entities = make([]Entity, 0)

		for {
			select {
			case <-t.stop:
				break
			case time := <-t.ticker.C:
				err := t.tick(time)
				if err != nil {
					log.Println(err)
					t.stop <- true
				}
			}
		}
	}()
}

func (t *EntityRunner) tick(time time.Time) (err error) {
	t.wait.Add(1)
	for entity := range t.spawnQueue {
		entity.waitGroup().Wait()
		entity.waitGroup().Add(1)
		entity.id(t.nextEntityId)
		t.entities = append(t.entities, entity)
		atomic.AddInt32(&t.nextEntityId, 1)
		entity.waitGroup().Done()
	}
	t.wait.Done()

	t.wait.Add(1)
	for i, entity := range t.entities {
		entity.waitGroup().Wait()
		entity.waitGroup().Add(1)
		entity.tick(time)

		if entity.remove() {
			t.entities = append(t.entities[:i], t.entities[i+1:]...)
		}

		entity.waitGroup().Done()
	}
	t.wait.Done()
	return
}
