FROM golang:1.14.4 as builder

WORKDIR /go/src/github.com/nachiguro1003/otenki/src/servicer

COPY ./src/servicer/go.mod ./
COPY ./src/servicer/go.sum ./

RUN go mod download

COPY ./src/servicer/  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -o /go/bin/main -ldflags '-s -w'

FROM scratch as runner

COPY --from=builder /go/src/github.com/nachiguro1003/otenki/src/servicer/env-prd.yaml ./env.yaml
COPY --from=builder /go/bin/main ./main

EXPOSE 8001

ENTRYPOINT ["./main","./env.yaml"]
