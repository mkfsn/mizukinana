PKG = github.com/mkfsn/mizukinana
PKG_LIST = $(shell go list $(PKG)/...)

.PHONY: all generate lint vet test test-coverage clean clean-coverage help

all: test build

generate:
	@go generate -v ./...

lint: ## Lint go files
	@golint -set_exit_status $(PKG_LIST)

vet: ## Run go vet
	@go vet $(PKG_LIST)

test: generate ## Run unittests
	@go test -short $(PKG_LIST)

test-coverage: generate ## Run unittests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic $(PKG_LIST)
	@cat cover.out >> coverage.txt

build: generate ## Build the binary
	@go build -i -o build/mizukinana $(PKG)/cmd/mizukinana

clean: ## Remove files from build
	@rm -rf build

clean-coverage: ## Remove files from test-coverage
	@rm -f cover.out coverage.txt

help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
