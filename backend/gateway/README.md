# API Gateway

Reverse proxy на Go и Echo. Gateway принимает запросы, выбирает upstream по самому длинному совпавшему префиксу и при необходимости убирает префикс перед отправкой запроса.

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

Маршрут состоит из `prefix`, `target` и `strip_prefix`:

```yaml
routes:
  - prefix: /api/users
    target: http://localhost:8081
    strip_prefix: true
```

При `strip_prefix: true` запрос `/api/users/42` уйдёт в upstream как `/42`. Секреты и адреса окружений не должны храниться в репозитории: для них используйте переменные окружения или отдельный конфигурационный файл.

## Структура

- `cmd/` - запуск приложения и сборка зависимостей.
- `internal/config/` - загрузка и валидация YAML.
- `internal/controller/` - HTTP handlers Echo.
- `internal/service/` - маршрутизация и reverse proxy.
- `internal/domain/` - доменные модели.
- `config/` - пример конфигурации.

## Docker

```powershell
docker build -t gateway .
docker run --rm -p 8080:8080 gateway
```
