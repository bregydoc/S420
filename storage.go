package s420

// ObjectOptions ...
type ObjectOptions struct {
	ContentType string
}

// SaveResponse ...
type SaveResponse struct {
	OK       bool
	FilePath string
}

// ContentType ...
type ContentType string

// StorageSystem represent the store struct to dialog with minio or any store type
type StorageSystem interface {
	SaveObjectInBucket(bucket string, fileName string, data []byte, options *ObjectOptions) (*SaveResponse, error)
	GetObjectFromBucket(bucket string, path string) ([]byte, ContentType, error) // data, content-type, error
	SaveNewObject(path string, data []byte, options *ObjectOptions) (*SaveResponse, error)
	GetObject(path string) ([]byte, ContentType, error) // data, content-type, error
	ListObjectsOfBucket(bucket string, withThumbnails bool) ([]*File, error)
}
