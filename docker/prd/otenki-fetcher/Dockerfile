FROM golang:1.14.4 as builder

WORKDIR /go/src/github.com/nachiguro1003/otenki/src/fetcher

COPY ./src/fetcher/go.mod ./
COPY ./src/fetcher/go.sum ./

RUN go mod download

COPY ./src/fetcher/  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -o /go/bin/main -ldflags '-s -w'

FROM alpine:3.7

COPY --from=builder /go/src/github.com/nachiguro1003/otenki/src/fetcher/env-prd.yaml ./env.yaml
COPY --from=builder /go/bin/main ./main

RUN apk add --no-cache curl

EXPOSE 8000

ENTRYPOINT ["./main","./env.yaml"]
