package image

import (
	"Leech-ru/internal/domain/common/errorz"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

// Get returns the image by uuid.
func (s *imageRepo) Get(ctx context.Context, id uuid.UUID) (io.ReadCloser, string, int64, time.Time, string, error) {
	obj, err := s.client.GetObject(ctx, s.bucketName, id.String(), minio.GetObjectOptions{})
	if err != nil {
		return nil, "", 0, time.Time{}, "", fmt.Errorf("failed to get image object: %w", err)
	}

	info, err := obj.Stat()
	if err != nil {
		return nil, "", 0, time.Time{}, "", errorz.ImageNotFound
	}

	originalFilename := info.UserMetadata["original-filename"]

	return obj, info.ContentType, info.Size, info.LastModified, originalFilename, nil
}
