package inmemory

import (
	"github.com/nickklode/ozon-urlshortener/internal/storage"
	"github.com/nickklode/ozon-urlshortener/internal/utils/generator"
	"github.com/nickklode/ozon-urlshortener/internal/utils/validator"
)

type Store struct {
	m map[string]storage.Links
}

func New() *Store {
	return &Store{m: map[string]storage.Links{}}
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
	sur = storage.Links{OriginalUrl: orig, Token: generator.GenerateToken()}
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
	return "", err
}
