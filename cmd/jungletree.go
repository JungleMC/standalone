package main

import (
	java "github.com/JungleMC/java-edition/pkg/startup"
	"github.com/JungleMC/standalone/internal/startup"
)

func main() {
	rdb := startup.RedisBootstrap()

	defer func() {
		rdb.Close()
	}()

	java.Start(rdb)
}
