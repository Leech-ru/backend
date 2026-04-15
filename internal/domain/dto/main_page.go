package dto

import "github.com/google/uuid"

// MainPage represents the main page content structure
type MainPage struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	ImageID  uuid.UUID `json:"image_id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174100"`
	Title    string    `json:"title" validate:"required,min=2,max=100" example:"Main banner title"`
	Content  string    `json:"content" validate:"required,min=2,max=1000" example:"Main banner description text goes here..."`
	Href     string    `json:"href" validate:"required,min=1,max=2000" example:"https://example.com/page"`
	IsHidden bool      `json:"is_hidden" validate:"required" example:"false"`
	Fluid    bool      `json:"fluid" validate:"required" example:"true"`
}

// CreateMainPageRequest represents a request to create a new main page content
type CreateMainPageRequest struct {
	ImageID  uuid.UUID `json:"image_id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Title    string    `json:"title" validate:"required,min=2,max=100" example:"Main banner title"`
	Content  string    `json:"content" validate:"required,min=2,max=1000" example:"Main banner description text goes here..."`
	Href     string    `json:"href" validate:"required,min=1,max=2000" example:"https://example.com/page"`
	IsHidden bool      `json:"is_hidden" validate:"required" example:"false"`
	Fluid    bool      `json:"fluid" validate:"required" example:"true"`
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
	Limit  *int `json:"limit,omitempty" form:"limit" validate:"omitempty,min=1,max=100" example:"10"`
	Offset *int `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0" example:"0"`
}

// GetAllMainPageResponse represents a main page contents
type GetAllMainPageResponse struct {
	Items      []*MainPage     `json:"items"`
	Pagination PaginationInfo `json:"pagination"`
}

// UpdateMainPageRequest represents a request to update main page content data
type UpdateMainPageRequest struct {
	ID       uuid.UUID  `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true"`
	ImageID  *uuid.UUID `json:"image_id,omitempty" validate:"omitempty,uuid" example:"123e4567-e89b-12d3-a456-426614174100"`
	Title    *string    `json:"title,omitempty" validate:"omitempty,min=2,max=100" example:"Updated main banner title"`
	Content  *string    `json:"content,omitempty" validate:"omitempty,min=2,max=1000" example:"Updated main banner description text..."`
	Href     *string    `json:"href,omitempty" validate:"omitempty,min=1,max=2000" example:"https://example.com/updated"`
	IsHidden *bool      `json:"is_hidden,omitempty" validate:"omitempty" example:"true"`
	Fluid    *bool      `json:"fluid,omitempty" validate:"omitempty" example:"false"`
}

// UpdateMainPageResponse represents an updated main page content
type UpdateMainPageResponse MainPage

// DeleteMainPageRequest represents a request to delete a main page content by ID
type DeleteMainPageRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true"`
}
