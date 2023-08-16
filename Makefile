build clean deps:
	go build -o bin/app

run: build
	./bin/app

clean:
	rm -rf ./bin

deps:
	go mod tidy
