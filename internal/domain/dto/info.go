package dto

import "Leech-ru/internal/domain/types"

// Info represents public corporation information.
type Info struct {
	Heading     string          `json:"heading" validate:"required,min=3,max=100" example:"Welcome to our company"`
	Description string          `json:"description" validate:"required,min=10,max=500" example:"We are a global leader in innovation and technology."`
	Fluid       *bool           `json:"fluid,omitempty" validate:"omitempty" example:"true"`
	Schedule    []ScheduleEntry `json:"schedule" validate:"dive,min=1,max=7"`
	Links       []InfoLinks     `json:"links" validate:"dive,min=0,max=10"`
}

// ScheduleEntry defines a working day and its hours.
type ScheduleEntry struct {
	Weekday types.Weekday `json:"weekday" validate:"required,weekday" example:"1"` // Monday
	Hours   Hours         `json:"hours" validate:"required"`
}

// Hours represents opening and closing time for a single day.
type Hours struct {
	Open  string `json:"open" validate:"required,len=5,datetime=15:04" example:"09:00"`
	Close string `json:"close" validate:"required,len=5,datetime=15:04" example:"18:00"`
}

// InfoLinks defines a label and a corresponding URL.
type InfoLinks struct {
	Label string `json:"label" validate:"required,min=2,max=50" example:"Instagram"`
	Href  string `json:"href" validate:"required,url,max=500" example:"https://instagram.com/company"`
}

// GetInfoResponse is the response structure for fetching corporation info.
type GetInfoResponse Info

// UpdateInfoRequest defines fields that can be updated for corporation info.
type UpdateInfoRequest struct {
	Heading     *string          `json:"heading,omitempty" validate:"omitempty,min=3,max=100" example:"Updated Heading"`
	Description *string          `json:"description,omitempty" validate:"omitempty,min=10,max=500" example:"Updated long description about the company."`
	Fluid       *bool            `json:"fluid,omitempty" validate:"omitempty" example:"false"`
	Schedule    *[]ScheduleEntry `json:"schedule,omitempty" validate:"omitempty,dive,min=1,max=7"`
	Links       *[]InfoLinks     `json:"links,omitempty" validate:"omitempty,dive,min=0,max=10"`
}

// UpdateInfoResponse is the response structure after updating corporation info.
type UpdateInfoResponse Info
