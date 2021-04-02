#!/bin/sh

BUILD_TAG="${1}"

if [ -z "${BUILD_TAG}" ]; then
    BUILD_TAG="dev"
fi

# Build for local use with given build tag (default is "dev").
# The binary is statically linked.
echo "Building with build tag \"${BUILD_TAG}\" ..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags=${BUILD_TAG} -ldflags="-w -s" -o out/gocut cmd/main.go

# Alternatively, build fast but not statically linked.
#go build -tags=${BUILD_TAG} -o out/gocut cmd/main.go
