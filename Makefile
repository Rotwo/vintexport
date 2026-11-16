# Makefile created following https://makefiletutorial.com/ adapted to this go project. 

TARGET := ./src/main.go
BUILD_DIR := ./build
OUTPUT := $(BUILD_DIR)/vintexport

# go build related variables
GOOS   ?= linux # linux / windows / darwin
GOARCH ?= amd64 # amd64 / arm64

run:
	go run $(TARGET)

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT) $(TARGET)

.PHONY: clean
clean:
	rm -r $(BUILD_DIR)