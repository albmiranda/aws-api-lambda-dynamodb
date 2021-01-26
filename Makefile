# Go Makefile

GOCMD=go
BINARY=main
BINARY_SRCDIR=cmd/main/main.go

PKG_LIST = cmd/main internal/db internal/handler internal/http internal/satellite pkg/dynamodb
SRCS_LIST=$(shell find $(SRCDIR) -name "*.go")

.DEFAULT_GOAL := all

all: clean build test

build:	## Build main application
	@echo "\nBuilding GO project"
	$(GOCMD) build $(BINARY_SRCDIR)

test: ## Run unit test and code coverage
	@echo "\nTest and coverage"
	$(GOCMD) test -coverpkg=./... -coverprofile=profile.cov ./...
	@echo "\n"
	$(GOCMD) tool cover -func profile.cov

clean:	## Clean generated artifacts
	@echo "\nCleaning"
	$(GOCMD) clean
	-rm -f $(BINARY) $(BINARY).zip profile.cov

package: ${BINARY}	## Generate binary application as .zip artifact to be used on AWS Lambda
	zip $(BINARY).zip $(BINARY)

lint:	## Run golint to check style mistakes
	@echo "\nRunning golint c"
	@golint -set_exit_status ${PKG_LIST}

fmt:	## List format suggestions to the code
	@echo "\nRunning gofmt"
	@gofmt -l ${PKG_LIST}

vet:	## Examines the code and reports suspicious constructs
	@echo "\nRunning go vet"
	@$(foreach var,$(SRCS_LIST), $(GOCMD) vet $(var);)

help:	## Show this help message
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-13s\033[0m %s\n", $$1, $$2}'

.PHONY: all build clean test package lint fmt vet help
