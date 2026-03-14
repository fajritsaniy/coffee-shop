# Build stage
FROM golang:1.22.2-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download && go get github.com/joho/godotenv

# Copy source code
COPY . .

# Update dependencies and build
RUN go mod tidy && go build -o coffee-shop main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/coffee-shop .
COPY --from=builder /app/.env .

# Expose port
EXPOSE 3001

# Command to run
CMD ["./coffee-shop"]
