build: clean deps
	go build -o bin/app

run: build docker-start
	./bin/app

docker-start:
	docker run --name some-mongo -p 27017:27017 -d mongo

clean:
	rm -rf ./bin

deps:
	go mod tidy
