MAKEFILE_DIR = $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
WORKFLOW_DIR = ~/Dropbox/Alfred/Alfred.alfredpreferences/workflows
BUNDLE_ID = com.ruedap.emma-css
CLI_CMD = ./workflow/awc

default: build cli test

ci: deps build cli coveralls

release: clean build unlink link

godep:
cli:
	@echo "--> Running CLI commands"
	@$(CLI_CMD) pos sta

build:
	@echo "--> Compiling packages and dependencies"
	@mkdir -p ./workflow/
	godep go build -ldflags '-s -w' -o $(CLI_CMD)

coveralls:
	@echo "--> Testing packages and sending coverage report"
	@godep go test -v -covermode=count -coverprofile=coverage.out
	@goveralls -coverprofile=coverage.out -service=travis-ci

deps:
	@echo "--> Installing build dependencies"
	go get -u github.com/tools/godep
	godep go install
	go get github.com/axw/gocov/gocov
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls

test:
	@echo "--> Testing packages"
	@godep go test -v -cover

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

.PHONY: default ci release cli build coveralls deps test clean link unlink
