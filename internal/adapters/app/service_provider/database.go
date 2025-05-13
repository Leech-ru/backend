package service_provider

import (
	"Leech-ru/pkg/closer"
	"Leech-ru/pkg/ent"
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/lib/pq"
)

func (s *ServiceProvider) DB() *ent.Client {
	if s.db == nil {
		s.Logger().Debugf("Connecting to database (dsn=%s)", s.PGConfig().DSN())
		db, err := ent.Open(dialect.Postgres, s.PGConfig().DSN())
		if err != nil {
			s.Logger().Panicf("failed to open database: %v", err)
		}
		fmt.Println(db)
		client := db

		loggerCfg := s.LoggerConfig()
		if loggerCfg.Debug() {
			client = client.Debug()
		}

		if errMigrate := client.Schema.Create(
			context.Background(),
		); errMigrate != nil {
			s.Logger().Panicf("Failed to run migrations: %v", errMigrate)
		}

		closer.Add(client.Close)
		s.db = client
	}

	return s.db
}
