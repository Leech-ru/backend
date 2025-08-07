package service_provider

import (
	"Leech-ru/pkg/closer"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func (s *ServiceProvider) MinIO() *minio.Client {
	if s.minio == nil {
		cfg := s.MinIOConfig()
		s.Logger().Debugf("Connecting to MinIO (endpoint=%s)", cfg.Endpoint())

		client, err := minio.New(cfg.Endpoint(), &minio.Options{
			Creds:  credentials.NewStaticV4(cfg.AccessKey(), cfg.SecretKey(), ""),
			Secure: cfg.SSL(),
		})
		if err != nil {
			s.Logger().Panicf("failed to initialize MinIO client: %v", err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout())
		defer cancel()

		_, err = client.ListBuckets(ctx)
		if err != nil {
			s.Logger().Panicf("failed to connect to MinIO: %v", err)
		}

		exists, err := client.BucketExists(ctx, cfg.BucketName())
		if err != nil {
			s.Logger().Panicf("failed to check bucket existence: %v", err)
		}

		if !exists {
			err = client.MakeBucket(ctx, cfg.BucketName(), minio.MakeBucketOptions{})
			if err != nil {
				s.Logger().Panicf("failed to create bucket: %v", err)
			}
			s.Logger().Debugf("Bucket '%s' created successfully", cfg.BucketName())
		}

		closer.Add(func() error {
			s.Logger().Debug("MinIO connection closed")
			return nil
		})

		s.minio = client
	}

	return s.minio
}
