//+build dev

package configuration

import (
	. "github.com/junglemc/JungleTree/pkg/util"
)

func root() RootConfiguration {
	return RootConfiguration{
		DebugMode:        true,
		Verbose:          true,
		MOTD:             "JungleTree Debug",
		MaxOnlinePlayers: 20,
		Network: NetConfig{
			IP:                          "",
			Port:                        25565,
			NetworkCompressionThreshold: 256,
		},
		Gamemode:    Survival.String(),
		Difficulty:  Normal.String(),
		JavaEdition: JavaEditionConfig{OnlineMode: false},
	}
}
