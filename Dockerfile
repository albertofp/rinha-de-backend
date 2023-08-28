FROM golang:1.21-alpine3.18 as build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o -v rinha .

FROM alpine:3.18 as runtime

WORKDIR /app

RUN mkdir /pprof

COPY --from=build /app/rinha .

EXPOSE 8080

CMD ["./rinha"]

