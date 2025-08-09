package image

import (
	"Leech-ru/internal/domain/dto"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (s *imageService) Create(ctx context.Context, req *dto.CreateImageRequest) (*dto.CreateImageResponse, error) {
	id := uuid.New()
	err := s.imageMinIORepo.Write(ctx, id, req.File.Content, req.File.Size, req.File.ContentType, req.File.Filename)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image to storage: %w", err)
	}

	return &dto.CreateImageResponse{
		ID: id,
		File: &dto.FilePackage{
			Content:      nil,
			ContentType:  req.File.ContentType,
			Size:         req.File.Size,
			Filename:     req.File.Filename,
			LastModified: time.Now(),
		},
	}, nil
}
