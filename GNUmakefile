# GNUmakefile

# Configure shell path
SHELL := /bin/bash

# Name of the binary to be built
BINARY_NAME := mechanical-drill
PRODUCT_VERSION := $(shell cat ./API_VERSION)

# Source directory
SRC_DIR := .

# Build directory
BUILD_DIR := ./build
ARTIFACTS_DIR := ./artifacts

# Exclude specific directories and/or file patterns
EXCLUDE_DIR := ./tests
EXCLUDE_PATTERN := *.test.go

# Find command adjusted to exclude the specified directories and patterns
SOURCES := $(shell find $(SRC_DIR) -name '*.go' ! -path "$(EXCLUDE_DIR)/*" ! -name "$(EXCLUDE_PATTERN)")

# Docker-related variables
DOCKER_IMAGE := mechanical-drill
DOCKER_TAG := latest
IMAGE_DISTRIBUTOR := cloudputation



# Phony targets for make commands
.PHONY: all build clean docker-build docker-push

# Default target
all: build docker-build docker-push

# Build the binary
build: $(SOURCES)
	@echo "Downloading dependencies..."
	@GO111MODULE=on go mod tidy
	@GO111MODULE=on go mod download
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)

# Build the Docker image
docker-build: build
	@echo "Building the Docker image..."
	docker build --build-arg PRODUCT_VERSION=$(PRODUCT_VERSION) -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

# Push the Docker image to the registry
docker-push:
	@echo "Pushing the Docker image..."
	docker tag $(DOCKER_IMAGE):$(DOCKER_TAG) $(DOCKER_REGISTRY)/$(IMAGE_DISTRIBUTOR)/$(DOCKER_IMAGE):$(DOCKER_TAG)
	docker push $(DOCKER_REGISTRY)/$(IMAGE_DISTRIBUTOR)/$(DOCKER_IMAGE):$(DOCKER_TAG)

# Clean up
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@rm -rf $(ARTIFACTS_DIR)
