package event

import "github.com/google/uuid"

const (
	thickLine = "===================================="
	thinLine  = "------------------------------------"
)

type (
	ServerStartupEvent struct{}
	ServerLoadedEvent  struct{}

	PlayerLoginEvent struct {
		ID               uuid.UUID
		Username         string
		cancel           bool
	}

	PlayerJoinEvent struct {
		ID               uuid.UUID
		Username         string
		cancel           bool
	}
)

func (e ServerStartupEvent) IsAsync() bool {
	return false
}

func (e ServerLoadedEvent) IsAsync() bool {
	return false
}

func (e PlayerLoginEvent) IsAsync() bool {
	return false
}

func (e PlayerLoginEvent) IsCancelled() bool {
	return e.cancel
}

func (e PlayerLoginEvent) Cancel() {
	e.cancel = true
}

func (e PlayerJoinEvent) IsAsync() bool {
	return true
}

func (e PlayerJoinEvent) Cancel() {
	e.cancel = true
}

func (e PlayerJoinEvent) IsCancelled() bool {
	return e.cancel
}
