package image

import (
	imageMinIO "Leech-ru/internal/adapters/repository/minio/image"
	imagePostges "Leech-ru/internal/adapters/repository/postgres/image"
	"Leech-ru/pkg/ent"
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

type imagePostgresRepo interface {
	Create(ctx context.Context, entity ent.Image) (*ent.Image, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.Image, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type imageMinIORepo interface {
	Write(
		ctx context.Context,
		id uuid.UUID,
		file io.Reader,
		size int64,
		contentType string,
		originalFilename string,
	) error
	Get(ctx context.Context, id uuid.UUID) (io.ReadCloser, string, int64, time.Time, string, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
type minioConfig interface {
	BucketName() string
}

type imageService struct {
	imagePostgresRepo imagePostgresRepo
	imageMinIORepo    imageMinIORepo
}

func NewImageService(entClient *ent.Client, minIOClient *minio.Client, minioConfig minioConfig) *imageService {
	return &imageService{
		imagePostgresRepo: imagePostges.NewImageRepo(entClient),
		imageMinIORepo:    imageMinIO.NewImageRepo(minIOClient, minioConfig.BucketName()),
	}
}
