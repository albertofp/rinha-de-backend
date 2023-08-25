FROM golang:1.21-alpine3.18 as build

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o main.go ./bin/rinha

FROM alpine:3.18 as runtime

WORKDIR /app

RUN mkdir /pprof

COPY --from=build /app/rinha .

CMD ["./rinha"]

