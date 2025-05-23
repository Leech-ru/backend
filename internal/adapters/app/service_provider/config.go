package service_provider

import (
	"Leech-ru/internal/adapters/config"
	"fmt"
)

func (s *ServiceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get logger config: %w", err))
		}
		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *ServiceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		var err error
		s.pgConfig, err = config.NewPGConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get PG config: %w", err))
		}
	}

	return s.pgConfig
}

func (s *ServiceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		var err error
		s.redisConfig, err = config.NewRedisConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get redis config: %w", err))
		}
	}

	return s.redisConfig
}

func (s *ServiceProvider) ServerConfig() config.ServerConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get http config: %w", err))
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}
func (s *ServiceProvider) MailConfig() config.MailConfig {
	if s.mailConfig == nil {
		cfg, err := config.NewMailConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get mail config: %w", err))
		}

		s.mailConfig = cfg
	}

	return s.mailConfig
}

func (s *ServiceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		cfg, err := config.NewJWTConfig()
		if err != nil {
			panic(fmt.Errorf("failed to get JWT config: %w", err))
		}

		s.jwtConfig = cfg
	}

	return s.jwtConfig
}
