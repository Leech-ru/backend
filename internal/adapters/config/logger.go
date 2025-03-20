package config

import (
	"time"

	"github.com/spf13/viper"
)

type LoggerConfig interface {
	Debug() bool
	LogToFile() bool
	LogsDir() string
	TimeLocation() *time.Location
}

type loggerConfig struct {
	debug        bool
	logToFile    bool
	logsDir      string
	timeLocation *time.Location
}

func NewLoggerConfig() (LoggerConfig, error) {
	location, err := time.LoadLocation(viper.GetString("settings.timezone"))
	if err != nil {
		return nil, err
	}

	return &loggerConfig{
		debug:        viper.GetBool("settings.debug"),
		logToFile:    viper.GetBool("settings.logger.log-to-file"),
		logsDir:      viper.GetString("settings.logger.logs-dir"),
		timeLocation: location,
	}, nil
}

func (cfg *loggerConfig) Debug() bool {
	return cfg.debug
}

func (cfg *loggerConfig) LogToFile() bool {
	return cfg.logToFile
}

func (cfg *loggerConfig) LogsDir() string {
	return cfg.logsDir
}

func (cfg *loggerConfig) TimeLocation() *time.Location {
	return cfg.timeLocation
}
