FROM golang:1.17.6-alpine
LABEL Name=go-mysql Version=0.0.1 maintainer="Juan Farber <juanfarberjob@gmail.com>"
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/go-mysql
#final stage
FROM alpine:latest  
WORKDIR /app
COPY --from=builder /build/go-mysql /app/go-mysql
EXPOSE 1323
CMD ["/app/go-mysql"]
# docker build . -t go-mysql
