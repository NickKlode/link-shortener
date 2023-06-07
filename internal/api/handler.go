package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nickklode/ozon-urlshortener/internal/storage"
)

func (api *API) endpoints() {
	api.r.HandleFunc("/", api.createToken).Methods("POST")
	api.r.HandleFunc("/{id}", api.getOriginal).Methods("GET")
}

func (api *API) createToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newLink storage.Links

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
