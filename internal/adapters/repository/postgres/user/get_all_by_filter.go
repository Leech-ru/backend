package user

import (
	"Leech-ru/internal/domain/types"
	"Leech-ru/pkg/ent"
	"Leech-ru/pkg/ent/user"
	"context"
	"fmt"
	"strings"
)

// GetAllByFilter retrieves all users with optional pagination and filters.
func (s *userRepo) GetAllByFilter(
	ctx context.Context,
	limit, offset int,
	role *types.Role,
	queryText, emailPrefix *string,
) ([]*ent.User, error) {

	query := s.client.User.Query()

	if role != nil {
		query = query.Where(user.RoleEQ(*role))
	}
	if queryText != nil {
		for _, token := range splitSearchTokens(*queryText) {
			query = query.Where(
				user.Or(
					user.NameContainsFold(token),
					user.SurnameContainsFold(token),
				),
			)
		}
	}
	if emailPrefix != nil {
		query = query.Where(user.EmailHasPrefix(*emailPrefix))
	}
	users, err := query.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to query db: %w", err)
	}

	return users, nil
}

func splitSearchTokens(searchText string) []string {
	normalized := strings.ToLower(searchText)
	return strings.Fields(normalized)
}
