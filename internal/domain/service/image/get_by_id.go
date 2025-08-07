package image

import (
	"Leech-ru/internal/domain/dto"
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
)

func (s *imageService) GetByID(ctx context.Context, id uuid.UUID) (*dto.GetByIdImageResponse, error) {
	reader, contentType, size, lastModified, originalFilename, err := s.imageMinIORepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read image content: %w", err)
	}

	filePkg := &dto.FilePackage{
		Content:       bytes.NewReader(buf.Bytes()),
		ContentType:   contentType,
		ContentLength: size,
		Filename:      originalFilename,
		LastModified:  lastModified,
	}

	return &dto.GetByIdImageResponse{
		ID:   id,
		File: filePkg,
	}, nil
}
