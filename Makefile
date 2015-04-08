DEPS = $(shell go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
MAKEFILE_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
WORKFLOW_DIR = ~/Dropbox/Alfred/Alfred.alfredpreferences/workflows
BUNDLE_ID = com.ruedap.emma-css
CLI_CMD = ./workflow/awc

default: build cli test

cli:
	@echo "--> Running CLI commands"
	@$(CLI_CMD) pos sta
	@echo

build:
	@echo "--> Compiling packages and dependencies"
	@mkdir -p ./workflow/
	go build -ldflags '-s -w' -o $(CLI_CMD)

test:
	@echo "--> Testing packages"
	@go test -v -cover

clean:
	@echo "--> Cleaning workflow files"
	@- rm $(CLI_CMD)

link:
	@echo "--> Linking workflow files"
	@- ln -s $(MAKEFILE_DIR)/workflow $(WORKFLOW_DIR)/$(BUNDLE_ID)

unlink:
	@echo "--> Unlinking workflow files"
	@- rm $(WORKFLOW_DIR)/$(BUNDLE_ID)
	@- rm ./workflow/workflow

release: clean build unlink link

.PHONY: default cli build test clean link unlink release
