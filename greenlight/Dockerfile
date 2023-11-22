FROM golang:1.20

WORKDIR usr/src/greenlight

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./
COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN chmod +x wait-for-postgres.sh

RUN go mod download && go mod verify

RUN go build -o main ./cmd/api/

EXPOSE 8080
EXPOSE 5432

# Command to run the executable
CMD ["./main"]