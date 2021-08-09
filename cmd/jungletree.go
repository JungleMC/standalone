package main

import (
	java "github.com/JungleMC/java-edition/pkg/service"
	login "github.com/JungleMC/login-service/pkg/service"
)

func main() {
	go login.Start()
	java.Start()
}
