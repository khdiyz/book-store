build:
	@go build -o ./bin/app ./cmd/main.go

run: build
	@./bin/app

swag:
	swag init -g cmd/main.go