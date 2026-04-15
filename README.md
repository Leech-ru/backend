# Leech-ru Backend

Бэкенд для платформы Leech-ru.

Что делает сервис:
- хранит и отдает данные для сайта;
- работает с пользователями и авторизацией;
- управляет категориями, косметикой, новостями, партнерами и контентом главной;
- хранит изображения в MinIO;
- использует PostgreSQL и Valkey.

Если вы разработчик, подробная техническая документация здесь: [DEV.md](DEV.md).

## Быстрый запуск через Docker

### 1. Подготовить окружение

Скопируйте `.env.example` в `.env`:

```bash
cp .env.example .env
```

Скопируйте `config.yaml.example` в `config.yaml`:

```bash
cp config.yaml.example config.yaml
```

При необходимости отредактируйте значения в `.env` и `config.yaml`.

### 2. Сгенерировать RSA-ключи (обязательно)

Сервис использует JWT с RSA-ключами. До сборки контейнера создайте ключи:

```bash
openssl genpkey -algorithm RSA -out keys/private.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in keys/private.pem -out keys/public.pem
```

### 3. Запустить проект

```bash
docker compose up --build -d
```

### 4. Проверить, что сервис поднялся

API: `http://localhost:8000`  
Swagger: `http://localhost:8000/api/v1/swagger/index.html`

Проверка статуса контейнеров:

```bash
docker compose ps
```

Логи приложения:

```bash
docker compose logs -f app
```

### 5. Остановить проект

```bash
docker compose down
```

