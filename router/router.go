package router

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"

	//pq is the PostgreSQL driver
	_ "github.com/lib/pq"
	database "github.com/reedkihaddi/REST-API/db"
)

//Env struct contains info about db and router for the router package only.
type Env struct {
	db     *database.DB
	router *mux.Router
}

var env Env

//NewRouter creates a new router.
func NewRouter(user, password, dbname string) *mux.Router {

	//Connect to the DB.
	conn := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname)
	var err error
	dbcon, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	// Pass the db connection to database package.
	env.db, err = database.New(dbcon)
	if err != nil {
		log.Fatal(err)
	}
	//Create a new router and initialize the routes.
	env.router = mux.NewRouter()
	env.initRoutes()

	return env.router
}
