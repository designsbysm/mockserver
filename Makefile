.DEFAULT_GOAL := start
.PHONY:
.SILENT:

benchmark:
	go test ./... -bench=. -benchmem

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html="coverage.out"

docs:
	godoc2md github.com/designsbysm/mockserver | sed -e s#src/target/##g > DOCUMENTATION.md

lint:
	golangci-lint run

start:
	gow -c -e=go,yaml run .

test:
	go test ./... -cover -race
