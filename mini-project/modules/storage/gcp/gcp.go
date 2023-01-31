package gcp

import (
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type gcpStorage struct {
	client *storage.Client
}

func NewStorage(client *storage.Client) *gcpStorage {
	return &gcpStorage{client}
}

func (g *gcpStorage) UploadImage(name, prefix string, file multipart.FileHeader)
