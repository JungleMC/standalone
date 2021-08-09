module github.com/JungleMC/standalone

go 1.16

require (
	github.com/JungleMC/java-edition v0.0.0-20210808094919-e7bc63e3faf6
	github.com/JungleMC/login-service v0.0.0-20210809133537-973d79d88972
	github.com/JungleMC/sdk v0.0.0-20210809130650-a3783d7db03f
	github.com/stretchr/testify v1.7.0 // indirect
)

replace github.com/JungleMC/java-edition => ../java-edition

replace github.com/JungleMC/login-service => ../login-service

replace github.com/JungleMC/sdk => ../sdk
