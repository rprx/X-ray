# NOTE: This MAKEFILE can be used to build Xray-core locally and in Automatic workflows. It is \
    provided for convenience in automatic building and functions as a part of it.
# NOTE: If you need to modify this file, please be aware that:\
    - This file is not the main Makefile; it only accepts environment variables and builds the \
    binary.\
    - Automatic building expects the correct binaries to be built by this Makefile. If you \
    intend to propose a change to this Makefile, carefully review the file below and ensure \
    that the change will not accidentally break the automatic building:\
        .github/workflows/release.yml \
    Otherwise it is recommended to contact the project maintainers.

# Define the name of the output binary
NAME = xray

# Define the version using the latest git commit description
VERSION = $(shell git describe --always --dirty)

# Linker flags and build parameters
LDFLAGS = -X github.com/xtls/xray-core/core.build=$(VERSION) -s -w -buildid=
PARAMS = -trimpath -ldflags "$(LDFLAGS)" -v

# Main package to build
MAIN = ./main

# Define the prefix for installation
PREFIX ?= $(shell go env GOPATH)

# Phony targets to avoid conflicts with files named 'clean', 'build', 'test', or 'deps'
.PHONY: clean build test deps

# Install dependencies
deps:
    go mod download

# Build target to compile the binary
build: deps
    CGO_ENABLED=0 go build -o $(NAME) $(PARAMS) $(MAIN)
ifeq ($(GOOS),windows)
    mv $(NAME) $(NAME).exe
    echo 'CreateObject("Wscript.Shell").Run "$(NAME).exe",0' > $(NAME)_no_window.vbs
else ifeq ($(GOARCH:0:4),mips)
    GOMIPS=softfloat CGO_ENABLED=0 go build -o $(NAME)_softfloat $(PARAMS) $(MAIN)
endif

# Run tests
test:
    go test ./...

# Clean target to remove generated files
clean:
    go clean -v -i $(PWD)
    rm -f $(NAME) $(NAME).exe $(NAME)_no_window.vbs $(NAME)_softfloat

# Default target
default: build
