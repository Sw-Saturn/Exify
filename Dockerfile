FROM golang:1.16.0-alpine3.13 AS build

ENV GO111MODULE=on
RUN apk --no-cache add git make build-base
WORKDIR /go/src/github.com/Sw-Saturn/Image2EXIF-API
COPY . .
RUN go build -a -tags "netgo" -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o app main.go

EXPOSE 8080
ENTRYPOINT ["/go/src/github.com/Sw-Saturn/Image2EXIF-API/app"]