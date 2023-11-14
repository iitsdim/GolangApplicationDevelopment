FROM golang:1.20

WORKDIR usr/src/greenlight

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8080

