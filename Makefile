.PHONY: build

MAIN_DIRECTORY := ./cmd/notes
BIN_NAME := notes

# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
DOCKER_BUILD_PLATFORMS ?= linux/amd64

dist:
	mkdir dist

build: dist
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o "./dist/${GOOS}/${GOARCH}/${BIN_NAME}" ${MAIN_DIRECTORY}

build-linux-arm64: export GOOS := linux
build-linux-arm64: export GOARCH := arm64
build-linux-arm64:
	make build

build-linux-amd64: export GOOS := linux
build-linux-amd64: export GOARCH := amd64
build-linux-amd64:
	make build

## Build multi-arch Docker images
multi-arch-image-%: build-linux-amd64 build-linux-arm64
	docker buildx build $(DOCKER_BUILDX_ARGS) -t youkoulayley/$(BIN_NAME):$* --platform=$(DOCKER_BUILD_PLATFORMS) -f buildx.Dockerfile .
