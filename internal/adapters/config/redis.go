package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type RedisConfig interface {
	Addr() string
	Password() string
	DB() int
}

type redisConfig struct {
	host     string
	port     int
	password string
	db       int
}

func NewRedisConfig() (RedisConfig, error) {
	return &redisConfig{
		host:     viper.GetString("service.redis.host"),
		port:     viper.GetInt("service.redis.port"),
		password: viper.GetString("service.redis.password"),
		db:       viper.GetInt("service.redis.db"),
	}, nil
}

func (cfg *redisConfig) Addr() string {
	return fmt.Sprintf("%s:%d", cfg.host, cfg.port)
}

func (cfg *redisConfig) Password() string {
	return cfg.password
}

func (cfg *redisConfig) DB() int {
	return cfg.db
}
