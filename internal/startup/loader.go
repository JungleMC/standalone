package startup

import (
	"context"
	"github.com/junglemc/JungleTree/pkg/rpc/client"
	"github.com/junglemc/JungleTree/pkg/rpc/message"
	"github.com/junglemc/JungleTree/pkg/rpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/junglemc/JungleTree/internal/storage"
	"github.com/junglemc/JungleTree/pkg/block"
	"github.com/junglemc/JungleTree/pkg/crafting"
	"github.com/junglemc/JungleTree/pkg/entity"
	"github.com/junglemc/JungleTree/pkg/event"
	"github.com/junglemc/JungleTree/pkg/item"
)

const (
	TicksPerSecond = 20
)

var (
	WorldRPC        service.WorldProviderClient
	WorldConnection *grpc.ClientConn
)

func Init() {
	loadStorage()

	var err error
	WorldRPC, WorldConnection, err = client.Connect("127.0.0.1", 50051)
	if err != nil {
		log.Panicln(err)
	}

	err = createDefaultWorld()
	if err != nil {
		log.Panicln(err)
	}

	rand.Seed(time.Now().Unix())

	event.Trigger(event.ServerStartupEvent{})
	loadBlocks()
	loadItems()
	loadRecipes()
	loadEntities()
}

func createDefaultWorld() error {
	name := os.Getenv("DEFAULT_WORLD_NAME")
	if name == "" {
		name = "world"
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := WorldRPC.GetWorld(ctx, &message.WorldGetRequest{})
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.NotFound {
			// Should create

			uuid := ""

			world := &message.World{
				Uuid:                &uuid,
				Name:                name,
				Seed:                rand.Uint64(),
				Height:              256,
				Dimension:           "minecraft:overworld",
				InitialGamemode:     message.Gamemode_SURVIVAL,
				EnableRespawnScreen: true,
				IsFlat:              false,
				IsHardcore:          false,
				ReducedDebugInfo:    false,
			}

			result, err := WorldRPC.CreateWorld(ctx, world)
			if err != nil {
				return err
			}
			log.Println(result)
		} else {
			return err
		}
	}
	return nil
}

func loadStorage() {
	log.Println("\t* Loading LevelDB")
	err := storage.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func loadBlocks() {
	log.Println("\t* Loading blocks")

	err := block.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func loadEntities() {
	log.Println("\t* Loading entities")
	err := entity.Load()
	if err != nil {
		log.Panicln(err)
	}

	entityThread := entity.Runner{TPS: TicksPerSecond}
	entityThread.Run()
}

func loadItems() {
	log.Println("\t* Loading items")
	err := item.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func loadRecipes() {
	log.Println("\t* Loading recipes")
	err := crafting.Load()
	if err != nil {
		log.Panicln(err)
	}
}
