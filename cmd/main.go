package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nickklode/ozon-urlshortener/internal/api"
	"github.com/nickklode/ozon-urlshortener/internal/storage"
	"github.com/nickklode/ozon-urlshortener/internal/storage/inmemory"
	"github.com/nickklode/ozon-urlshortener/internal/storage/postgres"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error init config. %s", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env. %s", err)
	}

	var si storage.StorageInterface

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("enter db type")
	}
	store := arguments[1]

	switch store {
	case "postgres":
		var err error
		si, err = postgres.New(postgres.Config{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: os.Getenv("DB_PASSWORD"),
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
			os.Exit(1)
		}
	case "inmemory":
		si = inmemory.New()
	default:
		log.Fatal("invalid db")

	}

	api := api.New(si)

	log.Fatal(http.ListenAndServe(":8080", api.Router()))

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
