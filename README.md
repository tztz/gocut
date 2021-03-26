# gocut

![GoBuild](https://github.com/tztz/gocut/workflows/GoBuild/badge.svg)
![CodeQL](https://github.com/tztz/gocut/workflows/CodeQL/badge.svg)

A simple swipe from development to rollout with Golang.

## Install required build tooling

Some build tasks, such as source code linting, require additional tools to be installed first:

    scripts/install_tools.sh

Note: Keep your tooling up to date, i.e. regularly update the versions declared in the script and re-run it.

## Build, test, run

### Build and run for local use

This creates a statically linked executable binary file named "gocut":

    scripts/build.sh

Note: This builds the service with the "dev" build tag.

In order to build with e.g. "prod" build tag, execute:

    scripts/build.sh prod

Then run the _gocut_ service (the binary is placed in the "./out" folder) via:

    out/gocut

Note: This runs the service with the "prod" profile (independent of the build tag used).

In order to run with e.g. "dev" profile, execute:

    RUN_PROFILE=dev out/gocut

#### Run locally without creating a binary file

Alternatively, this builds **and** runs the _gocut_ service (without creating a binary file):

    scripts/run.sh

Note: This builds the service with the "dev" build tag and runs it with "dev" profile.

In order to build and run with e.g. "prod" build tag resp. profile, execute:

    scripts/run.sh prod

### Build for production use

This creates a production-ready Docker image:

    scripts/prod-build.sh

Note: The service is build with the "prod" build tag.

### Run production Docker container locally

    scripts/docker-run.sh

Note: The service is run with the "prod" profile.

### Run all tests

    scripts/test.sh

Note: This runs the tests with "test" profile.

### Lint the source code (aka run static code analysis)

    scripts/lint.sh

Note: For the linter to work the required tools must be installed beforehand (see "Install required build tooling").

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
