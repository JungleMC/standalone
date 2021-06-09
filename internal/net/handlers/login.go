package handlers

import (
	"log"

	"github.com/google/uuid"

	"github.com/junglemc/JungleTree/internal/configuration"
	. "github.com/junglemc/JungleTree/internal/net"
	"github.com/junglemc/JungleTree/internal/net/auth"
	"github.com/junglemc/JungleTree/internal/net/protocol"
	. "github.com/junglemc/JungleTree/internal/pkg/net/packets"
	"github.com/junglemc/JungleTree/pkg"
	"github.com/junglemc/JungleTree/pkg/ability"
	. "github.com/junglemc/JungleTree/pkg/codec"
	"github.com/junglemc/JungleTree/pkg/crafting"
	"github.com/junglemc/JungleTree/pkg/event"
	"github.com/junglemc/JungleTree/pkg/util"
	"github.com/junglemc/JungleTree/pkg/world"
	"github.com/junglemc/JungleTree/pkg/world/dimensions"
)

func init() {
	event.Register(event.PlayerJoinEvent{}, event.PlayerJoinListener{})
}

func loginStart(c *Client, p Packet) (err error) {
	pkt := p.(ServerboundLoginStartPacket)

	c.Profile.Name = pkt.Username

	if c.Server.OnlineMode {
		pkt := &ClientboundLoginEncryptionRequest{
			ServerId:    "",
			PublicKey:   c.Server.PublicKey(),
			VerifyToken: c.ExpectedVerifyToken,
		}
		return c.Send(pkt)
	} else {
		if c.Server.CompressionThreshold >= 0 {
			if c.Server.Debug {
				log.Println("Compression enabled for " + c.Profile.Name)
			}
			err = c.EnableCompression()
			if err != nil {
				return
			}
		}

		c.Profile.ID, _ = uuid.NewRandom()
		err = loginSuccess(c)
		if err != nil {
			return
		}
		return joinGame(c)
	}
}

func loginEncryptionResponse(c *Client, p Packet) (err error) {
	pkt := p.(ServerboundLoginEncryptionResponsePacket)

	sharedSecret, err := auth.DecryptLoginResponse(c.Server.PrivateKey(), c.Server.PublicKey(), c.ExpectedVerifyToken, pkt.VerifyToken, pkt.SharedSecret, &c.Profile)
	if err != nil {
		return
	}

	err = c.EnableEncryption(sharedSecret)
	if err != nil {
		return
	}

	if c.Server.CompressionThreshold >= 0 {
		if c.Server.Debug {
			log.Println("Compression enabled for " + c.Profile.Name)
		}
		err = c.EnableCompression()
		if err != nil {
			return
		}
	}

	err = loginSuccess(c)
	if err != nil {
		return
	}
	return joinGame(c)
}

// TODO: Cleanup
func joinGame(c *Client) (err error) {
	err = sendJoinGame(c)
	if err != nil {
		return
	}

	Connect(c)

	err = sendServerBrand(c)
	if err != nil {
		return
	}

	err = sendServerDifficulty(c)
	if err != nil {
		return
	}

	err = sendPlayerAbilities(c)
	if err != nil {
		return
	}

	err = sendDeclaredRecipes(c)
	if err != nil {
		return err
	}

	err = sendAddPlayer(c)
	if err != nil {
		return err
	}

	err = sendUpdatePing(c)
	if err != nil {
		return err
	}

	return sendPositionLook(c)
}

func sendJoinGame(c *Client) (err error) {
	// TODO: Pull data from the application configuration, world generator, etc
	dimension, ok := dimensions.ByName("minecraft:overworld")
	if !ok {
		panic("dimension not found")
	}

	join := &ClientboundJoinGamePacket{
		EntityId:            0,
		IsHardcore:          false,
		GameMode:            util.Survival,
		PreviousGameMode:    -1,
		WorldNames:          []string{"minecraft:world"},
		DimensionCodec:      world.DimensionBiomes(),
		Dimension:           *dimension,
		WorldName:           "minecraft:world",
		HashedSeed:          0,
		MaxPlayers:          int32(configuration.Config().MaxOnlinePlayers),
		ViewDistance:        32,
		ReducedDebugInfo:    false,
		EnableRespawnScreen: true,
		IsDebug:             configuration.Config().DebugMode,
		IsFlat:              false,
	}
	return c.Send(join)
}

func sendServerBrand(c *Client) (err error) {
	return c.Send(&ClientboundPluginMessagePacket{
		Channel: "minecraft:brand",
		Data:    WriteString(pkg.Brand),
	})
}

func sendServerDifficulty(c *Client) (err error) {
	pkt := &ClientboundServerDifficultyPacket{
		Difficulty:       util.DifficultyByName(configuration.Config().Difficulty),
		DifficultyLocked: false,
	}
	return c.Send(pkt)
}

func sendPlayerAbilities(c *Client) (err error) {
	onlinePlayer, ok := GetOnlinePlayer(c)
	if !ok {
		return
	}

	abilities := ability.PlayerAbilities(0)

	if onlinePlayer.Gamemode == util.Creative {
		abilities = ability.Set(abilities, ability.Invulnerable)
		abilities = ability.Set(abilities, ability.AllowFlying)
		abilities = ability.Set(abilities, ability.CreativeMode)
	}

	pkt := &ClientboundPlayerAbilitiesPacket{
		Flags:        byte(abilities),
		FlyingSpeed:  0.5,
		WalkingSpeed: 0.1,
	}
	return c.Send(pkt)
}

func sendDeclaredRecipes(c *Client) (err error) {
	return c.Send(&ClientboundDeclareRecipesPacket{Recipes: crafting.Recipes()})
}

func sendAddPlayer(c *Client) (err error) {
	return c.Send(&ClientboundPlayerInfoPacket{
		Action: PlayerInfoActionAddPlayer,
		Data: []PlayerInfo{
			{
				UUID:           c.Profile.ID,
				Name:           c.Profile.Name,
				Properties:     c.Profile.Properties,
				Gamemode:       c.Gamemode,
				Ping:           0,
				HasDisplayName: false,
			},
		},
	})
}

func sendUpdatePing(c *Client) (err error) {
	return c.Send(&ClientboundPlayerInfoPacket{
		Action: PlayerInfoActionUpdateLatency,
		Data: []PlayerInfo{
			{
				UUID: c.Profile.ID,
				Ping: 0,
			},
		},
	})
}

func sendPositionLook(c *Client) (err error) {
	return c.Send(&ClientboundPositionLookPacket{
		X:          0,
		Y:          64,
		Z:          0,
		Yaw:        0,
		Pitch:      0,
		Flags:      0,
		TeleportId: 1,
	})
}

func loginSuccess(c *Client) (err error) {
	response := &ClientboundLoginSuccess{
		Uuid:     c.Profile.ID,
		Username: c.Profile.Name,
	}
	err = c.Send(response)
	if err != nil {
		return
	}

	event.Trigger(event.PlayerJoinEvent{
		Username: c.Profile.Name,
	})

	c.Protocol = protocol.Play
	return
}
