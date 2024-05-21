build:
	@go build -o bin/first-go-api cmd/main.go

run: build
	@./bin/first-go-api

test:
	@go test -v ./...