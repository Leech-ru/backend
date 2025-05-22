package service_provider

import (
	"Leech-ru/pkg/closer"
	"context"
	"github.com/redis/go-redis/v9"
)

func (s *ServiceProvider) Redis() *redis.Client {
	if s.redis == nil {
		s.Logger().Debugf("Connecting to Redis (addr=%s, db=%d)", s.RedisConfig().Addr(), s.RedisConfig().DB())

		client := redis.NewClient(&redis.Options{
			Addr:     s.RedisConfig().Addr(),
			Password: s.RedisConfig().Password(),
			DB:       s.RedisConfig().DB(),
		})

		// Проверим соединение
		if err := client.Ping(context.Background()).Err(); err != nil {
			s.Logger().Panicf("Failed to connect to Redis: %v", err)
		}

		// Добавим в graceful closer
		closer.Add(func() error {
			s.Logger().Info("Closing Redis connection")
			return client.Close()
		})

		s.redis = client
	}

	return s.redis
}
