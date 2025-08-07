package config

import (
	"github.com/spf13/viper"
	"time"
)

type MinIOConfig interface {
	Endpoint() string
	AccessKey() string
	SecretKey() string
	BucketName() string
	SSL() bool
	Timeout() time.Duration
}

type minioConfig struct {
	endpoint   string
	accessKey  string
	secretKey  string
	bucketName string
	ssl        bool
	timeout    time.Duration
}

func NewMinIOConfig() (MinIOConfig, error) {
	return &minioConfig{
		endpoint:   viper.GetString("service.minio.endpoint"),
		accessKey:  viper.GetString("service.minio.access-key"),
		secretKey:  viper.GetString("service.minio.secret-key"),
		bucketName: viper.GetString("service.minio.bucket-name"),
		ssl:        viper.GetBool("service.minio.ssl"),
		timeout:    viper.GetDuration("service.minio.timeout") * time.Second,
	}, nil
}

func (cfg *minioConfig) Endpoint() string {
	return cfg.endpoint
}

func (cfg *minioConfig) AccessKey() string {
	return cfg.accessKey
}

func (cfg *minioConfig) SecretKey() string {
	return cfg.secretKey
}

func (cfg *minioConfig) BucketName() string {
	return cfg.bucketName
}

func (cfg *minioConfig) SSL() bool {
	return cfg.ssl
}

func (cfg *minioConfig) Timeout() time.Duration {
	return cfg.timeout
}
