package image

import (
	"github.com/minio/minio-go/v7"
)

type imageRepo struct {
	client     *minio.Client
	bucketName string
}

func NewImageRepo(client *minio.Client, bucketName string) *imageRepo {
	return &imageRepo{
		client:     client,
		bucketName: bucketName,
	}
}
