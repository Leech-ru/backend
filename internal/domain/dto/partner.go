package dto

import (
	"github.com/google/uuid"
)

// Partner represents a partner entity.
type Partner struct {
	ID          uuid.UUID      `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Name        string         `json:"name" example:"Backend's department"`
	Description *string        `json:"description,omitempty" example:"Leading supplier of industrial equipment"`
	Links       []PartnersLink `json:"links,omitempty"`
}

// PartnersLink represents a hyperlink related to a partner.
type PartnersLink struct {
	Label string `json:"label" validate:"required,min=1,max=150" example:"Official Website"`
	Href  string `json:"href" validate:"required,min=1,max=2000" example:"https://example.com"`
}

// CreatePartnerRequest represents a request to create a partner.
type CreatePartnerRequest struct {
	Name        string         `json:"name" validate:"required,min=1,max=100" example:"Backend's department"`
	Description *string        `json:"description,omitempty" validate:"omitempty,max=500" example:"Leading supplier of industrial equipment"`
	Links       []PartnersLink `json:"links,omitempty" validate:"omitempty,dive"`
}

// CreatePartnerResponse represents the response after creating a partner.
type CreatePartnerResponse Partner

// GetByIdPartnerRequest represents a request to get a partner by ID.
type GetByIdPartnerRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true"`
}

// GetByIdPartnerResponse represents the response to a get-by-id request.
type GetByIdPartnerResponse Partner

// GetAllPartnerRequest represents a request to get all partners with optional pagination.
type GetAllPartnerRequest struct {
	Limit  *int `json:"limit,omitempty" form:"limit" validate:"omitempty,min=0,max=100" example:"10"`
	Offset *int `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0" example:"0"`
}

// GetAllPartnerResponse represents the response to a get-all request.
type GetAllPartnerResponse struct {
	Items      []*Partner      `json:"items"`
	Pagination PaginationInfo `json:"pagination"`
}

// UpdatePartnerRequest represents a request to update a partner.
type UpdatePartnerRequest struct {
	ID          uuid.UUID       `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true"`
	Name        *string         `json:"name,omitempty" validate:"omitempty,min=1,max=100" example:"New Name"`
	Description *string         `json:"description,omitempty" validate:"omitempty,max=500" example:"Updated description of the partner"`
	Links       []*PartnersLink `json:"links,omitempty" validate:"omitempty,dive"`
}

// UpdatePartnerResponse represents the response after updating a partner.
type UpdatePartnerResponse Partner

// DeletePartnerRequest represents a request to delete a partner.
type DeletePartnerRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000" swaggerignore:"true"`
}
