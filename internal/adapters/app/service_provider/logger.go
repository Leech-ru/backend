package service_provider

import (
	"Leech-ru/pkg/logger"
	"fmt"
)

func (s *ServiceProvider) Logger() *logger.Logger {
	if s.logger == nil {
		l, err := logger.Init(logger.Config{
			Debug:        s.LoggerConfig().Debug(),
			TimeLocation: s.LoggerConfig().TimeLocation(),
			LogToFile:    s.LoggerConfig().LogToFile(),
			LogsDir:      s.LoggerConfig().LogsDir(),
		})
		if err != nil {
			panic(fmt.Errorf("failed to init logger: %w", err))
		}

		s.logger = l
		if s.LoggerConfig().Debug() {
			s.logger.Debug("Debug mode enabled")
		}
	}

	return s.logger
}
