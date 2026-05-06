#Этап 1: Сборка
FROM golang:1.26.2-alpine AS builder

#Рабочая дериктория
WORKDIR /app

#Копирование файлов зависимостей
COPY go.mod go.sum ./
RUN go mod download

#Копирование всего исходного кода
COPY . .

#Компиляция бинарного файла
RUN CGO_ENABLED=0 GOOS=linux go build -o inspector ./cmd/api/main.go

#ЭТАП 2: Запуск
FROM alpine:3.19

#Установка сертификатов для HTTPS запросов
RUN apk --no-cache add ca-certificates

WORKDIR /root/

#Копирование скомпилированного файла из первотого этапа
COPY --from=builder /app/inspector .

#Открытие порта 
EXPOSE 8080


#Запуск приложения
CMD [ "./inspector" ]
