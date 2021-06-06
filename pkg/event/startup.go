package event

import (
	"log"

	"github.com/junglemc/JungleTree/pkg"
)

const (
	thickLine = "===================================="
	thinLine  = "------------------------------------"
)

type ServerStartupEvent struct{}
type ServerStartupListener struct{}

func (e ServerStartupEvent) IsAsync() bool {
    return false
}

type ServerLoadedEvent struct{}
type ServerLoadedListener struct{}

func (e ServerLoadedEvent) IsAsync() bool {
    return false
}

func (l ServerStartupListener) OnEvent(event Event) {
	log.Println(thickLine)
	log.Println("Starting JungleTree Server " + pkg.Version)
	log.Println(thickLine)
}

func (l ServerLoadedListener) OnEvent(event Event) {
	log.Println(thinLine)
	log.Println("Done!")
}
