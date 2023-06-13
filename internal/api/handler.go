package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nickklode/ozon-urlshortener/internal/storage"
	"github.com/nickklode/ozon-urlshortener/internal/utils/validator"
)

func (api *API) endpoints() {
	api.r.HandleFunc("/", api.createToken).Methods("POST")
	api.r.HandleFunc("/{id}", api.getOriginal).Methods("GET")
}

func (api *API) createToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newLink storage.Links

	err := json.NewDecoder(r.Body).Decode(&newLink)
	if err != nil {
		log.Printf("unnable to decode the reques body. %v", err)
	}
	err = validator.ValidateURL(newLink.OriginalUrl)
	if err != nil {
		log.Printf("wrong url. %s", err)
		return
	}
	st, err := api.db.CreateToken(newLink.OriginalUrl)
	if err != nil {
		log.Printf("unnable to create token. %v", err)
	}

	json.NewEncoder(w).Encode(map[string]string{"token": st})

}

func (api *API) getOriginal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := mux.Vars(r)["id"]
	err := validator.ValidateToken(s)
	if err != nil {
		log.Printf("wrong token. %s", err)
		return
	}

	st, err := api.db.GetByToken(s)

	if err != nil {
		log.Printf("unnable to get by token. %v", err)
	}
	json.NewEncoder(w).Encode(map[string]string{"original_url": st})
}
