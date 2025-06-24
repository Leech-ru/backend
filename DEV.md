Генерация таблиц
```shell
go run entgo.io/ent/cmd/ent new --target internal/domain/schema [TableName]
```

Генерация орм
```shell
go run -mod=mod entgo.io/ent/cmd/ent generate --target ./pkg/ent ./internal/domain/schema
```

Сгенерировать RSA ключи
```shell
openssl genpkey -algorithm RSA -out keys/private.pem -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in keys/private.pem -out keys/public.pem
```