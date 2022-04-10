#build stage
FROM golang:1.17.6-alpine AS builder
LABEL name=go-mysql Version=0.0.2 maintainer="Juan Farber <jfarberjob@gmail.com>"
WORKDIR /go/src/app
RUN apk add --no-cache git
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin /app

WORKDIR /app
CMD ./go-mysql
EXPOSE 1323