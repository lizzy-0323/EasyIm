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
BUILD_DIR=./bin
MSG_GATEWAY=msg-gateway
API_GATEWAY=api-gateway

# Targets
all: test build cli-build grpc-build

.PHONY:build
build: grpc-build msg-gateway-build cli-build 

.PHONY: msg-gateway-build
msg-gateway-build:
	$(GOBUILD) -o $(BUILD_DIR)/msg-gateway -v $(SRC_DIR)/$(MSG_GATEWAY)

.PHONY: cli-build
cli-build:	
	$(GOBUILD) -o $(BUILD_DIR)/chat  ./client/main.go

.PHONY: grpc-build
grpc-build:	
	protoc --go_out=.. --go-grpc_out=.. ./pkg/protocol/proto/*.proto -I ./pkg/protocol/proto

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