package event

import (
	"log"

	"github.com/junglemc/JungleTree/internal/configuration"
)

type PlayerLoginEvent struct {
	Username string
}

type PlayerLoginListener struct{}

func (e PlayerLoginEvent) IsAsync() bool {
	return true
}

func (l PlayerLoginListener) OnEvent(event Event) error {
	e := event.(PlayerLoginEvent)
	if configuration.Config().DebugMode {
		log.Printf("Player connecting: %s\n", e.Username)
	}
	return nil
}
