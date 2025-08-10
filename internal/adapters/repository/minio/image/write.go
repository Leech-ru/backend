package image

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"io"
	"strings"
)

// Write saves an image with a given UUID
func (r *imageRepo) Write(
	ctx context.Context,
	id uuid.UUID,
	file io.Reader,
	size int64,
	contentType string,
	originalFilename string,
) error {
	objectName := id.String()

	userMeta := map[string]string{
		"original-filename": strings.ToLower(originalFilename),
	}

	_, err := r.client.PutObject(ctx, r.bucketName, objectName, file, size, minio.PutObjectOptions{
		ContentType:  contentType,
		UserMetadata: userMeta,
	})
	if err != nil {
		return fmt.Errorf("failed to upload image to MinIO (id=%s): %w", objectName, err)
	}

	return nil
}
