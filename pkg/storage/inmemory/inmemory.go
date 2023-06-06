package inmemory

import (
	"github.com/jackc/pgx/v4"
	"github.com/nickklode/ozon-urlshortener/pkg/generator"
	"github.com/nickklode/ozon-urlshortener/pkg/validator"
)

type Links struct {
	OriginalUrl string `json:"original_url"`
	Token       string `json:"token"`
}

type Store struct {
	m map[string]Links
}

func New() *Store {
	return &Store{m: map[string]Links{}}
}

func (s *Store) CreateToken(orig string) (string, error) {
	err := validator.ValidateURL(orig)
	if err != nil {
		return "", err
	}
	sur, ok := s.m[orig]
	if ok {
		return sur.Token, nil
	}
	sur = Links{OriginalUrl: orig, Token: generator.GenerateToken()}
	s.m[orig] = sur
	return sur.Token, nil
}

func (s *Store) GetByToken(token string) (string, error) {
	err := validator.ValidateToken(token)
	if err != nil {
		return "", err
	}
	for _, sur := range s.m {
		if sur.Token == token {
			return sur.OriginalUrl, nil
		}
	}
	return "", pgx.ErrNoRows
}
