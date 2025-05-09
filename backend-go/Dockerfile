# Этап сборки
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o weather-server ./cmd/server

# Финальный этап
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/weather-server .
COPY .env .
EXPOSE 8000
CMD ["./weather-server"]
