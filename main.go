package main

import (
	"github.com/reedkihaddi/REST-API/logging"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/reedkihaddi/REST-API/router"
)

// @title Products API
// @version 1.0
// @description This is a sample REST-API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email saruabhbraryo@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

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
