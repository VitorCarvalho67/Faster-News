FROM golang:latest

WORKDIR /app/server

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go application
RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]