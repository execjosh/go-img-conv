PROJECT = go-img-conv
PACKAGE = github.com/execjosh/${PROJECT}
TARGET = img-conv
EXTLDFLAGS = -extldflags -static $(null)

all: build

.PHONY: build
build: main.go
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o 'bin/${TARGET}' -ldflags '${EXTLDFLAGS}' '${PACKAGE}'
