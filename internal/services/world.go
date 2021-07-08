package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/junglemc/JungleTree/internal/storage"
	"github.com/junglemc/JungleTree/pkg/rpc/message"
	"github.com/junglemc/JungleTree/pkg/rpc/service"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

func World(address string, port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	service.RegisterWorldProviderServer(s, &server{})

	log.Printf("World service running on %v:%v", address, port)
	if err = s.Serve(listener); err != nil {
		return err
	}
	return nil
}

func defaultWorldName() string {
	name := os.Getenv("DEFAULT_WORLD_NAME")
	if name == "" {
		return "world"
	}
	return name
}

type server struct {
	service.UnimplementedWorldProviderServer
}

func (s *server) CreateWorld(_ context.Context, msg *message.World) (*message.World, error) {
	name := msg.GetName()
	if name == "" {
		name = defaultWorldName()
	}
	worldId := fmt.Sprintf("world:%v", name)

	worlds := make([]string, 0, 0)
	err := storage.Get("minecraft:worlds", &worlds, nil)
	if err != nil && err != errors.ErrNotFound {
		return &message.World{}, status.Error(codes.Internal, err.Error())
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return &message.World{}, status.Error(codes.Internal, err.Error())
	}
	idString := id.String()

	worlds = append(worlds, idString)
	err = storage.Put("minecraft:worlds", worlds, nil)
	if err != nil {
		return &message.World{}, status.Error(codes.Internal, err.Error())
	}

	result := msg
	result.Uuid = &idString

	err = storage.Put(fmt.Sprintf("world:%v", worldId), result, nil)
	if err != nil {
		return &message.World{}, status.Error(codes.Internal, err.Error())
	}
	return &message.World{}, status.Error(codes.OK, "")
}

func (s *server) GetWorld(_ context.Context, msg *message.WorldGetRequest) (*message.World, error) {
	name := msg.GetName()
	if name == "" {
		name = defaultWorldName()
	}
	id := fmt.Sprintf("world:%v", name)

	ok, err := storage.Has(id, nil)
	if err != nil {
		return &message.World{}, status.Error(codes.Internal, err.Error())
	}

	if !ok {
		return &message.World{}, status.Error(codes.NotFound, "not found")
	}

	result := &message.World{}
	err = storage.Get(id, result, nil)
	if err != nil {
		return &message.World{}, status.Error(codes.Internal, err.Error())
	}
	return result, nil
}

func (s *server) UpdateWorld(context.Context, *message.World) (*message.World, error) {
	return nil, nil
}

func (s *server) DeleteWorld(context.Context, *message.WorldDeleteRequest) (*message.WorldDeleteResponse, error) {
	return nil, nil
}
