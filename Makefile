GOLANGCI_LINT_VERSION := v1.62.2

.PHONY: lint build test ci

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./...

build:
	go build ./...

test:
	go test ./...

ci: lint build test
