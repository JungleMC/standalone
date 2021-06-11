package event

import (
	"log"
	. "reflect"

	"github.com/junglemc/JungleTree/internal/configuration"
)

type Event interface {
	IsAsync() bool
}

type Cancellable interface {
	IsCancelled() bool
	Cancel()
}

type Listener interface {
	OnEvent(event Event) error
}

type listenerRegistry map[Type][]Listener

var (
	listeners       = make(listenerRegistry)
	cancellableType = TypeOf((*Cancellable)(nil)).Elem()
)

func Register(event Event, listener Listener) {
	// Debug logging, print event registration on function call
	if configuration.Config().DebugMode {
		log.Printf("Registering event listener: event=%s, listener=%s", TypeOf(event).Name(), TypeOf(listener).Name())
	}

	v := listeners[TypeOf(event)]
	if v == nil {
		v = make([]Listener, 0)
	}
	v = append(v, listener)
	listeners[TypeOf(event)] = v
}

func Trigger(event Event) (cancelled bool) {
	// Run on a separate goroutine to avoid hogging the spawning thread

	v := listeners[TypeOf(event)]
	if v == nil {
		return
	}

	cancelChannel := make(chan bool)

	for _, l := range v {
		if event.IsAsync() {
			// For long events, async it.
			// TODO: Thread pooling
			go func() {
				if err := l.OnEvent(event); err != nil {
					log.Panicln(err)
				}

				if TypeOf(event).Implements(cancellableType) {
					cancelChannel <- event.(Cancellable).IsCancelled()
				}
			}()
		} else {
			if err := l.OnEvent(event); err != nil {
				log.Panicln(err)
			}

			if TypeOf(event).Implements(cancellableType) && event.(Cancellable).IsCancelled() {
				logCancel(event)
				return true
			}
		}
	}
	close(cancelChannel)

	if TypeOf(event).Implements(cancellableType) {
		for cancel := range cancelChannel {
			if cancel {
				logCancel(event)
				return true
			}
		}
	}
	return false
}

func logCancel(e Event) {
	if configuration.Config().DebugMode {
		log.Printf("%s cancelled", TypeOf(e).Name())
	}
}
