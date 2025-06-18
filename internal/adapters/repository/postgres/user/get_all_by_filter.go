package user

import (
	"Leech-ru/internal/domain/types"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/user"
	"context"
)

// GetAllByFilter retrieves all users with optional pagination and filters.
func (s *userRepo) GetAllByFilter(
	ctx context.Context,
	limit, offset int,
	role *types.Role,
	namePrefix, surnamePrefix, emailPrefix *string,
) ([]*ent.User, error) {

	query := s.client.User.Query()

	if role != nil {
		query = query.Where(user.RoleEQ(*role))
	}
	if namePrefix != nil {
		query = query.Where(user.NameHasPrefix(*namePrefix))
	}
	if surnamePrefix != nil {
		query = query.Where(user.SurnameHasPrefix(*surnamePrefix))
	}
	if emailPrefix != nil {
		query = query.Where(user.EmailHasPrefix(*emailPrefix))
	}

	users, err := query.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return users, nil
}
