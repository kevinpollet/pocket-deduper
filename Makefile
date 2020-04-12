.PHONY: default clean lint build

BIN_NAME = "pocket-deduper"
DIST_DIR = "dist"

default: clean lint build

build: clean $(DIST_DIR)
	go build -o $(DIST_DIR)/$(BIN_NAME) $(BIN_NAME).go

lint:
	golangci-lint run

clean:
	rm -rf $(DIST_DIR)

$(DIST_DIR):
	mkdir -p $(DIST_DIR)
