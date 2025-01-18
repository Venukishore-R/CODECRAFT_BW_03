.PHONY: build

build:
	go mod vendor
	go build -o bin/main cmd/main.go

run:
	make build
	./bin/main
