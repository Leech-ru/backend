package dto

import "github.com/google/uuid"

type Image struct {
	ID   uuid.UUID    `json:"id"`
	File *FilePackage `json:"file"`
}

type CreateImageRequest struct {
	File *FilePackage `json:"file"`
}
type CreateImageResponse Image

type GetByIdImageRequest struct {
	ID uuid.UUID `json:"id"`
}
type GetByIdImageResponse Image

type DeleteImageRequest struct {
	ID uuid.UUID `json:"id"`
}
