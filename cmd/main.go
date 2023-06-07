package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nickklode/ozon-urlshortener/internal/api"
	"github.com/nickklode/ozon-urlshortener/internal/storage"
	"github.com/nickklode/ozon-urlshortener/internal/storage/inmemory"
	"github.com/nickklode/ozon-urlshortener/internal/storage/postgres"
)

func main() {

	var si storage.StorageInterface

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("enter db type")
	}
	store := arguments[1]
	coo := "postgres://postgres:ZAQzaqzaq97@localhost:5433/linkstest?sslmode=disable"
	switch store {
	case "postgres":
		var err error
		// "user=postgres password=ZAQzaqzaq97 host=db port=5432 dbname=postgres sslmode=disable"
		si, err = postgres.New(coo)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
			os.Exit(1)
		}
	case "inmemory":
		si = inmemory.New()
	}

	api := api.New(si)

	log.Fatal(http.ListenAndServe(":8080", api.Router()))

}
