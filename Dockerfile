FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD *.go go.mod go.sum /build/
WORKDIR /build
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o us-api .

FROM alpine:3.11.3
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/us-api .
ENTRYPOINT ["./us-api"]
