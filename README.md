# SSO (Single Sign-On) Service

Сервис аутентификации и авторизации на базе gRPC, реализующий единую точку входа для микросервисной архитектуры.

## Возможности

- Регистрация пользователей
- Аутентификация пользователей
- Управление JWT токенами
- gRPC API для интеграции с другими сервисами
- SQLite в качестве хранилища данных

## Требования

- Go 1.23.4 или выше
- SQLite3
- Make (для использования Makefile)

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd sso
```

2. Установите зависимости:
```bash
go mod download
```

3. Примените миграции базы данных:
```bash
make migrate-up
```

4. Запустите сервис:
```bash
make run
```

## Интеграция с другими проектами

### 1. Установка зависимостей

В вашем проекте добавьте следующие зависимости:
```bash
go get github.com/GolangLessons/protos
go get google.golang.org/grpc
```

### 2. Настройка клиента

```go
package main

import (
    "context"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    pb "github.com/GolangLessons/protos"
)

func main() {
    // Подключение к SSO сервису
    conn, err := grpc.Dial("localhost:44044", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Создание клиента
    client := pb.NewAuthClient(conn)

    // Пример регистрации пользователя
    registerResp, err := client.Register(context.Background(), &pb.RegisterRequest{
        Email:    "user@example.com",
        Password: "password123",
    })
    if err != nil {
        log.Fatalf("Failed to register: %v", err)
    }

    // Пример аутентификации
    loginResp, err := client.Login(context.Background(), &pb.LoginRequest{
        Email:    "user@example.com",
        Password: "password123",
    })
    if err != nil {
        log.Fatalf("Failed to login: %v", err)
    }

    // Использование полученного токена
    token := loginResp.Token
}
```

### 3. Проверка токена

```go
// Пример проверки токена
isValidResp, err := client.IsAdmin(context.Background(), &pb.IsAdminRequest{
    UserId: userId,
})
if err != nil {
    log.Fatalf("Failed to check admin status: %v", err)
}

if isValidResp.IsAdmin {
    // Пользователь является администратором
}
```

## Конфигурация

Сервис использует конфигурационный файл в формате TOML. Пример конфигурации:

```toml
grpc_port = 44044
token_ttl = "1h"
```

## Тестирование

Для запуска тестов используйте:
```bash
make test
```

## Лицензия

MIT

