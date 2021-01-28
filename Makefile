build:
	go get -u ./...
	go build

run:
	shop api

fmt:
	go fmt ./...

lint:
	golangci-lint run

