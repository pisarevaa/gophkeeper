# gophkeeper
GophKeeper представляет собой клиент-серверную систему, позволяющую пользователю надёжно и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.

# Шаги по запуску сервиса
1. Скопируйте и измените переменные окружения:
```
cp .env.example .env
```
2. Сгенерите публичный и приватный ключ:
```
make generate_keys
```
3. Запустите БД и Minio
```
make up_database
```
4. Запустите сервер
```
make run_server
```

# Swagger документация

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

# Для разработчиков
Запус линтеров
```
make lint
```
Запуск тестов
```
make test
```
Генерация swagger документации
```
make swag_format && make swag_generate
```
Генерация моков для тестов
```
make mock_generate
```