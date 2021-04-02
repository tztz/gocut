#!/bin/sh

# Convenience script to perform linting, testing, and building.
./scripts/lint.sh && \
./scripts/test.sh && \
./scripts/build.sh $1

# Last but not least:
go mod tidy
