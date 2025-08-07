package dto

import (
	"io"
	"time"
)

type FilePackage struct {
	Content       io.Reader // Тело файла
	ContentType   string    // MIME-тип (например, "image/png")
	ContentLength int64     // Размер файла в байтах
	Filename      string    // Имя файла (например, "cat.png")
	LastModified  time.Time // Время последнего изменения
}
