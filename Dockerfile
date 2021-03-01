# ------------------------
ARG SRV_NAME=gocut
ARG PACKAGE_NAME_PREFIX=srv.tztz.io/example
# ------------------------

ARG PACKAGE_NAME=${PACKAGE_NAME_PREFIX}/${SRV_NAME}

##################################
# STEP 1 - build executable binary
##################################

# Start from the latest Alpine Linux image with the latest version of Go installed
FROM golang:alpine@sha256:353e19718d4aa37cb38cf362e5aba23e22b8680bfc18255408ccd9b7b777c469 as builder

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

# Build the service (the app) with all its dependencies
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/${SRV_NAME}

##############################
# STEP 2 - build a small image
##############################

FROM scratch

# Import the user and group files from the builder stage (step 1)
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy the static executable
COPY --from=builder /go/bin/${SRV_NAME} /go/bin/${SRV_NAME}

# Use the unprivileged user
USER appuser:appuser

# Run the service by default when the container starts
ENTRYPOINT ["/go/bin/gocut"]

# Port on which the service will be exposed
EXPOSE 3000
