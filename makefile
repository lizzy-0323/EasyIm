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

# Binarys
CONN_SERVER=conn-server
LOGIC_SERVER = logic-server
BUSINESS_SERVER = business-server

# Targets
all: test build cli-build grpc-build

.PHONY:build
build: grpc-build cli-build conn-server logic-server business-server 

.PHONY: conn-server
conn-server:
	$(GOBUILD) -o $(BUILD_DIR)/$(CONN_SERVER) -v $(SRC_DIR)/$(CONN_SERVER)

.PHONY: logic-server
logic-server:
	$(GOBUILD) -o $(BUILD_DIR)/$(LOGIC_SERVER) -v $(SRC_DIR)/$(LOGIC_SERVER)

.PHONY: business-server
business-server:
	$(GOBUILD)  -o $(BUILD_DIR)/$(BUSINESS_SERVER) -v $(SRC_DIR)/$(BUSINESS_SERVER)
	
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