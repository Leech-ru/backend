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
	ApplicationMethod *string        `json:"application_method,omitempty"`
	Volume            *int           `json:"volume,omitempty"`
}

type CreateCosmeticsRequest struct {
	Category          types.Category `json:"category" validate:"required,category"`
	Title             string         `json:"title" validate:"required,min=3,max=100"`
	Description       *string        `json:"description,omitempty" validate:"omitempty,min=3,max=3000"`
	ApplicationMethod *string        `json:"application_method,omitempty" validate:"omitempty,min=3,max=500"`
	Volume            *int           `json:"volume,omitempty" validate:"omitempty,min=1,max=10000"`
}
type CreateCosmeticsResponse Cosmetics

type GetCosmeticsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
type GetCosmeticsResponse Cosmetics

type GetAllByFilterCosmeticsRequest struct {
	Limit       *int            `json:"limit,omitempty" form:"limit" validate:"omitempty,min=1,max=100"`
	Offset      *int            `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0"`
	Category    *types.Category `json:"category,omitempty" form:"category" validate:"omitempty,category"`
	TitlePrefix *string         `json:"title_prefix,omitempty" form:"titlePrefix" validate:"omitempty,min=1,max=100"`
	Volume      *int            `json:"volume,omitempty" form:"volume" validate:"omitempty,min=1,max=10000"`
}
type GetAllByFilterCosmeticsResponse []*Cosmetics

type UpdateCosmeticsRequest struct {
	ID                uuid.UUID       `json:"id" validate:"required,uuid"`
	Category          *types.Category `json:"category,omitempty" validate:"omitempty,category"`
	Title             *string         `json:"title,omitempty" validate:"omitempty,min=3,max=100"`
	Description       *string         `json:"description,omitempty" validate:"omitempty,min=3,max=3000"`
	ApplicationMethod *string         `json:"application_method,omitempty" validate:"omitempty,min=3,max=500"`
	Volume            *int            `json:"volume,omitempty" validate:"omitempty,min=1,max=10000"`
}
type UpdateCosmeticsResponse Cosmetics

type DeleteCosmeticsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid"`
}
