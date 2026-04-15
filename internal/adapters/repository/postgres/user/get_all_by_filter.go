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
) ([]*ent.User, int, error) {

	baseQuery := s.client.User.Query()

	if role != nil {
		baseQuery = baseQuery.Where(user.RoleEQ(*role))
	}
	if queryText != nil {
		for _, token := range splitSearchTokens(*queryText) {
			baseQuery = baseQuery.Where(
				user.Or(
					user.NameContainsFold(token),
					user.SurnameContainsFold(token),
					user.EmailContainsFold(token),
				),
			)
		}
	}
	if emailPrefix != nil {
		baseQuery = baseQuery.Where(user.EmailHasPrefix(*emailPrefix))
	}

	totalItems, err := baseQuery.Clone().Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count users in db: %w", err)
	}

	users, err := baseQuery.
		Limit(limit).
		Offset(offset).
		All(ctx)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to query db: %w", err)
	}

	return users, totalItems, nil
}

func splitSearchTokens(searchText string) []string {
	normalized := strings.ToLower(searchText)
	return strings.Fields(normalized)
}
