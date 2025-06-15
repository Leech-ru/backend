package dto

import (
	"Leech-ru/internal/domain/types"
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID  `json:"id"`
	Email   string     `json:"email"`
	Name    string     `json:"name"`
	Surname string     `json:"surname"`
	Role    types.Role `json:"role"`
}

type RegisterUserRequest struct {
	Email    string      `json:"email" validate:"required,email,min=6,max=254"`
	Password string      `json:"password" validate:"required,min=8,max=100"`
	Name     string      `json:"name" validate:"required,min=2,max=100"`
	Surname  string      `json:"surname" validate:"required,min=2,max=100"`
	Role     *types.Role `json:"role,omitempty" validate:"omitempty,role"`
}

type RegisterUserResponse struct {
	RefreshToken string `json:"-"`
	User
}

type GetUserRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}

type GetUserResponse User

type GetAllUsersRequest struct {
	Limit  *int `json:"limit,omitempty" form:"limit" validate:"omitempty,min=1,max=100"`
	Offset *int `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0"`
}

// TODO сделать []*User
type GetAllUsersResponse []User

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email,min=6,max=254"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type LoginUserResponse struct {
	RefreshToken string `json:"-"`
	User
}

type LogoutRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}

type UpdateUserRequest struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid"`
	Name    *string   `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Surname *string   `json:"surname,omitempty" validate:"omitempty,min=2,max=100"`
}

type UpdateUserResponse User

type ChangePasswordRequest struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid"`
	OldPassword string    `json:"old_password" validate:"required,min=8,max=100"`
	NewPassword string    `json:"new_password" validate:"required,min=8,max=100"`
}

type ChangePasswordResponse struct {
	RefreshToken string `json:"-"`
}

type DeleteUserRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
