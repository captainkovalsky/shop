build:
	go get -u ./...
	go build

gen:
	/usr/local/proto/protoc --proto_path=proto --go_out=build/gen --go_opt=paths=source_relative cv.proto

web:
	cd frontend && npm i && npm run build:aot
	cp -rv frontend/dist/frontend/* ./static


run:
	shop api

fmt:
	go fmt ./...

lint:
	golangci-lint run

