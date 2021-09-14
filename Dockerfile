FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

ENV PORT 8081

RUN go build -o /docker-gs-ping ./cmd/api


CMD [ "/docker-gs-ping" ]
