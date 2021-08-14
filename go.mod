module github.com/JungleMC/standalone

go 1.16

replace (
	github.com/JungleMC/java-edition => ../java-edition
	github.com/JungleMC/login-service => ../login-service
	github.com/JungleMC/protocol => ../protocol
	github.com/JungleMC/sdk => ../sdk
)

require (
	github.com/JungleMC/java-edition v0.0.0-20210809140412-861ede23c72b
	github.com/JungleMC/login-service v0.0.0-20210810041609-ba221521faec
)
