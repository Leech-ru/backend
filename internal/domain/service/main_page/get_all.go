package mainpage

import (
	"Leech-ru/internal/domain/dto"
	"context"
)

// GetAll gets all main page contents with pagination
func (s *mainPageService) GetAll(ctx context.Context, req *dto.GetAllMainPageRequest) (dto.GetAllMainPageResponse, error) {
	mainPages, err := s.mainPageRepo.GetAll(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	resp := make([]*dto.MainPage, len(mainPages))
	for i, page := range mainPages {
		resp[i] = &dto.MainPage{
			ID:       page.ID,
			ImageID:  page.ImageID,
			Title:    page.Title,
			Content:  page.Content,
			Href:     page.Href,
			IsHidden: page.IsHidden,
			Fluid:    page.Fluid,
		}
	}

	return resp, nil
}
