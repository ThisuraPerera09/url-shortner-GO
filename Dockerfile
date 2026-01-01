# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application (no CGO needed - using PostgreSQL or in-memory)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-shortener .

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/url-shortener .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./url-shortener"]
