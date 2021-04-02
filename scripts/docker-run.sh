#!/bin/sh

GIN_MODE=debug # default: release

# Run production Docker container locally with "prod" profile.
docker run -e GIN_MODE=${GIN_MODE} --rm --publish 3000:3000 gocut
