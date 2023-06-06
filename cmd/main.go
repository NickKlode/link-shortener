package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nickklode/ozon-urlshortener/pkg/api"
	"github.com/nickklode/ozon-urlshortener/pkg/storage"
	"github.com/nickklode/ozon-urlshortener/pkg/storage/inmemory"
	"github.com/nickklode/ozon-urlshortener/pkg/storage/postgres"
)

func main() {
	var store string
	var si storage.StorageInterface

	switch store {
	case "postgres":
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("error loading .env file")
		}
		si, err = postgres.New(os.Getenv("POSTGRES_CONN"))
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
