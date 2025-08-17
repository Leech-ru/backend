package dto

import (
	"io"
	"time"
)

type FilePackage struct {
	Content      io.Reader `json:"-" swaggerignore:"true"`
	ContentType  string    `json:"content_type" validate:"required,filetype=image/png;image/jpeg;image/jpg;image/gif;image/webp;image/bmp;image/tiff;image/heic"`
	Size         int64     `json:"size" validate:"required,gt=0,maxfilesize=5242880"`
	Filename     string    `json:"filename" validate:"required,max=250"`
	LastModified time.Time `json:"last_modified" validate:"lte"`
}
