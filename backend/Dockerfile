# Stage 1: Build the application
FROM golang:1.21 AS builder

WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Stage 2: Run the application
FROM alpine:latest AS runtime

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./main"]