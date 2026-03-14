GO := go
PKG := ./...

.PHONY: test lint coverage all

test:
	$(GO) test $(PKG)

coverage:
	$(GO) test -coverprofile=coverage.out $(PKG)
	$(GO) tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run

all: lint test coverage
