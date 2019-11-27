.PHONY: test

ALGORITHMS_DIR = ./algorithms

all: fmt test
	@echo "All done."

test:
	@echo "testing..."
	@go test -race -coverprofile ./output/coverage.out $(ALGORITHMS_DIR)/...

fmt:
	@./fmt.sh

lint:
	@golint $(ALGORITHMS_DIR)

mod:
	@go mod download