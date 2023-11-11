FROM golang:1.20

WORKDIR usr/src/greenlight

COPY . .

RUN go mod download

RUN go mod tidy

