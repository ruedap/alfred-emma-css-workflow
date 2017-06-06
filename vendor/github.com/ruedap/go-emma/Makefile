NAME := emma

.DEFAULT_GOAL := build

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif
	glide install

.PHONY: generate
generate:
ifeq ($(shell command -v esc 2> /dev/null),)
	go get -v github.com/mjibson/esc
endif
	go generate

.PHONY: deps
deps: glide generate

.PHONY: build
build:
	go build -o bin/$(NAME) cmd/$(NAME)/main.go

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: test
test:
	go test -cover -v `glide novendor`
