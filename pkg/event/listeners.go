package event

import (
	"log"

	"github.com/junglemc/JungleTree/internal/configuration"
	"github.com/junglemc/JungleTree/pkg"
)

type (
	ServerStartupListener struct{}
	ServerLoadedListener  struct{}

	PlayerLoginListener struct{}
	PlayerJoinListener  struct{}
)

func (l ServerStartupListener) OnEvent(event Event) error {
	log.Println(thickLine)
	log.Printf("Starting JungleTree Server (Version: %s)\n", pkg.Version)
	log.Println(thickLine)
	return nil
}

func (l ServerLoadedListener) OnEvent(Event) error {
	log.Println(thinLine)
	log.Println("Done!")
	return nil
}

func (l PlayerLoginListener) OnEvent(Event) error {
	log.Println(thinLine)
	log.Println("Done!")
	return nil
}

func (l PlayerJoinListener) OnEvent(event Event) error {
	e := event.(PlayerJoinEvent)
	if configuration.Config().DebugMode {
		log.Printf("%s joined the game\n", e.Username)
	}
	return nil
}
