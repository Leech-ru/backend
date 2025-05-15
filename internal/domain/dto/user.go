package dto

import "github.com/google/uuid"

// TODO валидацию сделать
type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Role     int    `json:"role"`
}

type RegisterUserResponse struct {
	Token   string    `json:"token"`
	ID      uuid.UUID `json:"id"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Role    int       `json:"role"`
}

type GetUserRequest struct {
	ID uuid.UUID `json:"id"`
}

type GetUserResponse struct {
	ID      uuid.UUID `json:"id"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Role    int       `json:"role"`
}

type GetAllUsersRequest struct {
	Limit  *int `json:"limit,omitempty" form:"limit"`
	Offset *int `json:"offset,omitempty" form:"offset"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token   string    `json:"token"`
	ID      uuid.UUID `json:"id"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Role    int       `json:"role"`
}

type UpdateUserRequest struct {
	ID      uuid.UUID `json:"id"`
	Name    *string   `json:"name,omitempty"`
	Surname *string   `json:"surname,omitempty"`
	Role    *int      `json:"role,omitempty"`
}

type UpdateUserResponse struct {
	ID      uuid.UUID `json:"id"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Role    int       `json:"role"`
}

type ChangePasswordRequest struct {
	ID          uuid.UUID `json:"id"`
	OldPassword string    `json:"old_password"`
	NewPassword string    `json:"new_password"`
}

type ChangePasswordResponse struct {
	Token string `json:"token"`
}

type DeleteUserRequest struct {
	ID uuid.UUID `json:"id"`
}
