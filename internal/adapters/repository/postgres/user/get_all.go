package user

import (
	"Leech-ru/pkg/ent"
	"context"
)

// GetAll retrieves all users with optional pagination
func (s *userRepo) GetAll(ctx context.Context, limit, offset int) ([]*ent.User, error) {
	users, err := s.client.User.
		Query().
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
