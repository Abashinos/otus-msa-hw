# syntax=docker/dockerfile:1

FROM golang:1.18-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./server

FROM alpine:3.16

WORKDIR /
COPY --from=builder /app/server ./

EXPOSE 8000

CMD ["/server"]

LABEL version="0.0.0"
LABEL org.opencontainers.image.authors="abashinos@gmail.com"