build:
	go get -u ./...
	go build

web:
	cd frontend && npm i && npm run build:aot
	cp -rv frontend/dist/frontend/* ./static


run:
	shop api

fmt:
	go fmt ./...

lint:
	golangci-lint run

