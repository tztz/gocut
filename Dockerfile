# ------------------------
ARG SERVICE_NAME=gocut
ARG PACKAGE_NAME_PREFIX=srv.tztz.io/example
ARG MAIN_GO_FILE=cmd/main.go
ARG GIN_MODE=release
# ------------------------

# Do not modify these ARGs:
ARG PACKAGE_NAME=${PACKAGE_NAME_PREFIX}/${SERVICE_NAME}
ARG BUILD_TAG=prod

##################################
# STEP 1 - build executable binary
##################################

# Start from the latest Alpine Linux image with the latest version of Go installed
FROM golang:alpine@sha256:353e19718d4aa37cb38cf362e5aba23e22b8680bfc18255408ccd9b7b777c469 as builder

ARG SERVICE_NAME
ARG MAIN_GO_FILE
ARG PACKAGE_NAME
ARG BUILD_TAG

# Create unprivileged app user
ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Copy the local package files to the container's workspace
ADD . /go/src/${PACKAGE_NAME}

# Set current workdir
WORKDIR /go/src/${PACKAGE_NAME}

# Build the service (the app) and all its dependencies with the given build tag
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags=${BUILD_TAG} -ldflags="-w -s" -o /go/bin/${SERVICE_NAME} ${MAIN_GO_FILE}

##############################
# STEP 2 - build a small image
##############################

FROM scratch

LABEL maintainer="markimo-the-dev@tztz.io"

ARG SERVICE_NAME
ARG GIN_MODE
ARG PACKAGE_NAME

# Import the user and group files from the builder stage (step 1)
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy the static executable
COPY --from=builder /go/bin/${SERVICE_NAME} /go/bin/main
# Copy the configs folder
COPY --from=builder /go/src/${PACKAGE_NAME}/configs /go/src/${PACKAGE_NAME}/configs
# Copy the web folder
COPY --from=builder /go/src/${PACKAGE_NAME}/web /go/src/${PACKAGE_NAME}/web

# Use the unprivileged user
USER appuser:appuser

ENV GIN_MODE=${GIN_MODE}

# Run the service by default when the container starts
ENTRYPOINT ["/go/bin/main"]

# Port on which the service will be exposed
EXPOSE 3000
