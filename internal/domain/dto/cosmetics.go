package dto

import (
	"Leech-ru/internal/domain/types"
	"github.com/google/uuid"
)

type Cosmetics struct {
	ID                uuid.UUID      `json:"id"`
	Category          types.Category `json:"category"`
	Title             string         `json:"title"`
	Description       *string        `json:"description,omitempty"`
	ApplicationMethod *string        `json:"applicationMethod,omitempty"`
	Volume            *int           `json:"volume,omitempty"`
}

type CreateCosmeticsRequest struct {
	Category          types.Category
	Title             string
	Description       *string
	ApplicationMethod *string
	Volume            *int
}
type CreateCosmeticsResponse Cosmetics

type GetCosmeticsRequest struct {
	ID uuid.UUID `json:"id"`
}
type GetCosmeticsResponse Cosmetics

type GetAllCosmeticsRequest struct {
	Limit  *int
	Offset *int
}

// TODO []*Cosmetics
type GetAllCosmeticsResponse []Cosmetics

type UpdateCosmeticsRequest struct {
	Category          *types.Category
	Title             *string
	Description       *string
	ApplicationMethod *string
	Volume            *int
}
type UpdateCosmeticsResponse Cosmetics

type DeleteCosmeticsRequest struct {
	ID uuid.UUID `json:"id"`
}
