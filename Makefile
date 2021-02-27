all: build

build:
	mkdir -p dist
	go build -x -o dist/gdnsupdater
	GOOS=windows go build -x -o dist/gdnsupdater.exe

test:
	go test -v

.PHONY: all build test
