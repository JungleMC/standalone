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
type ServerStartupListener struct{}

func (e ServerStartupEvent) IsAsync() bool {
    return false
}

type ServerLoadedEvent struct{}
type ServerLoadedListener struct{}

func (e ServerLoadedEvent) IsAsync() bool {
    return false
}

func (l ServerStartupListener) OnEvent(event Event) error {
    log.Println(thickLine)
    log.Println("Starting JungleTree Server v" + pkg.Version)
    log.Println(thickLine)
    return nil
}

func (l ServerLoadedListener) OnEvent(event Event) error {
    log.Println(thinLine)
    log.Println("Done!")
    return nil
}
