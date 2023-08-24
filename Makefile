.PHONY: default build run clean deps docs test
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
	@rm -rf ./docs
	@swag init

test:
	@go test ./ ...

run-with-docs:
	@swag init
	@air
