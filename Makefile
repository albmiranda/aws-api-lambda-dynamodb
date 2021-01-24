# Go Makefile

GOCMD=go
BINARY=main
BINARY_SRCDIR=cmd/main/main.go

PKG_LIST = cmd/main internal/db internal/handler internal/http internal/satellite pkg/dynamodb
SRCS_LIST=$(shell find $(SRCDIR) -name "*.go")

.DEFAULT_GOAL := all

all: clean validate build test

build:	## Build main application and compress its binary file into a zip file
	@echo "\nBuilding GO project"
	$(GOCMD) build $(BINARY_SRCDIR)
	zip $(BINARY).zip $(BINARY)

test: ## Run unit test and code coverage
	@echo "\nTest and coverage"
	$(GOCMD) test -coverpkg=./... -coverprofile=profile.cov ./...
	$(GOCMD) tool cover -func profile.cov

clean:	## Clean generated artifacts
	@echo "\nCleaning"
	$(GOCMD) clean
	-rm -f $(BINARY) $(BINARY).zip profile.cov

validate:	## Run golint, gofmt and go vet at every go package/files
	@echo "\nRunning golint c"
	@golint -set_exit_status ${PKG_LIST}

	@echo "\nRunning gofmt"
	@gofmt -l ${PKG_LIST}

	@echo "\nRunning go vet"
	@$(foreach var,$(SRCS_LIST), $(GOCMD) vet $(var);)

help:	## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-13s\033[0m %s\n", $$1, $$2}'

.PHONY: all build clean test validate help
