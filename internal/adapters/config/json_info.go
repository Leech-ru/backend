package config

import "github.com/spf13/viper"

// JsonInfoConfig defines an interface for HTTP server configuration.
type JsonInfoConfig interface {
	PathToJsonFile() string
}

type jsonInfoConfig struct {
	pathToJsonFile string
}

// NewJsonInfoConfig initializes a new JSON info configuration from environment variables.
func NewJsonInfoConfig() (JsonInfoConfig, error) {
	return &jsonInfoConfig{
		pathToJsonFile: viper.GetString("service.info.path-to-json-file"),
	}, nil
}

func (cfg *jsonInfoConfig) PathToJsonFile() string {
	return cfg.pathToJsonFile
}
