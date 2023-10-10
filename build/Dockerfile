# Dockerfile

FROM golang:latest

RUN apt-get update && apt-get install -y netcat-openbsd

COPY ../ /app

WORKDIR /app

RUN chmod a+x /app/scripts/*.sh

RUN go build -o main ./cmd/app/.
