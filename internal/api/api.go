package api

import (
	"github.com/gorilla/mux"
	"github.com/nickklode/ozon-urlshortener/internal/storage"
)

type API struct {
	db storage.StorageInterface
	r  *mux.Router
}

func New(db storage.StorageInterface) *API {
	a := API{db: db, r: mux.NewRouter()}
	a.endpoints()
	return &a
}

func (api *API) Router() *mux.Router {
	return api.r
}
