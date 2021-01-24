DENO = deno
DIST_FILE=Emma.alfredworkflow
EXEC_BIN=$(DENO)
EXEC_SRC=./src
BIN_FILE=./emma-css-workflow
PLIST_FILE=./info.plist
ICON_FILE=./icon.png
DATA_FILE=./emma-data.yml
ENTRY_FILE=./src/mod.ts

all: clean compile build

build:
	zip -r $(DIST_FILE) $(BIN_FILE) $(PLIST_FILE) $(ICON_FILE) $(DATA_FILE)

compile:
	$(DENO) compile --unstable --lite --allow-read --output ${BIN_FILE} ${ENTRY_FILE}

clean:
	rm -rf $(DIST_FILE)

open: clean build
	open $(DIST_FILE)

run:
	$(DENO) run --allow-read ${ENTRY_FILE} $(args)

test:
	$(DENO) test

fmt:
	$(DENO) fmt

lint:
	$(DENO) lint --unstable

check: fmt lint
