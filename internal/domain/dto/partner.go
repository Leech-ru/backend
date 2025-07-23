package dto

import "github.com/google/uuid"

type Partner struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

type CreatePartnerRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreatePartnerResponse Partner

type GetByIdPartnerRequest struct {
	ID uuid.UUID `json:"id"`
}
type GetByIdPartnerResponse Partner
type GetAllPartnerRequest struct {
	Limit  *int
	Offset *int
}
type GetAllPartnerResponse []*Partner

type UpdatePartnerRequest struct {
	ID          uuid.UUID `json:"id"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
}
type UpdatePartnerResponse Partner
type DeletePartnerRequest struct {
	ID uuid.UUID `json:"id"`
}
