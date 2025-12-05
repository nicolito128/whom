# Binary output name
BIN ?= whom

# Package name
PKG := github.com/nicolito128/whom

# Architecture
ARCH ?= $(shell go env GOOS)-$(shell go env GOARCH)

# Program version
VERSION ?= main

# Output directory
OUTPUT_DIR ?= _oputput

# Go environment
platform = $(subst -, ,$(ARCH))
GOOS = $(word 1, $(platform))
GOARCH = $(word 2, $(platform))
GOPROXY ?= "https://proxy.golang.org,direct"

.PHONY: all
all:
	@$(MAKE) build

build: _output/bin/$(GOOS)/$(GOARCH)/$(BIN)

_output/bin/$(GOOS)/$(GOARCH)/$(BIN): build-dirs
	@echo "building: $@"
		GOOS=$(GOOS) \
		GOARCH=$(GOARCH) \
		VERSION=$(VERSION) \
		PKG=$(PKG) \
		BIN=$(BIN) \
		OUTPUT_DIR=$$(pwd)/_output/bin/$(GOOS)/$(GOARCH) \
		./scripts/build.sh

build-dirs:
	@mkdir -p _output/bin/$(GOOS)/$(GOARCH)

.PHONY: clean
clean:
	rm -rf _output

.PHONY: modules
modules:
	go mod tidy

.PHONY: run
run: build
	$$(pwd)/_output/bin/$(GOOS)/$(GOARCH)/$(BIN) \

