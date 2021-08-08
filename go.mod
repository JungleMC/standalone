module github.com/JungleMC/standalone

go 1.16

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/JungleMC/java-edition v0.0.0-20210808094919-e7bc63e3faf6
	github.com/google/uuid v1.2.0
)

replace github.com/JungleMC/java-edition => ../java-edition

replace github.com/JungleMC/sdk => ../sdk
