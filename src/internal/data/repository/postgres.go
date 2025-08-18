package repository

import (
	"context"
	"database/sql"
	"time"
	"url-shortener/src/internal/data/model"

	"log/slog"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresRepo struct {
	db  *sql.DB
	log *slog.Logger
}

func NewPostgresRepo(dsn string, log *slog.Logger) (*PostgresRepo, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresRepo{db: db, log: log}, nil
}

func (r *PostgresRepo) Save(url model.URLData) error {
	if url.CreatedAt.IsZero() {
		url.CreatedAt = time.Now()
	}

	_, err := r.db.ExecContext(
		context.Background(),
		`INSERT INTO urls (id, original_url, created_at) VALUES ($1, $2, $3)`,
		url.ID, url.URL, url.CreatedAt,
	)
	if err != nil {
		r.log.Error("failed to save URL", slog.String("id", url.ID), slog.String("url", url.URL), slog.Any("err", err))
	}
	return err
}

func (r *PostgresRepo) Find(id string) (model.URLData, bool) {
	var u model.URLData
	err := r.db.QueryRowContext(
		context.Background(),
		`SELECT id, original_url, created_at FROM urls WHERE id=$1`,
		id,
	).Scan(&u.ID, &u.URL, &u.CreatedAt)

	if err != nil {
		if err != sql.ErrNoRows {
			r.log.Error("failed to find URL", slog.String("id", id), slog.Any("err", err))
		}
		return model.URLData{}, false
	}
	return u, true
}
