package storage

type StorageInterface interface {
	CreateToken(orig string) (string, error)
	GetByToken(token string) (string, error)
}
