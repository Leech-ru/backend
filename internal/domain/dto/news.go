package dto

import "github.com/google/uuid"

type News struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	ImageID  uuid.UUID `json:"image_id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title    string    `json:"title" validate:"required,max=255" example:"Breaking News"`
	Content  string    `json:"content" validate:"required,max=5000" example:"This is the full news content..."`
	IsHidden bool      `json:"is_hidden" validate:"required" example:"false"`
}

type NewsListItem struct {
	ID             uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	ImageID        uuid.UUID `json:"image_id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title          string    `json:"title" validate:"required,max=255" example:"Breaking News"`
	ContentPreview string    `json:"content_preview" validate:"required,max=255" example:"This is a short preview..."`
	IsHidden       bool      `json:"is_hidden" validate:"required" example:"false"`
}

type CreateNewsRequest struct {
	ImageID  uuid.UUID `json:"image_id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title    string    `json:"title" validate:"required,max=255" example:"Breaking News"`
	Content  string    `json:"content" validate:"required,max=5000" example:"This is the full news content..."`
	IsHidden *bool     `json:"is_hidden" validate:"required" example:"false"`
}

type CreateNewsResponse News

type GetByIdNewsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}
type GetByIdNewsResponse News

type GetAllByFilterNewsRequest struct {
	Limit         *int `json:"limit,omitempty" form:"limit" validate:"omitempty,min=0" example:"10"`
	Offset        *int `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0" example:"0"`
	PreviewLength *int `json:"preview_length,omitempty" form:"preview_length" validate:"omitempty,min=10,max=500" example:"100"`
}
type GetAllByFilterNewsResponse struct {
	Items      []*NewsListItem `json:"items"`
	Pagination PaginationInfo `json:"pagination"`
}

type GetAllByFilterForAdminsNewsRequest struct {
	Limit         *int  `json:"limit,omitempty" form:"limit" validate:"omitempty,min=0" example:"10"`
	Offset        *int  `json:"offset,omitempty" form:"offset" validate:"omitempty,min=0" example:"0"`
	PreviewLength *int  `json:"preview_length,omitempty" form:"preview_length" validate:"omitempty,min=10,max=500" example:"100"`
	IsHidden      *bool `json:"is_hidden,omitempty" form:"is_hidden" validate:"omitempty" example:"true"`
}
type GetAllByFilterForAdminsNewsResponse struct {
	Items      []*NewsListItem `json:"items"`
	Pagination PaginationInfo `json:"pagination"`
}

type UpdateNewsRequest struct {
	ID       uuid.UUID  `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	ImageID  *uuid.UUID `json:"image_id,omitempty" validate:"omitempty,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
	Title    *string    `json:"title,omitempty" validate:"omitempty,max=255" example:"Updated News Title"`
	Content  *string    `json:"content,omitempty" validate:"omitempty,max=5000" example:"Updated full news content..."`
	IsHidden *bool      `json:"is_hidden,omitempty" validate:"omitempty" example:"true"`
}
type UpdateNewsResponse News

type DeleteNewsRequest struct {
	ID uuid.UUID `json:"id" validate:"required,uuid" example:"3fa85f64-5717-4562-b3fc-2c963f66afa6"`
}
