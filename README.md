# itk-test-task

## Запуск

1. Скопируй конфиг:
```
cp config.env.example config.env
```

2. Заполни `config.env`

3. Запусти:
```
docker-compose up --build
```

4. Примени миграции:
```
goose -dir iternal/repository/migrations postgres "postgres://user:password@localhost:5432/itk?sslmode=disable" up
```

## Тесты
```
go test ./...
```