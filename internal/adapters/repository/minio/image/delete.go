package image

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (r *imageRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.client.RemoveObject(ctx, r.bucketName, id.String(), minio.RemoveObjectOptions{})
}
