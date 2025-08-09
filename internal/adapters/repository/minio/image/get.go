package image

import (
	"Leech-ru/internal/domain/common/errorz"
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

// Get returns the image by uuid, fully loaded into memory.
// The returned io.Reader is safe to use without Close().
func (s *imageRepo) Get(ctx context.Context, id uuid.UUID) (io.Reader, string, int64, time.Time, string, error) {
	obj, err := s.client.GetObject(ctx, s.bucketName, id.String(), minio.GetObjectOptions{})
	if err != nil {
		return nil, "", 0, time.Time{}, "", fmt.Errorf("failed to get image object: %w", err)
	}
	defer obj.Close()

	info, err := obj.Stat()
	if err != nil {
		fmt.Println("NOT FOUND")
		
		return nil, "", 0, time.Time{}, "", errorz.ImageNotFound
	}

	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, "", 0, time.Time{}, "", fmt.Errorf("failed to read image object: %w", err)
	}

	originalFilename := info.UserMetadata["original-filename"]

	return bytes.NewReader(data), info.ContentType, int64(len(data)), info.LastModified, originalFilename, nil
}
