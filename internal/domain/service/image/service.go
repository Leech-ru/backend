package image

import (
	imageMinIO "Leech-ru/internal/adapters/repository/minio/image"
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

type imageMinIORepo interface {
	Write(
		ctx context.Context,
		id uuid.UUID,
		file io.Reader,
		size int64,
		contentType string,
		originalFilename string,
	) error
	Get(ctx context.Context, id uuid.UUID) (io.Reader, string, int64, time.Time, string, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
}
type minioConfig interface {
	BucketName() string
}

type imageService struct {
	imageMinIORepo imageMinIORepo
}

func NewImageService(minIOClient *minio.Client, minioConfig minioConfig) *imageService {
	return &imageService{
		imageMinIORepo: imageMinIO.NewImageRepo(minIOClient, minioConfig.BucketName()),
	}
}
