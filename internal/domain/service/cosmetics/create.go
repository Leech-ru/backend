package cosmetics

import (
	"Leech-ru/internal/domain/dto"
	"Leech-ru/pkg/ent"
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"time"
)

// Create create cosmetics and returns it.
func (s *cosmeticsService) Create(ctx context.Context, req *dto.CreateCosmeticsRequest, file *multipart.FileHeader) (*dto.CreateCosmeticsResponse, error) {
	cosmetics := &ent.Cosmetics{
		Title:             req.Title,
		Description:       req.Description,
		ApplicationMethod: req.ApplicationMethod,
		Volume:            req.Volume,
		IsHidden:          *req.IsHidden,
		Edges: ent.CosmeticsEdges{
			Category: &ent.Category{ID: req.CategoryID},
		},
	}

	if req.Links != nil {
		cosmetics.OzonLink = req.Links.Ozon
		cosmetics.WildberriesLink = req.Links.Wildberries
	}
	var imageID *uuid.UUID
	if file != nil {
		src, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open image: %w", err)
		}
		defer src.Close()

		buf, err := io.ReadAll(src)
		if err != nil {
			return nil, fmt.Errorf("failed to read image: %w", err)
		}

		imageDTO, err := s.imageService.Create(ctx, &dto.CreateImageRequest{
			File: &dto.FilePackage{
				Content:      bytes.NewReader(buf),
				ContentType:  file.Header.Get("Content-Type"),
				Size:         file.Size,
				Filename:     file.Filename,
				LastModified: time.Now(),
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create image: %w", err)
		}
		imageID = &imageDTO.ID
		cosmetics.ImageID = imageID
	}

	cosmetics, err := s.cosmeticsRepo.Create(ctx, *cosmetics)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCosmeticsResponse{
		ID: cosmetics.ID,
		Category: dto.Category{
			ID:   cosmetics.Edges.Category.ID,
			Name: cosmetics.Edges.Category.Name,
		},
		Title:             cosmetics.Title,
		Description:       cosmetics.Description,
		ApplicationMethod: cosmetics.ApplicationMethod,
		Volume:            cosmetics.Volume,
		Links: &dto.Links{
			Ozon:        cosmetics.OzonLink,
			Wildberries: cosmetics.WildberriesLink,
		},
		IsHidden: cosmetics.IsHidden,
		ImageID:  imageID,
	}, nil
}
