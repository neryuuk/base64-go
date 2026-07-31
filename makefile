# Variables
PROJECT=neryuuk/base64-go
BINARY_NAME=base64
VERSION=1.0.0
BUILD_DIR=bin

.PHONY: all build clean

build:
	@echo "Building $(PROJECT)..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags "-X main.Version=$(VERSION)" -o $(BUILD_DIR)/$(BINARY_NAME) $(PROJECT)

clean:
	@echo "Cleaning up $(PROJECT)..."
	@rm -rf $(BUILD_DIR)
