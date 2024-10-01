package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type LiteRepo struct {
	DB *sql.DB
}

func (l *LiteRepo) Open(path string) error {
	var err error

	l.DB, err = sql.Open("sqlite3", path)
	if err != nil {
		return fmt.Errorf("sql open: %w", err)
	}

	if err = uploadMigrations(l.DB); err != nil {
		return fmt.Errorf("upload migrations: %w", err)
	}

	return nil
}

func (l *LiteRepo) Close() error {
	if err := l.DB.Close(); err != nil {
		return fmt.Errorf("sql close: %w", err)
	}
	return nil
}

func (l *LiteRepo) GetThumbnail(ctx context.Context, videoID string) ([]byte, bool, error) {
	var thumbnail []byte
	query := `SELECT thumbnail FROM thumbnails WHERE video_id = ?`
	err := l.DB.QueryRowContext(ctx, query, videoID).Scan(&thumbnail)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("sql query: %w", err)
	}
	return thumbnail, true, nil
}

func (l *LiteRepo) StoreThumbnail(ctx context.Context, videoID string, thumbnail []byte) error {
	tx, err := l.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("sql tx begin: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	query := `INSERT OR REPLACE INTO thumbnails (video_id, thumbnail) VALUES (?, ?)`
	_, err = tx.ExecContext(ctx, query, videoID, thumbnail)
	if err != nil {
		return fmt.Errorf("sql thumbnails store: %w", err)
	}

	return nil
}

func uploadMigrations(db *sql.DB) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("create driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/server/repository/sqlite/migrations",
		"sqlite3", driver)
	if err != nil {
		return fmt.Errorf("migrate instance: %w", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("run migrations: %w", err)
	}
	return nil
}
