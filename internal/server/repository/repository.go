package repository

import "context"

type Repository interface {
	ThumbNailRepository
	Open(path string) error
	Close() error
}

type ThumbNailRepository interface {
	GetThumbnail(ctx context.Context, videoID string) ([]byte, bool, error)
	StoreThumbnail(ctx context.Context, videoID string, thumbnail []byte) error
}
