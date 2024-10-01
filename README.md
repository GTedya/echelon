# Тестовое задание для Эшелон Технологии

Сервис для загрузки миниатюр с видеороликов YouTube.
Кэшируем данные с помощью SQLite (Коннект базы данных обязательный).

Перед началом работы необходимо проинициализировать файлы конфигураций.

Примеры файлов находятся в `config/*.yaml`:

   ```yaml
  grpcPort: ":50051"
  dbPath: "path_to_your_database.db"
   ```

## Запуск

Запуск gRPC сервера:

   ```bash
   go run cmd/server/main.go
   ```

## Клиентская утилита

Утилита командной строки

### Использование:

Синхронная загрузка:

   ```bash
   go run cmd/client/main.go --async=false  "https://www.youtube.com/watch?v=-gYpCIbZjUQ&list=LL&index=23" "https://www.youtube.com/watch?v=rhjiANJVR6g&list=LL&index=24"
   ```

Асинхронная загрузка:

   ```bash
   go run cmd/client/main.go --async=true "https://www.youtube.com/watch?v=rhjiANJVR6g&list=LL&index=24" "https://www.youtube.com/watch?v=-gYpCIbZjUQ&list=LL&index=23"
   ```

## Тестирование

Выполнить тестирование:

```bash
go test ./...
```

