module github.com/JungleMC/standalone

go 1.16

require (
	github.com/JungleMC/java-edition v0.0.0-20210808094919-e7bc63e3faf6
	github.com/go-redis/redis/v8 v8.11.2
	github.com/google/uuid v1.2.0
	github.com/hdt3213/godis v1.2.7
	github.com/junglemc/JungleTree v0.0.12
)

replace github.com/JungleMC/java-edition => ../java-edition

replace github.com/JungleMC/sdk => ../sdk
