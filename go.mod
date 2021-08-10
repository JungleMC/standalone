module github.com/JungleMC/standalone

go 1.16

require (
	github.com/JungleMC/java-edition v0.0.0-20210809140412-861ede23c72b
	github.com/JungleMC/login-service v0.0.0-20210810041609-ba221521faec
)

replace (
	github.com/JungleMC/java-edition v0.0.0-20210809140412-861ede23c72b => ../java-edition
	github.com/JungleMC/login-service v0.0.0-20210810041609-ba221521faec => ../login-service
	github.com/JungleMC/sdk v0.0.0-20210809140359-e8dcfa68f6af => ../sdk
)
