package event

import (
	"log"

	"github.com/junglemc/JungleTree/pkg"
)

const (
	thickLine = "===================================="
	thinLine  = "------------------------------------"
)

type (
	ServerStartupEvent struct{}
	ServerLoadedEvent  struct{}
)

type (
	ServerStartupListener struct{}
	ServerLoadedListener  struct{}
)

func (l ServerStartupListener) OnEvent(event Event) {
	log.Println(thickLine)
	log.Println("Starting JungleTree Server " + pkg.Version)
	log.Println(thickLine)
}

func (l ServerLoadedListener) OnEvent(event Event) {
	log.Println(thinLine)
	log.Println("Done!")
}
