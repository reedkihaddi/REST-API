package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/reedkihaddi/REST-API/config"
	"github.com/reedkihaddi/REST-API/logging"
	"github.com/reedkihaddi/REST-API/router"
)

// @title Products API
// @version 1.0
// @description This is a sample REST-API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	
	srv := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: router,
	}
	
	// Create a channel to listen for shutdown
	signalChan := make(chan os.Signal, 1)
	// Notify channel for signals
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // sent when a program loses its controlling terminal
		syscall.SIGINT,  // sent when controlling terminal presses the interrupt character (Control-C)
		syscall.SIGQUIT, // sent when controlling terminal presses the quit character (Control-Backslash)
	)
	
	// Listen to connections
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logging.Log.Fatalf("listen: %s\n", err)
		}
	}()
	logging.Log.Info("Server started.")
	
	// Block
	<-signalChan
	logging.Log.Info("Server stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()
	
	if err := srv.Shutdown(ctx); err != nil {
		logging.Log.Fatalf("Server shutdown failed:%+v", err)
	}
	
	logging.Log.Info("Server exited properly")

}
