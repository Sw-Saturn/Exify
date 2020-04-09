FROM golang:1.14.1-alpine3.11 AS build

ENV GO111MODULE=on
RUN apk --no-cache add git make build-base
WORKDIR /go/src/github.com/Sw-Saturn/Image2EXIF-API
COPY . .
RUN mkdir -p /build
RUN go build -a -tags "netgo" -installsuffix netgo -ldflags="-s -w -extldflags \"-static\"" -o=/build/app main.go

FROM alpine:3.11.5

RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
COPY --from=build /build/app /build/app
RUN chmod u+x /build/app
RUN mkdir -p upload

EXPOSE 8080
ENTRYPOINT ["/build/app"]
