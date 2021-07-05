package local

import (
	"errors"
	"io/ioutil"
	"os"

	"go_service/internal/service/storage/request"
)

type StorageLocal struct {
}

var localObject *StorageLocal

func GetStorageLocal() (*StorageLocal, error) {
	if localObject == nil {
		err := localInit()
		if err != nil {
			return nil, err
		}
	}

	return localObject, nil
}

func localInit() error {
	return nil
}

func (sl *StorageLocal) Init() error {
	return nil
}

func (sl *StorageLocal) PutObject(filename string, value []byte) (string, error) {
	if len(value) == 0 {
		return "", errors.New("param[value] is empty")
	}

	if filename == "" {
		return "", errors.New("param[filename] is empty")
	}

	err := sl.Init()
	if err != nil {
		return "", err
	}

	writeFile, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer writeFile.Close()

	_, err = writeFile.Write(value)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (sl *StorageLocal) GetObject(filename string) ([]byte, error) {
	if filename == "" {
		return nil, errors.New("param[filename] is empty")
	}

	err := sl.Init()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadFile(filename)
}

func (sl *StorageLocal) GetStoragePath(fileName string) string {
	return fileName
}

func (sl *StorageLocal) GetSignUrl(storageKey, fileName string) (string, error) {
	return fileName, nil
}

func (sl *StorageLocal) GetUploadRequest(fileName, callbackUrl string) (*request.FileUploadRequest, error) {
	return nil, nil
}
