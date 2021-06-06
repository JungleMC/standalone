package event

import (
    "github.com/junglemc/JungleTree/pkg"
    "log"
)

const (
    thickLine = "===================================="
    thinLine  = "------------------------------------"
)

type ServerStartupEvent struct{}
type ServerLoadedEvent struct{}

type ServerStartupListener struct{}
type ServerLoadedListener struct{}

func (l ServerStartupListener) OnEvent(event Event) {
    log.Println(thickLine)
    log.Println("Starting JungleTree Server v" + pkg.Version)
    log.Println(thickLine)
}

func (l ServerLoadedListener) OnEvent(event Event) {
    log.Println(thinLine)
    log.Println("Done!")
}
