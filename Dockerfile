FROM golang:1.16.0-alpine3.13 AS build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk --no-cache add git make build-base
WORKDIR /go/src/github.com/Sw-Saturn/Exify
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app main.go

FROM alpine:3.13

WORKDIR /go/src/github.com/Sw-Saturn/Exify
RUN apk --no-cache add ca-certificates
COPY --from=build /go/src/github.com/Sw-Saturn/Exify/ .
RUN mkdir -p ./upload

ENTRYPOINT ["/go/src/github.com/Sw-Saturn/Exify/app"]