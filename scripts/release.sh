#!/usr/bin/env sh

BUILD_DIR=./dist

mkdir -p ${BUILD_DIR}

GOOS=linux GOARCH=amd64; go build -o ${BUILD_DIR}/JungleTree_${GOOS}_${GOARCH} cmd/JungleTree.go
GOOS=linux GOARCH=arm; go build -o ${BUILD_DIR}/JungleTree_${GOOS}_${GOARCH} cmd/JungleTree.go
GOOS=linux GOARCH=arm64; go build -o ${BUILD_DIR}/JungleTree_${GOOS}_${GOARCH} cmd/JungleTree.go
GOOS=android GOARCH=arm64; go build -o ${BUILD_DIR}/JungleTree_${GOOS}_${GOARCH} cmd/JungleTree.go
GOOS=darwin GOARCH=amd64; go build -o ${BUILD_DIR}/JungleTree_${GOOS}_${GOARCH} cmd/JungleTree.go
GOOS=darwin GOARCH=arm64; go build -o ${BUILD_DIR}/JungleTree_${GOOS}_${GOARCH} cmd/JungleTree.go
GOOS=windows GOARCH=amd64; go build -o ${BUILD_DIR}/JungleTree_${GOOS}_${GOARCH} cmd/JungleTree.go