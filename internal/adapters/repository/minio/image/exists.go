package image

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

// Exists checks if an image with the given UUID exists in the bucket.
// Returns true if the object exists, false if not found, and error if any other issue occurs.
func (r *imageRepo) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	_, err := r.client.StatObject(ctx, r.bucketName, id.String(), minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		}
		return false, fmt.Errorf("failed to stat image object: %w", err)
	}
	return true, nil
}
