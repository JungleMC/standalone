package event

import (
	"log"

	"github.com/junglemc/JungleTree/internal/configuration"
)

type PlayerJoinEvent struct {
	Username string
}

type PlayerJoinListener struct{}

func (e PlayerJoinEvent) IsAsync() bool {
	return true
}

func (l PlayerJoinListener) OnEvent(event Event) error {
	e := event.(PlayerJoinEvent)
	if configuration.Config().DebugMode {
		log.Printf("%s joined the game\n", e.Username)
	}
	return nil
}
