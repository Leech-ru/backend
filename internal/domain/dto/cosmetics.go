package dto

import (
	"Leech-ru/internal/domain/types"
	"github.com/google/uuid"
)

// Cosmetics represents the cosmetic product structure.
// @Description Contains product information such as category, title, description, and volume.
type Cosmetics struct {
	ID                uuid.UUID      `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Category          types.Category `json:"category" example:"3"`
	Title             string         `json:"title" example:"Hair Shampoo"`
	Description       *string        `json:"description,omitempty" example:"Suitable for daily use."`
	ApplicationMethod *string        `json:"application_method,omitempty" example:"Apply to wet hair, lather, rinse."`
	Volume            *int           `json:"volume,omitempty" example:"250"`
	Links             *Links         `json:"links"`
}

// Links contains links information.
type Links struct {
	Ozon        *string `json:"ozon" validate:"omitempty,max=500,ozonlink" example:"https://www.ozon.ru/product/gel-girudo-dr-nikonov-dlya-tela-100-ml-1907286044/?at=OgtEDAg59hR8AYgGimrnA9YIqYo9mocJYEzPjHR666Gm"`
	Wildberries *string `json:"wildberries" validate:"omitempty,max=500,wildberrieslink" example:"https://www.wildberries.ru/catalog/344283033/detail.aspx"`
}

// CreateCosmeticsRequest represents a request to create a new cosmetic product.
type CreateCosmeticsRequest struct {
	Category          types.Category `json:"category" validate:"required,category" example:"3"`
	Title             string         `json:"title" validate:"required,min=3,max=100" example:"Hair Shampoo"`
	Description       *string        `json:"description,omitempty" validate:"omitempty,min=3,max=3000" example:"Suitable for daily use."`
	ApplicationMethod *string        `json:"application_method,omitempty" validate:"omitempty,min=3,max=500" example:"Apply to wet hair, lather, rinse."`
	Volume            *int           `json:"volume,omitempty" validate:"omitempty,min=1,max=10000" example:"250"`
	Links             *Links         `json:"links,omitempty" validate:"omitempty"`
}

// CreateCosmeticsResponse represents the response after creating a cosmetic product.
type CreateCosmeticsResponse Cosmetics

// GetCosmeticsRequest represents a request to get a cosmetic product by ID.
type GetCosmeticsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

// GetCosmeticsResponse returns a cosmetic product by ID.
type GetCosmeticsResponse Cosmetics

// GetAllByFilterCosmeticsRequest is used to filter and list cosmetic products.
// TODO add volume from - volume to
type GetAllByFilterCosmeticsRequest struct {
	Limit       *int            `json:"limit,omitempty" form:"limit" validate:"omitempty,min=1,max=100" example:"10"`
	Offset      *int            `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0" example:"0"`
	Category    *types.Category `json:"category,omitempty" form:"category" validate:"omitempty,category" example:"2"`
	TitlePrefix *string         `json:"title_prefix,omitempty" form:"titlePrefix" validate:"omitempty,min=1,max=100" example:"Hair"`
	Volume      *int            `json:"volume,omitempty" form:"volume" validate:"omitempty,min=1,max=10000" example:"250"`
}

// GetAllByFilterCosmeticsResponse is the list of cosmetics returned by filters.
type GetAllByFilterCosmeticsResponse []*Cosmetics

// UpdateCosmeticsRequest represents an update request for a cosmetic product.
type UpdateCosmeticsRequest struct {
	ID                uuid.UUID       `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
	Category          *types.Category `json:"category,omitempty" validate:"omitempty,category" example:"4"`
	Title             *string         `json:"title,omitempty" validate:"omitempty,min=3,max=100" example:"New Hair Shampoo"`
	Description       *string         `json:"description,omitempty" validate:"omitempty,min=3,max=3000" example:"Updated product description."`
	ApplicationMethod *string         `json:"application_method,omitempty" validate:"omitempty,min=3,max=500" example:"Apply evenly and rinse well."`
	Volume            *int            `json:"volume,omitempty" validate:"omitempty,min=1,max=10000" example:"500"`
	Links             *Links          `json:"links,omitempty" validate:"omitempty"`
}

// UpdateCosmeticsResponse returns the updated cosmetic product.
type UpdateCosmeticsResponse Cosmetics

// DeleteCosmeticsRequest represents a request to delete a cosmetic product by ID.
type DeleteCosmeticsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}
