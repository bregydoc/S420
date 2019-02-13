package backends

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	s420 "github.com/bregydoc/S420"
	minio "github.com/minio/minio-go"
)

// MinioStorage ...
type MinioStorage struct {
	client *minio.Client
}

// NewMinioStore return a minio storage
func NewMinioStore(c *s420.Config) (*MinioStorage, error) {
	client, err := minio.New(c.Endpoint, c.AccessKeyID, c.SecretAccessKey, c.UseSSL)
	if err != nil {
		return nil, err
	}

	return &MinioStorage{
		client: client,
	}, nil
}

// SaveObjectInBucket implement a s420 storage
func (m *MinioStorage) SaveObjectInBucket(bucket string, fileName string, data []byte, options *s420.ObjectOptions) (*s420.SaveResponse, error) {
	buf := bytes.NewBuffer(data)

	if strings.HasPrefix(fileName, "/") {
		fileName = string(fileName[1:])
	}
	_, err := m.client.PutObject(bucket,
		fileName,
		buf,
		int64(buf.Len()),
		minio.PutObjectOptions{
			ContentType: options.ContentType,
		})

	if err != nil {
		return nil, err
	}

	return &s420.SaveResponse{
		OK:       true,
		FilePath: bucket + "/" + fileName,
	}, nil

}

// GetObjectFromBucket implement a s420 storage
func (m *MinioStorage) GetObjectFromBucket(bucket string, path string) ([]byte, s420.ContentType, error) {
	object, err := m.client.GetObject(bucket, path, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", err
	}

	data, err := ioutil.ReadAll(object)
	if err != nil {
		return nil, "", err
	}

	info, err := object.Stat()
	if err != nil {
		return nil, "", err
	}

	return data, s420.ContentType(info.ContentType), nil
}

// SaveNewObject implement a s420 storage
func (m *MinioStorage) SaveNewObject(path string, data []byte, options *s420.ObjectOptions) (*s420.SaveResponse, error) {
	if strings.HasSuffix(path, "/") {
		return nil, errors.New("invalid path, it not can have '/' at the last")
	}

	chunks := strings.Split(path, "/")
	if len(chunks) < 1 {
		return nil, fmt.Errorf("invalid path: %s", chunks)
	}

	bucket := chunks[0]
	filePath := strings.Join(chunks[1:], "/")

	return m.SaveObjectInBucket(bucket, filePath, data, options)
}

// GetObject implement a s420 storage
func (m *MinioStorage) GetObject(path string) ([]byte, s420.ContentType, error) {
	if strings.HasSuffix(path, "/") {
		return nil, "", errors.New("invalid path, it not can have '/' at the last")
	}

	chunks := strings.Split(path, "/")
	if len(chunks) < 1 {
		return nil, "", fmt.Errorf("invalid path: %s", chunks)
	}

	bucket := chunks[0]
	filePath := strings.Join(chunks[1:], "/")

	return m.GetObjectFromBucket(bucket, filePath)
}
