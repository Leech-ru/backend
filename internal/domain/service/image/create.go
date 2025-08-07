package image

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"context"
	"fmt"
	"mime/multipart"
	"time"
)

func (s *imageService) Create(ctx context.Context, req *dto.CreateImageRequest, file *multipart.FileHeader) (*dto.CreateImageResponse, error) {
	image := &ent.Image{Name: req.File.Filename}
	image, err := s.imagePostgresRepo.Create(ctx, *image)
	if err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	err = s.imageMinIORepo.Write(ctx, image.ID, src, file.Size, file.Header.Get("Content-Type"), file.Filename)
	if err != nil {
		_ = s.imagePostgresRepo.Delete(ctx, image.ID)
		return nil, fmt.Errorf("failed to upload image to storage: %w", err)
	}

	return &dto.CreateImageResponse{
		ID: image.ID,
		File: &dto.FilePackage{
			Content:       nil,
			ContentType:   file.Header.Get("Content-Type"),
			ContentLength: file.Size,
			Filename:      file.Filename,
			LastModified:  time.Now(),
		},
	}, nil
}
