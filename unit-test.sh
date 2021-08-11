#!/usr/bin/env bash

ALGORITHMS_DIR=./algorithms

echo "testing..."
go test -race -coverprofile ./output/coverage.out ${ALGORITHMS_DIR}/... -tags=ci
go tool cover -func=./output/coverage.out | tail -n1 | awk '{print "Total test coverage: " $3}'
