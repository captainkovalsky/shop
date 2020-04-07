FROM golang:latest AS builder
LABEL maintainer="vdzundza@protonmail.com"
# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
ADD . ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN scripts/build.bash

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/dist .
COPY --from=builder /app/ ./static
ENTRYPOINT "./shop-linux-amd64"
