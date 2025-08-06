package dto

import (
	"github.com/google/uuid"
)

type Category struct {
	ID   uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Name string    `json:"name" validate:"required,min=1,max=150" example:"Shampoo"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=1,max=150" example:"Shampoo"`
}

type CreateCategoryResponse Category

type GetByIdCategoryRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}

type GetByIdCategoryResponse Category

type GetAllCategoriesRequest struct {
	Limit  *int `json:"limit,omitempty" form:"limit" validate:"omitempty,min=1,max=100" example:"10"`
	Offset *int `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0" example:"0"`
}

type GetAllCategoriesResponse []*Category

type UpdateCategoryRequest struct {
	ID   uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Name *string   `json:"name" validate:"omitempty,min=1,max=150" example:"Hair conditioner"`
}

type UpdateCategoryResponse Category

type DeleteCategoryRequest struct {
	ID uuid.UUID `json:"id" validate:"required" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}
