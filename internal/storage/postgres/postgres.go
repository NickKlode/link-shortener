package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nickklode/ozon-urlshortener/internal/service/generator"
	"github.com/nickklode/ozon-urlshortener/internal/service/validator"
	"github.com/nickklode/ozon-urlshortener/internal/storage"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(p string) (*DB, error) {

	c, err := pgxpool.Connect(context.Background(), p)
	if err != nil {
		return nil, err
	}

	s := DB{
		pool: c,
	}

	return &s, nil
}

func (db *DB) CreateToken(orig string) (string, error) {
	err := validator.ValidateURL(orig)
	if err != nil {
		return "", err
	}
	token := generator.GenerateToken()

	query := "INSERT INTO links(original_url, token) VALUES ($1, $2) ON CONFLICT (original_url) DO NOTHING RETURNING token"
	err = db.pool.QueryRow(context.Background(), query, orig, token).Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (db *DB) GetByToken(token string) (string, error) {
	err := validator.ValidateToken(token)
	if err != nil {
		return "", err
	}
	var original storage.Links
	query := "SELECT original_url FROM links WHERE token = $1"
	err = db.pool.QueryRow(context.Background(), query, token).Scan(&original)
	if err != nil {
		return "", err
	}

	return original.OriginalUrl, nil
}
