FROM golang:1.23-alpine AS builder
WORKDIR /app

# Модули и кэш
COPY go.mod go.sum ./
RUN go mod download

# Исходники и синхронизация зависимостей
COPY . .
RUN go mod tidy

# Сборка компактного бинаря без отладочной инфы
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app ./cmd/

FROM alpine:3.20
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app .
COPY --from=builder /app/cmd/favorites.json .
RUN adduser -D -g '' appuser
USER appuser
EXPOSE 8000
CMD ["./app"]

