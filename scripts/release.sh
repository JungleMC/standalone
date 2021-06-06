#!/usr/bin/env bash

package=cmd/JungleTree.go
package_name=JungleTree
version=0.0.10
build_dir=./dist
# For cross-compilation
CGO=0

platforms=("linux/amd64" "linux/arm" "linux/arm64" "android/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH'-'$version
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env CGO_ENABLED=$CGO GOOS=$GOOS GOARCH=$GOARCH go build -o ${build_dir}/${output_name} $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
