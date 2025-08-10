package image

import (
	"context"
	"github.com/google/uuid"
)

func (s *imageService) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	return s.imageMinIORepo.Exists(ctx, id)
}
