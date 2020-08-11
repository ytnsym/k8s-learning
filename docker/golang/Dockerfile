FROM golang:1.14.0-alpine3.11 AS builder

RUN mkdir -p /go/src

COPY . /go/src

WORKDIR /go/src

RUN GOOS=linux go build -o bin/main

FROM alpine:3.11

EXPOSE 8080

COPY --from=builder /go/src/bin/main .

ENTRYPOINT [ "./main" ]
