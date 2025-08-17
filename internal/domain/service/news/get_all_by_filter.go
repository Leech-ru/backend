package news

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAllByFilter realizes a search for news with filtering parameters.
func (s *newsService) GetAllByFilter(ctx context.Context, req *dto.GetAllByFilterNewsRequest) (*dto.GetAllByFilterNewsResponse, error) {
	limit := 10
	if req.Limit != nil {
		limit = *req.Limit
	}
	offset := 0
	if req.Offset != nil {
		offset = *req.Offset
	}
	isHidden := false
	allNews, err := s.newsRepo.GetAllByFilter(ctx, limit, offset, &isHidden)
	if err != nil {
		return nil, err
	}

	previewLen := basePreviewLength
	if req.PreviewLength != nil {
		previewLen = *req.PreviewLength
	}

	var resp dto.GetAllByFilterNewsResponse
	for _, news := range allNews {
		contentPreview := news.Content
		if len(contentPreview) > previewLen {
			contentPreview = contentPreview[:previewLen]
		}

		resp = append(resp, &dto.NewsListItem{
			ID:             news.ID,
			ImageID:        news.ImageID,
			Title:          news.Title,
			ContentPreview: contentPreview,
			Href:           news.Href,
			IsHidden:       news.IsHidden,
		})
	}
	return &resp, nil
}
