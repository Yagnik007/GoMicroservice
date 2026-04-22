# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum* ./

# Download dependencies
# (uncomment after you have go.sum in your local directory)
# RUN go mod download 

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .
# Copy the env file
COPY .env.example .env

# Expose the API port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
