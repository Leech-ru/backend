package mainpage

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAll gets all main page contents with pagination
func (s *mainPageService) GetAll(ctx context.Context, req *dto.GetAllMainPageRequest) (dto.GetAllMainPageResponse, error) {
	limit := 10
	if req.Limit != nil && *req.Limit > 0 {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}

	mainPages, totalItems, err := s.mainPageRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dto.GetAllMainPageResponse{}, err
	}

	respItems := make([]*dto.MainPage, len(mainPages))
	for i, page := range mainPages {
		respItems[i] = &dto.MainPage{
			ID:       page.ID,
			ImageID:  page.ImageID,
			Title:    page.Title,
			Content:  page.Content,
			Href:     page.Href,
			IsHidden: page.IsHidden,
			Fluid:    page.Fluid,
		}
	}

	totalPages := 0
	if totalItems > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}
	currentPage := (offset / limit) + 1

	resp := dto.GetAllMainPageResponse{
		Items: respItems,
		Pagination: dto.PaginationInfo{
			TotalItems:  totalItems,
			TotalPages:  totalPages,
			CurrentPage: currentPage,
			HasNext:     currentPage < totalPages,
			HasPrevious: currentPage > 1,
		},
	}

	return resp, nil
}
