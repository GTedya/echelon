package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestLiteRepo_GetThumbnail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := &LiteRepo{DB: db}

	ctx := context.Background()

	tests := []struct {
		name          string
		videoID       string
		setupMock     func()
		expectedError error
		expectedFound bool
		expectedThumb []byte
	}{
		{
			name:    "success",
			videoID: "sampleVideoID",
			setupMock: func() {
				rows := sqlmock.NewRows([]string{"thumbnail"}).AddRow([]byte("sampleThumbnail"))
				mock.ExpectQuery("SELECT thumbnail FROM thumbnails WHERE video_id = ?").
					WithArgs("sampleVideoID").
					WillReturnRows(rows)
			},
			expectedError: nil,
			expectedFound: true,
			expectedThumb: []byte("sampleThumbnail"),
		},
		{
			name:    "not found",
			videoID: "sampleVideoID",
			setupMock: func() {
				mock.ExpectQuery("SELECT thumbnail FROM thumbnails WHERE video_id = ?").
					WithArgs("sampleVideoID").
					WillReturnError(sql.ErrNoRows)
			},
			expectedError: nil,
			expectedFound: false,
			expectedThumb: nil,
		},
		{
			name:    "sql error",
			videoID: "sampleVideoID",
			setupMock: func() {
				mock.ExpectQuery("SELECT thumbnail FROM thumbnails WHERE video_id = ?").
					WithArgs("sampleVideoID").
					WillReturnError(errors.New("some sql error"))
			},
			expectedError: errors.New("sql query: some sql error"),
			expectedFound: false,
			expectedThumb: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			thumbnail, found, err := repo.GetThumbnail(ctx, tt.videoID)

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.expectedFound, found)
			require.Equal(t, tt.expectedThumb, thumbnail)
		})
	}
}

func TestLiteRepo_StoreThumbnail(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := &LiteRepo{DB: db}

	ctx := context.Background()

	tests := []struct {
		name          string
		videoID       string
		thumbnail     []byte
		setupMock     func()
		expectedError error
	}{
		{
			name:      "success",
			videoID:   "sampleVideoID",
			thumbnail: []byte("sampleThumbnail"),
			setupMock: func() {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT OR REPLACE INTO thumbnails").
					WithArgs("sampleVideoID", []byte("sampleThumbnail")).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: nil,
		},
		{
			name:      "sql error during exec",
			videoID:   "sampleVideoID",
			thumbnail: []byte("sampleThumbnail"),
			setupMock: func() {
				mock.ExpectBegin()
				mock.ExpectExec("INSERT OR REPLACE INTO thumbnails").
					WithArgs("sampleVideoID", []byte("sampleThumbnail")).
					WillReturnError(errors.New("some sql error"))
				mock.ExpectRollback()
			},
			expectedError: errors.New("sql thumbnails store: some sql error"),
		},
		{
			name:      "sql error during begin",
			videoID:   "sampleVideoID",
			thumbnail: []byte("sampleThumbnail"),
			setupMock: func() {
				mock.ExpectBegin().WillReturnError(errors.New("begin tx error"))
			},
			expectedError: errors.New("sql tx begin: begin tx error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()

			err = repo.StoreThumbnail(ctx, tt.videoID, tt.thumbnail)

			if tt.expectedError != nil {
				require.EqualError(t, err, tt.expectedError.Error())
			} else {
				require.NoError(t, err)
			}

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
