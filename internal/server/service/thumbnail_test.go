package service

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/echelon/internal/server/repository/sqlite"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestCacheThumbnailService_GetCache(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := &sqlite.LiteRepo{DB: db}
	cacheService := NewCacheThumbnailService(repo)

	ctx := context.Background()

	tests := []struct {
		name          string
		videoID       string
		setupMock     func()
		expectedExist bool
		expectedError error
	}{
		{
			name:    "thumbnail exists",
			videoID: "sampleVideoID",
			setupMock: func() {
				mock.ExpectQuery(`SELECT thumbnail FROM thumbnails WHERE video_id = ?`).
					WithArgs("sampleVideoID").
					WillReturnRows(sqlmock.NewRows([]string{"thumbnail"}).AddRow([]byte("thumbnailData")))
			},
			expectedExist: true,
			expectedError: nil,
		},
		{
			name:    "thumbnail does not exist",
			videoID: "sampleVideoID",
			setupMock: func() {
				mock.ExpectQuery(`SELECT thumbnail FROM thumbnails WHERE video_id = ?`).
					WithArgs("sampleVideoID").
					WillReturnError(sql.ErrNoRows)
			},
			expectedExist: false,
			expectedError: nil,
		},
		{
			name:    "error fetching thumbnail",
			videoID: "sampleVideoID",
			setupMock: func() {
				mock.ExpectQuery(`SELECT thumbnail FROM thumbnails WHERE video_id = ?`).
					WithArgs("sampleVideoID").
					WillReturnError(errors.New("db error"))
			},
			expectedExist: false,
			expectedError: errors.New("repo getting thumbnail: sql query: db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			_, exist, err := cacheService.GetCache(ctx, tt.videoID)

			require.Equal(t, tt.expectedExist, exist)
			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestCacheThumbnailService_SaveThumbnail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := &sqlite.LiteRepo{DB: db}
	cacheService := NewCacheThumbnailService(repo)

	ctx := context.Background()

	tests := []struct {
		name          string
		videoID       string
		thumbnail     []byte
		setupMock     func()
		expectedError error
	}{
		{
			name:      "successful save",
			videoID:   "sampleVideoID",
			thumbnail: []byte("thumbnailData"),
			setupMock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT OR REPLACE INTO thumbnails`).
					WithArgs("sampleVideoID", []byte("thumbnailData")).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
		{
			name:      "error during save",
			videoID:   "sampleVideoID",
			thumbnail: []byte("thumbnailData"),
			setupMock: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`INSERT OR REPLACE INTO thumbnails`).
					WithArgs("sampleVideoID", []byte("thumbnailData")).
					WillReturnError(errors.New("db error"))
				mock.ExpectRollback()
			},
			expectedError: errors.New("store thumbnail: sql thumbnails store: db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			err = cacheService.SaveThumbnail(ctx, tt.videoID, tt.thumbnail)

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestFetchThumbnail(t *testing.T) {
	tests := []struct {
		name          string
		videoID       string
		httpResponse  *http.Response
		expectedError error
	}{
		{
			name:    "successful fetch",
			videoID: "sampleVideoID",
			httpResponse: &http.Response{StatusCode: http.StatusOK,
				Body: io.NopCloser(bytes.NewReader([]byte("thumbnailData")))},
			expectedError: nil,
		},
		{
			name:          "error during fetch",
			videoID:       "sampleVideoID",
			httpResponse:  &http.Response{StatusCode: http.StatusInternalServerError},
			expectedError: errors.New("get request: 500 Internal Server Error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			thumbnail, err := FetchThumbnail(tt.videoID)
			require.NoError(t, err)
			require.NotZero(t, thumbnail)
		})
	}
}

func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		name          string
		videoURL      string
		expectedID    string
		expectedError error
	}{
		{
			name:       "valid URL",
			videoURL:   "https://www.youtube.com/watch?v=abcd1234",
			expectedID: "abcd1234",
		},
		{
			name:          "invalid URL without video ID",
			videoURL:      "https://www.youtube.com/watch",
			expectedError: errors.New("ID not found in the URL"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			videoID, err := ExtractVideoID(tt.videoURL)
			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedID, videoID)
			}
		})
	}
}
