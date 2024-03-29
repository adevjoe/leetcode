.PHONY: test

ALGORITHMS_DIR = ./algorithms
OUTPUT = ./output

all: fmt test
	@echo "All done."

test: check-dir
	@./unit-test.sh

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

dep:
	@python3 -m pip install -r Requirements.txt

fetch-all:
	@./leetcode.py -a

new:
	@./leetcode.py -n --id $(id)

finish:
	@./leetcode.py -f --id $(id)

ff: new finish