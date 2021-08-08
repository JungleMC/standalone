package main

import (
	java "github.com/JungleMC/java-edition/pkg/startup"
	"github.com/JungleMC/standalone/internal/startup"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = os.Remove("/tmp/godis.sock")

	rdb, listener, handler := startup.RedisBootstrap()

	defer func() {
		rdb.Close()
		handler.Close()
		listener.Close()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL)
	go func() {
		<-quit
		_ = os.Remove("/tmp/godis.sock")
		close(quit)
		os.Exit(0)
	}()

	java.Start(rdb)
}
