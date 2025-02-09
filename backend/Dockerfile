FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy

COPY . .
RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
