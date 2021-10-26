package storage

import "context"

type FileStorage interface {
	CreateSignedURL(ctx context.Context, fileID string) (string, error)
}
