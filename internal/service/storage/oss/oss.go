package oss

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"go-service/config"
	"go-service/internal/service/storage/request"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type StorageOss struct {
	ossInstance *oss.Client
	ossBucket   string
	ossEndpoint string
	saveOssPath string
}

var ossObject *StorageOss

func GetStorageOSS() (*StorageOss, error) {
	if ossObject == nil {
		ossObject = &StorageOss{
			ossInstance: GetOss(),
			ossBucket:   config.GetConfig().Oss.Bucket,
			ossEndpoint: config.GetConfig().Oss.Endpoint,
			saveOssPath: config.GetConfig().Oss.PathPrefix,
		}
	}

	return ossObject, nil
}

func (so *StorageOss) PutObject(filename string, value []byte) (string, error) {
	if filename == "" {
		return "", errors.New("oss put filename is empty")
	}

	bucket, err := so.ossInstance.Bucket(so.ossBucket)
	if err != nil {
		return "", err
	}

	resultFileName := "https://" + so.ossBucket + "." + so.ossEndpoint + "/" + so.saveOssPath + filename
	fullFileName := so.saveOssPath + filename
	err = bucket.PutObject(fullFileName, bytes.NewReader(value))
	if err != nil {
		return "", err
	}

	return resultFileName, nil
}

func (so *StorageOss) GetStoragePath(fileName string) string {
	return "https://" + so.ossBucket + "." + so.ossEndpoint + "/" + so.saveOssPath + fileName
}

func (so *StorageOss) GetObject(filename string) ([]byte, error) {
	if filename == "" {
		return nil, errors.New("oss put filename is empty")
	}

	bucket, err := so.ossInstance.Bucket(so.ossBucket)
	if err != nil {
		return nil, err
	}

	reader, err := bucket.GetObject(so.saveOssPath + filename)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(reader)
}

func (so *StorageOss) GetSignUrl(storageKey, fileName string) (string, error) {
	bucket, err := so.ossInstance.Bucket(so.ossBucket)
	if err != nil {
		return "", err
	}

	optionArr := []oss.Option{
		oss.ResponseContentDisposition(fmt.Sprintf("attachment; filename=%v", fileName)),
		oss.ResponseContentType("application/octet-stream"),
	}

	signUrl, err := bucket.SignURL(so.saveOssPath+storageKey, oss.HTTPGet, 60*59, optionArr...)
	urlParseResult, _ := url.Parse(signUrl)
	cfg := config.GetConfig()
	originalUrl := fmt.Sprintf("%s://%s", urlParseResult.Scheme, urlParseResult.Host)
	if cfg.File.StorageUrl != "" {
		signUrl = strings.Replace(signUrl, originalUrl, cfg.File.StorageUrl, 1)
	}
	return signUrl, err
}

func (so *StorageOss) GetUploadRequest(filename, callbackUrl string) (*request.FileUploadRequest, error) {

	uploadRequest := &request.FileUploadRequest{}
	uploadRequest.URL = "https://" + so.ossBucket + "." + so.ossEndpoint

	bodyParams, err := getPolicyToken(callbackUrl)
	if err != nil {
		return nil, nil
	}

	bodyParams.SuccessActionStatus = 200
	bodyParams.Key = so.saveOssPath + filename

	// struct to map
	bodyParamsMap := make(map[string]interface{})
	bodyParamsByte, err := json.Marshal(bodyParams)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bodyParamsByte, &bodyParamsMap); err != nil {
		return nil, err
	}

	uploadRequest.BodyParams = bodyParamsMap
	return uploadRequest, nil
}
