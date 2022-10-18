# syntax=docker/dockerfile:1
# FROM golang:1.17-latest
FROM golang:1.17-alpine

WORKDIR /app

LABEL maintainer="Ashley and Gin"

COPY . .
#COPY go.mod ./
# COPY go.sum ./
RUN go mod download

RUN go build -o /app/ascii-art-web-ctn

EXPOSE 8080

CMD [ "/app/ascii-art-web-ctn" ]