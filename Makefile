all: lint

run:
	go run main.go

## golangci-lint run
lint:
	golangci-lint cache clean
	golangci-lint run --config .golangci.yml --timeout=5m
