Генерация таблиц
```shell
go run entgo.io/ent/cmd/ent new --target internal/domain [TableName]
```

Генерация орм
```shell
go run -mod=mod entgo.io/ent/cmd/ent generate .\internal\domain\schema\ --target ./pkg/ent
```