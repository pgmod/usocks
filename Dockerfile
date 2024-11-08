# Используем официальный образ Go для сборки
FROM golang:1.23-bookworm AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем проект
RUN go build -o socks ./cmd/main/
# Создаём финальный минималистичный образ
FROM scratch

COPY --from=builder /app/socks /socks


# Определяем команду по умолчанию
CMD ["/socks"]
