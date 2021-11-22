package storage

import (
	"context"
	"time"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
)

type Storage struct {
	ProjectID      string
	StorageID      string
	expirationTime int
	allowedType    string
	BucketHandle   *storage.BucketHandle
}

func NewStorage(ctx context.Context, projectID, storageID string, expirationTime int, allowedType string) (Storage, error) {
	if projectID == "" {
		return Storage{}, errors.New("missing projectID dependency")
	}

	if storageID == "" || expirationTime == 0 || allowedType == "" {
		return Storage{}, errors.New("missing environment variables dependency")
	}

	c, err := storage.NewClient(ctx)
	if err != nil {
		return Storage{}, errors.Wrap(err, "couldn't create storage client")
	}
	bkt := c.Bucket(storageID)
	return Storage{
		ProjectID:      projectID,
		StorageID:      storageID,
		expirationTime: expirationTime,
		allowedType:    allowedType,
		BucketHandle:   bkt,
	}, nil
}

func MustNewStorage(ctx context.Context, projectID, storageID string, expirationTime int, allowedType string) Storage {
	s, err := NewStorage(ctx, projectID, storageID, expirationTime, allowedType)
	if err != nil {
		panic(errors.Wrap(err, "couldn't create storage instance"))
	}
	return s
}

func (s Storage) CreateSignedURL(ctx context.Context, fileID string) (string, error) {
	et := time.Now().Add(time.Duration(s.expirationTime))
	u, err := s.BucketHandle.SignedURL(fileID, &storage.SignedURLOptions{
		Expires:     et,
		ContentType: s.allowedType,
		Method:      "PUT",
	})
	if err != nil {
		return "", errors.Wrap(err, "couldn't generate signedUrl")
	}
	return u, nil
}
