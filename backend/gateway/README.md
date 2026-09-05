# API Gateway

Gateway на Go и Echo. Он принимает HTTP-запросы и вызывает соответствующие gRPC-методы `theatre-service`.

## Запуск

```powershell
go run ./cmd
```

Путь к конфигурации можно изменить переменной `CONFIG_PATH`:

```powershell
$env:CONFIG_PATH = "config/config.yaml"
go run ./cmd
```

Проверка доступности:

```text
GET http://localhost:8080/health
GET http://localhost:8080/ready
```

## Конфигурация

Адрес theatre-service задаётся в `config/config.yaml`:

```yaml
theatre_service:
  address: "localhost:9090"
```

Поддерживаются следующие HTTP-маршруты:

```text
GET /api/spectacles?page=0&size=20
GET /api/spectacles/{id}
GET /api/shows/{id}
```

Gateway не хранит данные предметной области и не генерирует mock-ответы. Protobuf-клиент сгенерирован из `theatre-service/src/main/proto/spectacle.proto`.

## Структура

- `cmd/` - запуск приложения и сборка зависимостей.
- `internal/config/` - загрузка и валидация YAML.
- `internal/controller/` - HTTP handlers Echo и преобразование ошибок.
- `internal/service/` - порт приложения для theatre-service.
- `internal/adapter/theatregrpc/` - gRPC-адаптер и преобразование protobuf в domain.
- `internal/domain/` - доменные модели.
- `config/` - пример конфигурации.

## Docker

```powershell
docker build -t gateway .
docker run --rm -p 8080:8080 gateway
```
