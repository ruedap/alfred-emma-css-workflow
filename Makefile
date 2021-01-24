DENO = deno
DIST_FILE=Emma.alfredworkflow
EXEC_BIN=$(DENO)
EXEC_SRC=./src
PLIST_FILE=./info.plist
ICON_FILE=./icon.png
DATA_FILE=./emma-data.yml

all: clean build

build:
	zip -r $(DIST_FILE) $(EXEC_BIN) $(EXEC_SRC) $(PLIST_FILE) $(ICON_FILE) $(DATA_FILE)

compile:
	$(DENO) compile --unstable --lite --allow-read --output ./emma-css-workflow ./src/mod.ts

clean:
	rm -rf $(DIST_FILE)

open: clean build
	open $(DIST_FILE)

run:
	$(DENO) run --allow-read ./src/mod.ts $(args)

test:
	$(DENO) test

fmt:
	$(DENO) fmt

lint:
	$(DENO) lint --unstable

check: fmt lint
