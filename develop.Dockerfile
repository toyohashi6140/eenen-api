FROM golang:1.21.3-alpine

WORKDIR /go/src/github.com/toyohashi6140/eenen-api
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git alpine-sdk

RUN set -x \ 
    && go mod download \
    && go install github.com/cosmtrek/air@latest \
    && go install github.com/99designs/gqlgen@latest

COPY . .