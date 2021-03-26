#!/bin/sh

# run production Docker container locally with "prod" profile
docker run --rm --publish 3000:3000 gocut
