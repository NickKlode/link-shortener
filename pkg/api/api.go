package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nickklode/ozon-urlshortener/pkg/storage"
	"github.com/nickklode/ozon-urlshortener/pkg/storage/postgres"
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

func (api *API) endpoints() {
	api.r.HandleFunc("/", api.createToken).Methods("POST")
	api.r.HandleFunc("/{id}", api.getOriginal).Methods("GET")
}

func (api *API) createToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newLink postgres.Links

	json.NewDecoder(r.Body).Decode(&newLink)
	st, err := api.db.CreateToken(newLink.OriginalUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(st)

}

func (api *API) getOriginal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := mux.Vars(r)["id"]

	st, err := api.db.GetByToken(s)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(st)
}
