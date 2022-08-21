# syntax=docker/dockerfile:1

FROM golang:1.18-alpine as builder

ARG BUILD_PACKAGE

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY pkg ./pkg
COPY cmd/$BUILD_PACKAGE cmd/$BUILD_PACKAGE

WORKDIR /app/cmd/$BUILD_PACKAGE
RUN go build -o /tmp/app

FROM alpine:3.16

WORKDIR /
COPY --from=builder /tmp/app ./

EXPOSE 8000

CMD ["/app"]

LABEL version="0.0.0"
LABEL org.opencontainers.image.authors="abashinos@gmail.com"