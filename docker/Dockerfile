FROM node:latest as frontend
RUN node -v
WORKDIR /build
ADD frontend ./
RUN npm doctor && \
    npm ci && \
    npm run build

FROM golang AS builder
LABEL maintainer="vdzundza@protonmail.com"

WORKDIR /app
ADD . ./
RUN go mod download
RUN go mod verify
RUN scripts/build.bash

FROM alpine:latest as certs
RUN apk --update add ca-certificates
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/dist /
COPY --from=frontend /build/dist/frontend /static
ENTRYPOINT ["/shop-linux-amd64"]
