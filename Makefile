# Go parameters
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOMOD=$(GOCMD) mod

# Version information
VERSION=$(shell git describe --tags --always 2>/dev/null || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Build flags
LDFLAGS=-ldflags "-s -w -X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.BuildDate=$(BUILD_DATE)"
BUILD_FLAGS=-trimpath $(LDFLAGS)

# Binary paths
BIN_DIR=build
BINARY=$(BIN_DIR)/inwx

# Default suffix is empty, can be overridden for different platforms
SUFFIX=
ifeq ($(GOOS),windows)
	SUFFIX=.exe
endif

.PHONY: all build clean deps fmt vet lint run help version dist

all: build

version:
	@echo "Version: $(VERSION)"
	@echo "Commit: $(COMMIT)"
	@echo "Build Date: $(BUILD_DATE)"

build: deps
	mkdir -p $(BIN_DIR)
	CGO_ENABLED=0 $(GOCMD) build $(BUILD_FLAGS) -o $(BINARY)$(SUFFIX) ./cmd/inwx

clean:
	$(GOCLEAN)
	rm -rf $(BIN_DIR)
	rm -rf dist

deps:
	$(GOMOD) tidy

fmt:
	$(GOCMD) fmt ./...

vet:
	$(GOCMD) vet ./...

lint:
	golangci-lint run

run: build
	./$(BINARY)

install:
	$(GOCMD) install $(BUILD_FLAGS) ./cmd/inwx

help:
	@echo "Make targets:"
	@echo "  all        - Build the binary"
	@echo "  build      - Build the binary"
	@echo "  clean      - Remove binary and build artifacts"
	@echo "  deps       - Install dependencies"
	@echo "  fmt        - Run go fmt"
	@echo "  vet        - Run go vet"
	@echo "  lint       - Run golangci-lint"
	@echo "  run        - Build and run the application"
	@echo "  install    - Install binary to GOPATH"
	@echo "  version    - Show version information"
	@echo "  dist       - Build binaries for multiple platforms"

dist: deps
	@echo "Building for multiple platforms..."
	@mkdir -p dist
	
	# Linux amd64
	@echo "Building for Linux (amd64)..."
	@mkdir -p dist/linux_amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build $(BUILD_FLAGS) -o dist/linux_amd64/inwx ./cmd/inwx
	
	# Linux arm64
	@echo "Building for Linux (arm64)..."
	@mkdir -p dist/linux_arm64
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOCMD) build $(BUILD_FLAGS) -o dist/linux_arm64/inwx ./cmd/inwx
	
	# Windows amd64
	@echo "Building for Windows (amd64)..."
	@mkdir -p dist/windows_amd64
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOCMD) build $(BUILD_FLAGS) -o dist/windows_amd64/inwx.exe ./cmd/inwx
	
	# Windows arm64
	@echo "Building for Windows (arm64)..."
	@mkdir -p dist/windows_arm64
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 $(GOCMD) build $(BUILD_FLAGS) -o dist/windows_arm64/inwx.exe ./cmd/inwx
	
	# macOS amd64
	@echo "Building for macOS (amd64)..."
	@mkdir -p dist/darwin_amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOCMD) build $(BUILD_FLAGS) -o dist/darwin_amd64/inwx ./cmd/inwx
	
	# macOS arm64
	@echo "Building for macOS (arm64)..."
	@mkdir -p dist/darwin_arm64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOCMD) build $(BUILD_FLAGS) -o dist/darwin_arm64/inwx ./cmd/inwx
	
	@echo "All builds completed in the dist/ directory"
