package repository

import (
	"cloud.google.com/go/storage"
	"github.com/nustvondev/vjetmemorize/account/model"
)

type gcImageRepository struct {
	Storage    *storage.Client
	BucketName string
}

// NewImageRepository is a factory for initializing User Repositories
func NewImageRepository(gcClient *storage.Client, bucketName string) model.ImageRepository {
	return &gcImageRepository{
		Storage:    gcClient,
		BucketName: bucketName,
	}
}
