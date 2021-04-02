#!/bin/sh

# Install linter in order to perform static code analysis.
# The binary will be $(go env GOPATH)/bin/golangci-lint.
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.38.0

# --------

# Last but not least:
go mod tidy
