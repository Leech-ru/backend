package dto

import "github.com/google/uuid"

type News struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	ImageID  uuid.UUID `json:"image_id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Href     string    `json:"href"`
	IsHidden bool      `json:"is_hidden"`
}

type NewsListItem struct {
	ID             uuid.UUID `json:"id"`
	ImageID        uuid.UUID `json:"image_id"`
	Title          string    `json:"title"`
	ContentPreview string    `json:"content_preview"`
	Href           string    `json:"href"`
	IsHidden       bool      `json:"is_hidden"`
}

type CreateNewsRequest struct {
	ImageID  uuid.UUID `json:"image_id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Href     string    `json:"href"`
	IsHidden bool      `json:"is_hidden"`
}
type CreateNewsResponse News
type GetByIdNewsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}
type GetByIdNewsResponse News
type GetAllByFilterNewsRequest struct {
	Limit         *int
	Offset        *int
	PreviewLength *int
}

type GetAllByFilterNewsResponse []*NewsListItem

type GetAllByFilterForAdminsNewsRequest struct {
	Limit         *int
	Offset        *int
	PreviewLength *int
	IsHidden      *bool
}
type GetAllByFilterForAdminsNewsResponse []*NewsListItem
type UpdateNewsRequest struct {
	ID       uuid.UUID  `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	ImageID  *uuid.UUID `json:"image_id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title    *string    `json:"title"`
	Content  *string    `json:"content"`
	Href     *string    `json:"href"`
	IsHidden *bool      `json:"is_hidden"`
}
type UpdateNewsResponse News
type DeleteNewsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}
