package client

import (
	"fmt"
	"github.com/junglemc/JungleTree/pkg/rpc/service"
	"google.golang.org/grpc"
)

func Connect(address string, port int) (service.WorldProviderClient, *grpc.ClientConn, error) {
	worldConnection, err := grpc.Dial(fmt.Sprintf("%v:%v", address, port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, err
	}
	return service.NewWorldProviderClient(worldConnection), worldConnection, nil
}
