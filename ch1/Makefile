# What target is run if no target is specified
.DEFAULT_GOAL := build

.PHONY:fmt vet build # keeps make from doing something with a file named fmt in the project

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean
