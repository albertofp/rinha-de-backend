.PHONY: all
APP_NAME=rinha-go

default: run-with-docs

build: clean deps
	@go build -o bin/$(APP_NAME) main.go

run:
	@air

clean:
	@rm -rf ./bin

deps:
	@go mod tidy

docs:
	@swag init

test:
	@go test ./ ...

run-with-docs:
	@swag init
	@air
