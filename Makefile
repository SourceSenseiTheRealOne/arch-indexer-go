SHELL := bash
BUILD_GOOS ?= linux
BUILD_GOARCH ?= amd64
BINARY := bin/indexerd-$(BUILD_GOOS)-$(BUILD_GOARCH)

.PHONY: fmt fmt-check test test-race vet build verify clean

fmt:
	@shopt -s globstar nullglob; files=(**/*.go); (($${#files[@]} == 0)) || gofmt -w "$${files[@]}"

fmt-check:
	@shopt -s globstar nullglob; files=(**/*.go); output=$$(gofmt -l "$${files[@]}"); test -z "$$output" || \
		(printf '%s\n%s\n' 'Go files require gofmt:' "$$output"; exit 1)

test:
	go test -count=1 ./...

test-race:
	go test -count=1 -race ./...

vet:
	go vet ./...

build:
	mkdir -p bin
	GOOS=$(BUILD_GOOS) GOARCH=$(BUILD_GOARCH) CGO_ENABLED=0 go build -trimpath -o $(BINARY) ./cmd/indexerd

verify:
	bash scripts/verify.sh

clean:
	rm -rf bin coverage
