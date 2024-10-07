build:
	@go build -o bin/fhc examples/fhc/main.go

run:
	@go run examples/fhc/main.go

bump:
	@go get -u ./...
	@go mod tidy

test:
	@go test -v ./...

all: build test run 