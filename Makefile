PROJECT="cpds-analyzer"
BIN=./bin/$(PROJECT)

BUILD_LDFLAGS := "-s -w"

.PHONY: all
all: clean build

.PHONY: build
build: 
	@echo "Building $(PROJECT)"
	@go build -ldflags=$(BUILD_LDFLAGS) -o $(BIN) ./cmd/main.go

clean:
	@echo "clean"
	@rm -rf ./bin

default: build
