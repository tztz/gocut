# gocut

![GoBuild](https://github.com/tztz/gocut/workflows/GoBuild/badge.svg)
![CodeQL](https://github.com/tztz/gocut/workflows/CodeQL/badge.svg)

A simple swipe from development to rollout with Golang.

## Build for local use

    scripts/build.sh

## Run locally

    scripts/run.sh

Then point your browser to `localhost:3000`

Try the following endpoints:

- `localhost:3000/api/ping`
- `localhost:3000/admin/metrics`
- `localhost:3000/admin/healthcheck`

## Build for production use

This creates a production-ready Docker image:

    scripts/prod-build.sh

## Run production Docker container locally

    scripts/docker-run.sh

Then point your browser to `localhost:3000` and try the above endpoints.
