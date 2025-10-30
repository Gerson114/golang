# Build stage
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

# Install git for private modules
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags='-w -s -extldflags "-static"' -o main .

# Production stage
FROM scratch

# Copy SSL certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary
COPY --from=builder /app/main /main

# Expose port
EXPOSE 8080

# Run binary
CMD ["/main"]
