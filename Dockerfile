FROM golang:latest

RUN apt-get update && apt-get install -y netcat-openbsd
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN chmod a+x /app/docker/*.sh

RUN go build -o main .