.PHONY: build install build-release clean

BIN_NAME := rabbithole
CMD_DIR := ./cmd/rabbithole
VERSION := 0.1.0
COMMIT_HASH := $(shell git rev-parse --short HEAD)
BUILD_FLAGS := -ldflags '-X main.version=$(VERSION) -X main.commitHash=$(COMMIT_HASH) -w -extldflags "-static"'

export CGO_ENABLED=0

build:
	go build $(BUILD_FLAGS) -o $(BIN_NAME) $(CMD_DIR)

install:
	go install $(BUILD_FLAGS) $(CMD_DIR)

build-release:
	@echo "==> Checking for uncommitted changes"
	@test -z "$(shell git status --porcelain)" || { echo -e "\033[31mERR: Uncommitted changes found. Please commit your changes and try again\033[00m"; exit 1 ; }
	
	@echo "==> Linux 64-bit"
	@GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BIN_NAME)-linux-amd64 $(CMD_DIR)

	@echo "==> Linux 32-bit"
	@GOOS=linux GOARCH=386 go build $(BUILD_FLAGS) -o $(BIN_NAME)-linux-386 $(CMD_DIR)

	@echo "==> Linux ARM 64-bit"
	@GOOS=linux GOARCH=arm64 go build $(BUILD_FLAGS) -o $(BIN_NAME)-linux-arm64 $(CMD_DIR)

	@echo "==> Linux ARM 32-bit"
	@GOOS=linux GOARCH=arm go build $(BUILD_FLAGS) -o $(BIN_NAME)-linux-arm $(CMD_DIR)

	@echo "==> Darwin 64-bit"
	@GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) -o $(BIN_NAME)-darwin-amd64 $(CMD_DIR)

	@echo "==> Darwin 32-bit"
	@GOOS=darwin GOARCH=386 go build $(BUILD_FLAGS) -o $(BIN_NAME)-darwin-386 $(CMD_DIR)

clean:
	@rm rabbithole*