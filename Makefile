VERS=0.2.0

all: build

build:
	mkdir -p dist
	go build -x -o dist/gdnsupdater
	GOOS=windows go build -x -o dist/gdnsupdater.exe

pack:
	(cd dist/; tar -czf gdnsupdater_${VERS}_Linux-64bit.tar.gz gdnsupdater)
	(cd dist/; zip -r gdnsupdater_${VERS}_Windows-64bit.zip gdnsupdater.exe)

release: build pack

test:
	go test -v

.PHONY: all build pack release test
