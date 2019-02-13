package s420

import (
	"context"

	s420con "github.com/bregydoc/s420/connection"
)

// SaveObjectInBucket implements a s420 server
func (s *S420Service) SaveObjectInBucket(c context.Context, params *s420con.NewObjectInBucket) (*s420con.SaveResponse, error) {
	r, err := s.store.SaveObjectInBucket(params.Bucket, params.FileName, params.Data, &ObjectOptions{
		ContentType: params.Options.ContentType,
	})

	if err != nil {
		return nil, err
	}

	return &s420con.SaveResponse{
		Ok:       r.OK,
		FilePath: r.FilePath,
	}, nil
}

// GetObjectFromBucket implements a s420 server
func (s *S420Service) GetObjectFromBucket(c context.Context, params *s420con.ObjectFromBucket) (*s420con.GetResponse, error) {
	data, contentType, err := s.store.GetObjectFromBucket(params.Bucket, params.Path)
	if err != nil {
		return nil, err
	}

	return &s420con.GetResponse{
		ContentType: string(contentType),
		Data:        data,
	}, nil
}

// SaveNewObject implements a s420 server
func (s *S420Service) SaveNewObject(c context.Context, params *s420con.NewObjectParams) (*s420con.SaveResponse, error) {
	r, err := s.store.SaveNewObject(params.Path, params.Data, &ObjectOptions{
		ContentType: params.Options.ContentType,
	})

	if err != nil {
		return nil, err
	}

	return &s420con.SaveResponse{
		Ok:       r.OK,
		FilePath: r.FilePath,
	}, nil
}

// GetObject implements a s420 server
func (s *S420Service) GetObject(c context.Context, params *s420con.ObjectPath) (*s420con.GetResponse, error) {
	data, contentType, err := s.store.GetObject(params.Path)
	if err != nil {
		return nil, err
	}

	return &s420con.GetResponse{
		ContentType: string(contentType),
		Data:        data,
	}, nil
}
