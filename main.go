package main

import (
	"github.com/reedkihaddi/REST-API/logging"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/reedkihaddi/REST-API/router"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logging.InitLogger()
	router := router.NewRouter(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	logging.Log.Info("Starting up the server.")
	http.ListenAndServe(os.Getenv("PORT"), router)

}
