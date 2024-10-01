package service

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	pb "github.com/echelon/api/grpc"
	z "github.com/echelon/pkg/logger"
)

const requestTimeout = 3 * time.Second

var filePermission = 0600

func DownloadThumbnailsSync(client pb.ThumbnailServiceClient, videoURLs []string) {
	for _, videoURL := range videoURLs {
		if err := downloadThumbnail(client, videoURL); err != nil {
			z.Logger.Errorf("err download thumbnail for %s: %v", videoURL, err)
		}
	}
}

func DownloadThumbnailsAsync(client pb.ThumbnailServiceClient, videoURLs []string) {
	var wg sync.WaitGroup
	wg.Add(len(videoURLs))

	for _, videoURL := range videoURLs {
		go func(url string) {
			defer wg.Done()
			if err := downloadThumbnail(client, url); err != nil {
				z.Logger.Errorf("err download thumbnail for %s: %v", videoURL, err)
			}
		}(videoURL)
	}

	wg.Wait()
}

func downloadThumbnail(client pb.ThumbnailServiceClient, videoURL string) error {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req := &pb.ThumbnailRequest{
		VideoUrl: videoURL,
	}

	resp, err := client.GetThumbnail(ctx, req)
	if err != nil {
		return fmt.Errorf("gRPC request failed: %w", err)
	}

	videoID := extractVideoID(videoURL)

	// сохраняю в память, для наглядности)
	fileName := fmt.Sprintf("%s_thumbnail.jpg", videoID)
	if err = os.WriteFile(fileName, resp.ThumbnailData, os.FileMode(filePermission)); err != nil {
		return fmt.Errorf("save thumbnail: %w", err)
	}

	z.Logger.Debugf("Thumbnail for %s saved as %s", videoURL, fileName)
	return nil
}

func extractVideoID(videoURL string) string {
	parts := strings.Split(videoURL, "v=")
	if len(parts) > 1 {
		return strings.Split(parts[1], "&")[0]
	}
	return videoURL
}
