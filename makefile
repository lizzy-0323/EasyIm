# Makefile for go-im project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=go-im
BINARY_UNIX=$(BINARY_NAME)_unix

# Directories
SRC_DIR=./cmd
BUILD_DIR=./build
MSG_GATEWAY=/msg-gateway

# Targets
all: test build

.PHONY:build
build: 
	$(GOBUILD) -o $(BUILD_DIR)/gateway -v $(SRC_DIR)/$(MSG_GATEWAY)

.PHONY:test
test: 
	$(GOTEST) -v ./...

.PHONY:clean
clean: 
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

.PHONY: install
install:
	$(GOGET) -v ./...

.PHONY: all build clean test run deps