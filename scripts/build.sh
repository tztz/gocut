#!/bin/sh

BUILD_TAG="${1}"

if [ -z "${BUILD_TAG}" ]; then
    BUILD_TAG="dev"
fi

# build for local use with given build tag (default is "dev")
echo "Building with build tag \"${BUILD_TAG}\" ..."
go build -tags=${BUILD_TAG} -o out/gocut cmd/main.go
