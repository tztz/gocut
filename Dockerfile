############################
# STEP 1 build executable binary
############################

# Start from the latest Alpine image with the latest version of Go installed
FROM golang:alpine as builder

# Copy the local package files to the container's workspace
ADD . /go/src/tztz.io/example/gocut

WORKDIR /go/src/tztz.io/example/gocut

# Build the app with all its dependencies
RUN go build

# Install the app inside the container
# (Fetch or manage dependencies here, either manually or with a tool like "godep")
RUN go install srv.tztz.io/example/gocut

############################
# STEP 2 build a small image
############################

FROM scratch

# Copy static executable
COPY --from=builder /go/bin/gocut /go/bin/gocut

# Run the app by default when the container starts
ENTRYPOINT /go/bin/gocut

# Document that the service listens on port 3000
EXPOSE 3000
