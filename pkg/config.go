package pkg

import (
	"github.com/junglemc/mc"
	"github.com/pelletier/go-toml"
	"io/fs"
	"io/ioutil"
	"os"
)

const configFile = "config.toml"

var config RootConfiguration

type RootConfiguration struct {
	DebugMode        bool
	Verbose          bool
	MOTD             string
	MaxOnlinePlayers int
	Network          NetConfig
	Gamemode         string
	Difficulty       string
	JavaEdition      JavaEditionConfig
}

type NetConfig struct {
	IP                          string
	Port                        uint16
	NetworkCompressionThreshold int
}

type JavaEditionConfig struct {
	OnlineMode bool
}

func init() {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		createDefaults()
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
}

func createDefaults() {
	serverConfig := NetConfig{
		IP:                          "",
		Port:                        25565,
		NetworkCompressionThreshold: 256,
	}

	jeConfig := JavaEditionConfig{OnlineMode: true}

	root := RootConfiguration{
		DebugMode:        true,
		Verbose:          false,
		MOTD:             "A JungleTree Server",
		MaxOnlinePlayers: 20,
		Network:          serverConfig,
		Gamemode:         mc.Survival.String(),
		Difficulty:       mc.Normal.String(),
		JavaEdition:      jeConfig,
	}

	data, err := toml.Marshal(&root)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(configFile, data, fs.FileMode(0664))
	if err != nil {
		panic(err)
	}
}

func Config() *RootConfiguration {
	return &config
}
