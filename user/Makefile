PROJECT_NAME := github.com/cyicz123/todolist
PKG := $(PROJECT_NAME)
USER_NAME := user
PKG_USER_LIST := $(shell go list $(PKG)/$(USER_NAME)/...)
CONFIG_DIR := $(HOME)/.config/todolist

.PHONY: all dep lint vet test test-coverage build clean

all: build

dep: ## Get the dependencies
	if [ ! -d $(CONFIG_DIR) ]; then \
        mkdir $(CONFIG_DIR); \
		cp -p ./config/config.yml $(CONFIG_DIR); \
    fi
	go mod tidy

test: ## Run unittests
	go test -short $(PKG_USER_LIST)

test-coverage: ## Run tests with coverage
	go test -short -coverprofile coverage.txt -covermode=atomic $(PKG_USER_LIST)

build: dep ## Build the binary file
	go build -i -o build/main $(PKG)

clean: ## Remove previous build
	rm -f ./build

help: ## Display this help screen
	grep -h -E '^[a-zA-Z_-]+:.*?## .*?' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", ?1, ?2}'
