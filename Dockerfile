FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

RUN apk add --no-cache bash

COPY . .

RUN go build -o /docker-gs-ping ./cmd/api

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
