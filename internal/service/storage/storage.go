package storage

import (
	"go_service/config"
	"go_service/internal/service/storage/local"
	"go_service/internal/service/storage/oss"
	"go_service/internal/service/storage/request"
)

const (
	TypeLocal = iota // 本地存储
	TypeOss          // 阿里云oss
)

type IStorage interface {
	PutObject(filename string, value []byte) (string, error)
	GetObject(filename string) ([]byte, error)
	GetStoragePath(fileName string) string
	GetSignUrl(storageKey, fileName string) (string, error)
	GetUploadRequest(filename, callbackUrl string) (*request.FileUploadRequest, error)
}

var storageType int

func GetStorage() (IStorage, error) {
	storageType = config.GetConfig().File.StorageType

	switch storageType {
	case TypeLocal:
		storageLocal, err := local.GetStorageLocal()
		if err != nil {
			return nil, err
		}

		return storageLocal, nil

	case TypeOss:
		storageOSS, err := oss.GetStorageOSS()
		if err != nil {
			return nil, err
		}

		return storageOSS, nil

	default:
		storageOSS, err := oss.GetStorageOSS()
		if err != nil {
			return nil, err
		}

		return storageOSS, nil
	}
}
