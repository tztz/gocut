#!/bin/sh

# run all tests with "test" profile
echo "Building and running tests with build tag \"test\" and profile \"test\" ..."
RUN_PROFILE="test" go test -tags=test ./... -cover
