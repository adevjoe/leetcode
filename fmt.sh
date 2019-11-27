#!/usr/bin/env bash

for f in $(gofmt -l $(find . -type f -name '*.go')) ; do
    go fmt ${f}
done

echo "All files be formatted."