package dto

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
type CreateCategoryRequest struct {
	Name string `json:"name"`
}
type CreateCategoryResponse Category

type GetAllCategoriesRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
type GetAllCategoriesResponse []*Category

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}
type UpdateCategoryResponse Category

type DeleteCategoryRequest struct {
	ID uuid.UUID `json:"id"`
}
