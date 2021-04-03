#!/bin/sh

#GIN_MODE=debug # default: release

# Run production Docker container locally with "prod" profile.
docker run -e GIN_MODE=${GIN_MODE} --rm --publish 8080:8080 gocut
