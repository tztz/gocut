# gocut

![GoBuild](https://github.com/tztz/gocut/workflows/GoBuild/badge.svg)
![CodeQL](https://github.com/tztz/gocut/workflows/CodeQL/badge.svg)

A simple swipe from development to rollout with Golang.

## Build, test, run

### Build and run for local use

This creates a statically linked executable binary file named "gocut":

    scripts/build.sh

Then run the _gocut_ service (the binary is placed in the "./out" folder) via:

    out/gocut

#### Run locally without creating a binary

Alternatively, this builds and runs the _gocut_ service (without creating a binary file):

    scripts/run.sh

### Build for production use

This creates a production-ready Docker image:

    scripts/prod-build.sh

### Run production Docker container locally

    scripts/docker-run.sh

### Run all tests

    scripts/test.sh

## Try the service

Point your browser to `localhost:3000`

Try the following endpoints:

- `localhost:3000/api/ping`
- `localhost:3000/admin/metrics`
- `localhost:3000/admin/healthcheck`

## Tech Stack

- Logrus
- Viper
- Gin Web Framework
- Prometheus

todo ...
