# parameteres
BINARY_NAME=fogoquaser
SUBDIRS := $(wildcard */.)



all: test build

build: $(SUBDIRS)


.phony: build all