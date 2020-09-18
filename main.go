package main

import (
	"net/http"
	"os"

	"github.com/reedkihaddi/REST-API/config"
	"github.com/reedkihaddi/REST-API/logging"
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

	logging.InitLogger()
	config.InitConfig()
	router := router.NewRouter()
	logging.Log.Info("Starting up the server.")
	http.ListenAndServe(os.Getenv("PORT"), router)

}
