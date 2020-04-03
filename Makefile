build:
	go get -u ./...
	go build

run:
	shop web

fmt:
	go fmt ./...

lint:
	golangci-lint run

