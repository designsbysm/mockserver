.DEFAULT_GOAL := start
.PHONY:
.SILENT:

benchmark:
	go test ./... -bench=. -benchmem

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html="coverage.out"

docs:
# 	gomarkdoc --output DOCUMENTATION.md ./...
	gomarkdoc --output DOCUMENTATION.md ./api/...

lint:
	golangci-lint run

start:
	gow -c -e=go,yaml run .

test:
	go test ./... -cover -race
