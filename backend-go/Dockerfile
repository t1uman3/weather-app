FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o app ./cmd/

# Use a smaller base image for the final container
FROM alpine:3.19

WORKDIR /app

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/cmd/favorites.json .

# Create a non-root user to run the application
RUN adduser -D -g '' appuser
USER appuser

# Expose the application port
EXPOSE 8000

# Command to run the application
CMD ["./app"]
