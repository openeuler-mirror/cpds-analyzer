PROJECT="cpds-analyzer"
OUT=./out/$(PROJECT)

BUILD_LDFLAGS := "-s -w"

.PHONY: all
all: clean build

.PHONY: build
build: 
	@echo "Building $(PROJECT)"
	@go build -ldflags=$(BUILD_LDFLAGS) -o $(OUT) ./cmd/$(PROJECT)/main.go

clean:
	@echo "clean"
	@rm -rf ./out

default: build
