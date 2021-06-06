package event

import (
    "log"
    "reflect"
)

type ServerLoadEvent struct {
}

type ServerLoadListener struct {
}

func (l ServerLoadListener) OnEvent(event Event) {
    log.Println("Test: " + reflect.TypeOf(event).Name())
}