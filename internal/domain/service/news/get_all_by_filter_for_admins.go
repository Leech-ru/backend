package news

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAllByFilterForAdmins realizes a search for news with filtering parameters, but only for admins.
func (s *newsService) GetAllByFilterForAdmins(ctx context.Context, req *dto.GetAllByFilterForAdminsNewsRequest) (*dto.GetAllByFilterForAdminsNewsResponse, error) {
	limit := 10
	if req.Limit != nil && *req.Limit > 0 {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}
	allNews, totalItems, err := s.newsRepo.GetAllByFilter(ctx, limit, offset, req.IsHidden)
	if err != nil {
		return nil, err
	}

	previewLen := basePreviewLength
	if req.PreviewLength != nil {
		previewLen = *req.PreviewLength
	}

	respItems := make([]*dto.NewsListItem, 0, len(allNews))
	for _, news := range allNews {
		contentPreview := news.Content
		if len(contentPreview) > previewLen {
			contentPreview = contentPreview[:previewLen]
		}

		respItems = append(respItems, &dto.NewsListItem{
			ID:             news.ID,
			ImageID:        news.ImageID,
			Title:          news.Title,
			ContentPreview: contentPreview,
			IsHidden:       news.IsHidden,
		})
	}

	totalPages := 0
	if totalItems > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}
	currentPage := (offset / limit) + 1

	resp := dto.GetAllByFilterForAdminsNewsResponse{
		Items: respItems,
		Pagination: dto.PaginationInfo{
			TotalItems:  totalItems,
			TotalPages:  totalPages,
			CurrentPage: currentPage,
			HasNext:     currentPage < totalPages,
			HasPrevious: currentPage > 1,
		},
	}

	return &resp, nil
}
