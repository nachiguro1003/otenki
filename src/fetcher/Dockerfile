FROM golang:1.12.6-alpine3.9
LABEL maintainer=yoseiito
WORKDIR /go/src/github.com/nachiguro1003/otenki/src/fetcher
ENV GO111MODULE=on CGO_ENABLED=0
RUN apk add --no-cache \
        git \
        && go get github.com/pilu/fresh
CMD ["/bin/true"]
