build:
	@go build -o bin/fhc cmd/main.go

run:
	@go run cmd/main.go

test:
	@go test -v ./...