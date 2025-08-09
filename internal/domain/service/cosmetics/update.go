package cosmetics

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
	"mime/multipart"
	"time"
)

// Update change cosmetics data and optionally upload a new image file.
func (s *cosmeticsService) Update(ctx context.Context, req *dto.UpdateCosmeticsRequest, file *multipart.FileHeader) (*dto.UpdateCosmeticsResponse, error) {
	cosmeticToUpdate, err := s.cosmeticsRepo.GetById(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.CosmeticsNotFound):
		return nil, errorz.CosmeticsNotFound
	case err != nil:
		return nil, err
	}

	if file != nil {
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		imageDTO, err := s.imageService.Create(ctx, &dto.CreateImageRequest{
			File: &dto.FilePackage{
				Content:      src,
				ContentType:  file.Header.Get("Content-Type"),
				Size:         file.Size,
				Filename:     file.Filename,
				LastModified: time.Now(),
			},
		})
		if err != nil {
			return nil, err
		}

		if cosmeticToUpdate.ImageID != nil {
			err := s.imageService.Delete(ctx, &dto.DeleteImageRequest{
				ID: *cosmeticToUpdate.ImageID,
			})
			switch {
			case errors.Is(err, errorz.ImageNotFound): //ignore
			case err != nil:
				return nil, err
			}
		}

		cosmeticToUpdate.ImageID = &imageDTO.ID
	}

	if req.CategoryID != nil {
		cosmeticToUpdate.Edges.Category.ID = *req.CategoryID
	}
	if req.Title != nil {
		cosmeticToUpdate.Title = *req.Title
	}
	if req.Description != nil {
		cosmeticToUpdate.Description = req.Description
	}
	if req.ApplicationMethod != nil {
		cosmeticToUpdate.ApplicationMethod = req.ApplicationMethod
	}
	if req.Volume != nil {
		cosmeticToUpdate.Volume = req.Volume
	}
	if req.IsHidden != nil {
		cosmeticToUpdate.IsHidden = *req.IsHidden
	}
	if req.Links != nil {
		if req.Links.Ozon != nil {
			cosmeticToUpdate.OzonLink = req.Links.Ozon
		}
		if req.Links.Wildberries != nil {
			cosmeticToUpdate.WildberriesLink = req.Links.Wildberries
		}
	}

	updatedCosmetic, err := s.cosmeticsRepo.Update(ctx, *cosmeticToUpdate)
	switch {
	case errors.Is(err, errorz.InvalidCosmeticsFormat):
		return nil, errorz.InvalidCosmeticsFormat
	case errors.Is(err, errorz.CategoryNotFound):
		return nil, errorz.CategoryNotFound
	case errors.Is(err, errorz.ImageNotFound):
		return nil, errorz.ImageNotFound
	case err != nil:
		return nil, err
	}

	return &dto.UpdateCosmeticsResponse{
		ID:      updatedCosmetic.ID,
		ImageID: updatedCosmetic.ImageID,
		Category: dto.Category{
			ID:   updatedCosmetic.Edges.Category.ID,
			Name: updatedCosmetic.Edges.Category.Name,
		},
		Title:             updatedCosmetic.Title,
		Description:       updatedCosmetic.Description,
		ApplicationMethod: updatedCosmetic.ApplicationMethod,
		Volume:            updatedCosmetic.Volume,
		Links: &dto.Links{
			Ozon:        updatedCosmetic.OzonLink,
			Wildberries: updatedCosmetic.WildberriesLink,
		},
		IsHidden: updatedCosmetic.IsHidden,
	}, nil
}
