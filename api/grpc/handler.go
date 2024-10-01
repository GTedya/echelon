package thumbnailapi

import (
	"context"
	"fmt"

	"github.com/echelon/internal/server/service"
)

type ThumbnailService struct {
	srv service.CacheThumbnailService
}

func NewThumbnailService(srv service.CacheThumbnailService) *ThumbnailService {
	return &ThumbnailService{srv: srv}
}

func (s *ThumbnailService) GetThumbnail(ctx context.Context, req *ThumbnailRequest) (*ThumbnailResponse, error) {
	videoID, err := service.ExtractVideoID(req.VideoUrl)
	if err != nil {
		return nil, fmt.Errorf("extract video ID: %w", err)
	}

	thumbnail, found, err := s.srv.GetCache(ctx, videoID)
	if err != nil {
		return nil, fmt.Errorf("check cache: %w", err)
	}
	if found {
		return &ThumbnailResponse{ThumbnailData: thumbnail}, nil
	}

	thumbnail, err = service.FetchThumbnail(videoID)
	if err != nil {
		return nil, fmt.Errorf("fetch thumbnail: %w", err)
	}

	err = s.srv.SaveThumbnail(ctx, videoID, thumbnail)
	if err != nil {
		return nil, fmt.Errorf("save thumbnail to cache: %w", err)
	}

	return &ThumbnailResponse{ThumbnailData: thumbnail}, nil
}

func (s *ThumbnailService) mustEmbedUnimplementedThumbnailServiceServer() {
}
