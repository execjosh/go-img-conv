PROJECT = go-img-conv
PACKAGE = github.com/execjosh/${PROJECT}
TARGET = img-conv
EXTLDFLAGS = -extldflags -static $(null)

all: build

.PHONY: build
build: main.go imgconv/imgconv.go internal/decoder/decoder.go internal/encoder/encoder.go
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o 'bin/${TARGET}' -ldflags '${EXTLDFLAGS}' '${PACKAGE}'

.PHONY: test
test:
	go test

.PHONY: coverage
coverage:
	@rm prof.out || true
	go get -u github.com/haya14busa/goverage
	'${GOPATH}/bin/goverage' -v -coverprofile prof.out ./imgconv/ ./internal/decoder/ ./internal/encoder/
	go tool cover -html=prof.out
