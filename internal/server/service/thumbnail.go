package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/echelon/internal/server/repository"
)

type CacheThumbnailService struct {
	repo repository.ThumbNailRepository
}

func NewCacheThumbnailService(repo repository.ThumbNailRepository) CacheThumbnailService {
	return CacheThumbnailService{repo: repo}
}

func (s *CacheThumbnailService) GetCache(ctx context.Context, videoURL string) ([]byte, bool, error) {
	thumbnail, exist, err := s.repo.GetThumbnail(ctx, videoURL)
	if err != nil {
		return nil, false, fmt.Errorf("repo getting thumbnail: %w", err)
	}
	return thumbnail, exist, nil
}

func (s *CacheThumbnailService) SaveThumbnail(ctx context.Context, videoID string, thumbnail []byte) error {
	if err := s.repo.StoreThumbnail(ctx, videoID, thumbnail); err != nil {
		return fmt.Errorf("store thumbnail: %w", err)
	}
	return nil
}

func FetchThumbnail(videoURL string) (thumbnail []byte, err error) {
	res, err := http.Get(fmt.Sprintf("https://img.youtube.com/vi/%s/hqdefault.jpg", videoURL))
	defer func() {
		err = res.Body.Close()
	}()
	if err != nil {
		return nil, fmt.Errorf("get request: %w", err)
	}

	thumbnail, err = io.ReadAll(res.Body)
	if err != nil && errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("reader : %w", err)
	}

	return thumbnail, nil
}

func ExtractVideoID(videoURL string) (string, error) {
	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return "", fmt.Errorf("parse URL: %w", err)
	}

	queryParams := parsedURL.Query()

	videoID := queryParams.Get("v")
	if videoID == "" {
		return "", fmt.Errorf("ID not found in the URL")
	}

	return videoID, nil
}
