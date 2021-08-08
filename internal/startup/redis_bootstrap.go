package startup

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hdt3213/godis/config"
	"github.com/hdt3213/godis/redis/server"
	"github.com/hdt3213/godis/tcp"
	"net"
)

const internalRedisPort = 6380

func RedisBootstrap() (*redis.Client, net.Listener, *server.Handler) {
	config.Properties = &config.ServerProperties{
		Bind:           "127.0.0.1",
		Port:           internalRedisPort,
		AppendOnly:     false,
		AppendFilename: "",
		MaxClients:     100,
	}

	listener, err := net.Listen("unix", "/tmp/godis.sock")
	if err != nil {
		panic(err)
	}

	handler := server.MakeHandler()
	tcp.ListenAndServe(listener, handler, nil)

	if err != nil {
		panic(err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "localhost", internalRedisPort),
		Password: "",
		DB:       0,
	}), listener, handler
}
