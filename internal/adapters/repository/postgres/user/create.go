package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new user in the database
func (s *userRepo) Create(ctx context.Context, entity ent.User) (*ent.User, error) {
	created, err := s.client.User.
		Create().
		SetEmail(entity.Email).
		SetPassword(entity.Password).
		SetName(entity.Name).
		SetSurname(entity.Surname).
		SetRole(entity.Role).
		Save(ctx)

	switch {
	case ent.IsConstraintError(err):
		return nil, errorz.EmailAlreadyExist
	case err != nil:
		return nil, err
	}
	return &ent.User{
		ID:       created.ID,
		Email:    created.Email,
		Password: created.Password,
		Name:     created.Name,
		Surname:  created.Surname,
		Role:     created.Role,
	}, nil
}
