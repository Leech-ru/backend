package image

import (
	"Leech-ru/internal/domain/dto"
	"bytes"
	"context"
	"fmt"
	"io"
)

func (s *imageService) GetById(ctx context.Context, req *dto.GetByIdImageRequest) (*dto.GetByIdImageResponse, error) {
	reader, contentType, size, lastModified, originalFilename, err := s.imageMinIORepo.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read image content: %w", err)
	}

	return &dto.GetByIdImageResponse{
		ID: req.ID,
		File: &dto.FilePackage{
			Content:      bytes.NewReader(buf.Bytes()), // теперь это новый reader с началом потока
			ContentType:  contentType,
			Size:         size,
			Filename:     originalFilename,
			LastModified: lastModified,
		},
	}, nil
}
