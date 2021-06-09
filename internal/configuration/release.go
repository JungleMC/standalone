//+build !dev

package configuration

import (
	. "github.com/junglemc/JungleTree/pkg/util"
)

func root() RootConfiguration {
	return RootConfiguration{
		DebugMode:        false,
		Verbose:          false,
		MOTD:             "A JungleTree Server",
		MaxOnlinePlayers: 20,
		Network: NetConfig{
			IP:                          "",
			Port:                        25565,
			NetworkCompressionThreshold: 256,
		},
		Gamemode:      Survival.String(),
		Difficulty:    Normal.String(),
		JavaEdition:   JavaEditionConfig{OnlineMode: true},
		BannedPlayers: make([]Ban, 0),
	}
}
