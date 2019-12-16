.PHONY: test

ALGORITHMS_DIR = ./algorithms
OUTPUT = ./output

all: fmt test
	@echo "All done."

test: check-dir
	@echo "testing..."
	@go test -race -coverprofile ./output/coverage.out $(ALGORITHMS_DIR)/...

fmt:
	@./fmt.sh

lint:
	@golint $(ALGORITHMS_DIR)

mod:
	@go mod download

check-dir:
	@if [ ! -d $(OUTPUT) ]; then mkdir $(OUTPUT); fi

clean:
	@rm -rf $(OUTPUT)