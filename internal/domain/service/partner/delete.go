package partner

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/internal/domain/dto"
	"context"
	"errors"
)

// Delete delete partner by ID.
func (s *partnerService) Delete(ctx context.Context, req *dto.DeletePartnerRequest) error {
	err := s.partnerRepo.Delete(ctx, req.ID)
	switch {
	case errors.Is(err, errorz.PartnerNotFound):
		return errorz.PartnerNotFound
	case err != nil:
		return err
	}

	return nil
}
