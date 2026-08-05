package fastdfs_storage

import (
	"context"

	s3_storage "github.com/storage-lock/go-s3-storage"
	"github.com/storage-lock/go-storage"
)

// FastdfsStorage fastdfs对象存储实现 (S3兼容)
type FastdfsStorage struct {
	*s3_storage.S3Storage
}

func NewFastdfsStorage(ctx context.Context, options *s3_storage.S3StorageOptions) (*FastdfsStorage, error) {
	s3Storage, err := s3_storage.NewS3Storage(ctx, options)
	if err != nil {
		return nil, err
	}
	return &FastdfsStorage{S3Storage: s3Storage}, nil
}

const StorageName = "fastdfs-storage"

func (x *FastdfsStorage) GetName() string {
	return StorageName
}

func (x *FastdfsStorage) Capabilities() []storage.StorageCapability {
	return x.S3Storage.Capabilities()
}
