package user

import (
	"Leech-ru/internal/domain/common/errorz"
	"Leech-ru/pkg/ent"
	"context"
)

// Create creates a new user in the database
func (s *userRepo) Create(ctx context.Context, userEntity ent.User) (*ent.User, error) {
	created, err := s.client.User.
		Create().
		SetEmail(userEntity.Email).
		SetPassword(userEntity.Password).
		SetName(userEntity.Name).
		SetSurname(userEntity.Surname).
		SetRole(userEntity.Role).
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
