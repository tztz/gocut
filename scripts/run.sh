#!/bin/sh

PROFILE="${1}"

if [ -z "${PROFILE}" ]; then
    PROFILE="dev"
fi

# build with given build tag and run with given profile (default is "dev")
echo "Building and running with build tag \"${PROFILE}\" and profile \"${PROFILE}\" ..."
RUN_PROFILE=${PROFILE} go run -tags=${PROFILE} cmd/main.go
