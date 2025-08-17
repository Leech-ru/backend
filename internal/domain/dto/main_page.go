package dto

import "github.com/google/uuid"

// MainPage represents the main page content structure
type MainPage struct {
	ID       uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	ImageID  uuid.UUID `json:"image_id" example:"123e4567-e89b-12d3-a456-426614174100"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Href     string    `json:"href"`
	IsHidden bool      `json:"is_hidden"`
	Fluid    bool      `json:"fluid"`
}

// CreateMainPageRequest represents a request to create a new main page content
type CreateMainPageRequest struct {
	ImageID  uuid.UUID `json:"image_id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title    string    `json:"title" validate:"required,min=2,max=100"`
	Content  string    `json:"content" validate:"required,min=2,max=1000"`
	Href     string    `json:"href" validate:"required,min=2,max=1000"`
	IsHidden bool      `json:"is_hidden"`
	Fluid    bool      `json:"fluid"`
}

// CreateMainPageResponse represents the response after creation a main page content
type CreateMainPageResponse MainPage

// GetByIdMainPageRequest represents a request to get a main page content by id
type GetByIdMainPageRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true"`
}

// GetByIdMainPageResponse represents a main page content
type GetByIdMainPageResponse MainPage

// GetAllMainPageRequest represents a request to get all main page contents with pagination
type GetAllMainPageRequest struct {
	Limit  int `json:"limit" example:"10"`
	Offset int `json:"offset" example:"0"`
}

// GetAllMainPageResponse represents a main page contents
type GetAllMainPageResponse []*MainPage

// UpdateMainPageRequest represents a request to update main page content data
type UpdateMainPageRequest struct {
	ID       uuid.UUID  `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true"`
	ImageID  *uuid.UUID `json:"image_id" validate:"omitempty,uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title    *string    `json:"title" validate:"omitempty,min=2,max=100"`
	Content  *string    `json:"content" validate:"omitempty,min=2,max=1000"`
	Href     *string    `json:"href" validate:"omitempty,min=2,max=1000"`
	IsHidden *bool      `json:"is_hidden"`
	Fluid    *bool      `json:"fluid"`
}

// UpdateMainPageResponse represents an updated main page content
type UpdateMainPageResponse MainPage

// DeleteMainPageRequest represents a request to delete a main page content by ID
type DeleteMainPageRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true"`
}
