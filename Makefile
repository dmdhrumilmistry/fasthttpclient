build:
	@go build -o bin/fhc cmd/fhc/main.go

run:
	@go run cmd/fhc/main.go

test:
	@go test -v ./...