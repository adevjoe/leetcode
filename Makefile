.PHONY: test

ALGORITHMS_DIR = ./algorithms

all: fmt test
	@echo "All done."

test:
	@echo "testing..."
	@go test -race -cover $(ALGORITHMS_DIR)/...

fmt:
	@./fmt.sh

lint:
	@golint $(ALGORITHMS_DIR)