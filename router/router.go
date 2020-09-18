package router

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/reedkihaddi/REST-API/logging"

	"github.com/gorilla/mux"

	//pq is the PostgreSQL driver
	_ "github.com/lib/pq"
	database "github.com/reedkihaddi/REST-API/db"
)

//Env struct contains info about db and router for the router package only.
type Env struct {
	DB     *database.DB
	Router *mux.Router
}

var env Env

//NewRouter creates a new router.
func NewRouter() *mux.Router {

	//Connect to the DB.
	logging.Log.Info("Connecting to the database.")
	conn := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	var err error
	dbcon, err := sql.Open("postgres", conn)
	if err != nil {
		logging.Log.Fatalf("db opening err: %s", err)
	}

	// Pass the db connection to database package.
	env.DB, err = database.New(dbcon)
	if err != nil {
		logging.Log.Error("Error in passing db conn.")
	}
	//Create a new router and initialize the routes.
	logging.Log.Info("Creating the router.")
	env.Router = mux.NewRouter()
	logging.Log.Info("Initalizing routes.")
	env.initRoutes()

	return env.Router
}
