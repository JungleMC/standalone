package startup

import (
	"github.com/go-redis/redis/v8"
	"github.com/hdt3213/godis/config"
	"github.com/hdt3213/godis/redis/server"
	"github.com/hdt3213/godis/tcp"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const internalRedisPort = 6380

func RedisBootstrap() *redis.Client {
	config.Properties = &config.ServerProperties{
		AppendOnly:     false,
		AppendFilename: "",
		MaxClients:     100,
	}

	_ = os.Remove("/tmp/godis.sock")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		<-quit
		_ = os.Remove("/tmp/godis.sock")
		close(quit)
		os.Exit(0)
	}()

	listener, err := net.Listen("unix", "/tmp/godis.sock")
	if err != nil {
		panic(err)
	}

	handler := server.MakeHandler()
	go tcp.ListenAndServe(listener, handler, nil)

	if err != nil {
		panic(err)
	}

	return redis.NewClient(&redis.Options{
		Network: "unix",
		Addr:     "/tmp/godis.sock",
		Password: "",
		DB:       0,
	})
}
