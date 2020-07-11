DENO = ./deno
DIST_FILE=emma.alfredworkflow
EXEC_BIN=$(DENO)
EXEC_SRC=./src
PLIST_FILE=./info.plist
ICON_FILE=./icon.png
DATA_FILE=./emma-data.yml

all: clean $(DIST_FILE)

$(DIST_FILE):
	zip -r $(DIST_FILE) $(EXEC_BIN) $(EXEC_SRC) $(PLIST_FILE) $(ICON_FILE) $(DATA_FILE) 

clean:
	rm -rf $(DIST_FILE)

run:
	$(DENO) run --allow-read ./src/mod.ts $(args)

test:
	$(DENO) test

fmt:
	$(DENO) fmt

lint:
	$(DENO) lint --unstable
