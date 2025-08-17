package dto

import "github.com/google/uuid"

type News struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	ImageID uuid.UUID `json:"image_id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Href    string    `json:"href"`
}

type NewsListItem struct {
	ID             uuid.UUID `json:"id"`
	ImageID        uuid.UUID `json:"image_id"`
	Title          string    `json:"title"`
	ContentPreview string    `json:"content_preview"`
	Href           string    `json:"href"`
}

type CreateNewsRequest struct {
	ImageID uuid.UUID `json:"image_id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Href    string    `json:"href"`
}
type CreateNewsResponse News
type GetByIdNewsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}
type GetByIdNewsResponse News
type GetAllNewsRequest struct {
	Limit      *int
	Offset     *int
	PreContent *int
}
type GetAllNewsResponse []*NewsListItem
type UpdateNewsRequest struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	ImageID uuid.UUID `json:"image_id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Href    string    `json:"href"`
}
type UpdateNewsResponse News
type DeleteNewsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}
