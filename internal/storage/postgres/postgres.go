package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nickklode/ozon-urlshortener/internal/utils/generator"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type DB struct {
	pool *pgxpool.Pool
}

func New(cfg Config) (*DB, error) {

	c, err := pgxpool.Connect(context.Background(), fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	s := DB{
		pool: c,
	}

	return &s, nil
}

func (db *DB) CreateToken(orig string) (string, error) {

	token := generator.GenerateToken()

	query := "INSERT INTO links(original_url, token) VALUES ($1, $2) ON CONFLICT (original_url) DO NOTHING RETURNING token"
	err := db.pool.QueryRow(context.Background(), query, orig, token).Scan(&token)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (db *DB) GetByToken(token string) (string, error) {

	var original string
	query := "SELECT original_url FROM links WHERE token = $1"
	err := db.pool.QueryRow(context.Background(), query, token).Scan(&original)
	if err != nil {
		return "", err
	}

	return original, nil
}
