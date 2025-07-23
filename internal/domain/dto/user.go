package dto

import (
	"Leech-ru/internal/domain/types"
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID  `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Email   string     `json:"email" example:"user@example.com"`
	Name    string     `json:"name" example:"Ivan"`
	Surname string     `json:"surname" example:"Ivanov"`
	Role    types.Role `json:"role" example:"0"`
}

type RegisterUserRequest struct {
	Email    string      `json:"email" validate:"required,email,min=6,max=254" example:"user@example.com"`
	Password string      `json:"password" validate:"required,min=8,max=100" example:"SecurePass123!" format:"password"`
	Name     string      `json:"name" validate:"required,min=2,max=100" example:"Ivan"`
	Surname  string      `json:"surname" validate:"required,min=2,max=100" example:"Ivanov"`
	Role     *types.Role `json:"role,omitempty" validate:"omitempty,role" swaggerignore:"true"`
}

type RegisterUserResponse struct {
	RefreshToken string `json:"-"`
	User
}

type GetUserRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000" swaggerignore:"true"`
}

type GetUserResponse User

type GetAllUsersRequest struct {
	Limit  *int `json:"limit,omitempty" form:"limit" validate:"omitempty,min=1,max=100" example:"10"`
	Offset *int `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0" example:"0"`
}

type GetAllByFilterUsersRequest struct {
	Limit         *int        `json:"limit,omitempty" form:"limit" validate:"omitempty,min=1,max=100" example:"10"`
	Offset        *int        `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0" example:"0"`
	Role          *types.Role `json:"role,omitempty" form:"role" validate:"omitempty,role" example:"0"`
	NamePrefix    *string     `json:"name_prefix,omitempty" form:"name_prefix" validate:"omitempty,min=1,max=100" example:"Iv"`
	SurnamePrefix *string     `json:"surname_prefix,omitempty" form:"surname_prefix" validate:"omitempty,min=1,max=100" example:"Ivan"`
	EmailPrefix   *string     `json:"email_prefix,omitempty" form:"email_prefix" validate:"omitempty,email" example:"user@"`
}

type GetAllByFilterUsersResponse []*User

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email,min=6,max=254" example:"user@example.com"`
	Password string `json:"password" validate:"required,min=8,max=100" example:"SecurePass123!" format:"password"`
}

type LoginUserResponse struct {
	RefreshToken string `json:"-"`
	User
}

type LogoutRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" swaggerignore:"true"`
}

type UpdateUserRequest struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" swaggerignore:"true"`
	Name    *string   `json:"name,omitempty" validate:"omitempty,min=2,max=100" example:"Ivan"`
	Surname *string   `json:"surname,omitempty" validate:"omitempty,min=2,max=100" example:"Petrov"`
}

type UpdateUserResponse User

type ChangePasswordRequest struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid" swaggerignore:"true"`
	OldPassword string    `json:"old_password" validate:"required,min=8,max=100" example:"OldPass123!" format:"password"`
	NewPassword string    `json:"new_password" validate:"required,min=8,max=100" example:"NewPass456!" format:"password"`
}

type ChangePasswordResponse struct {
	RefreshToken string `json:"-" swaggerignore:"true"`
}

type DeleteUserRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" swaggerignore:"true"`
}
