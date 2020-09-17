package router

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	database "github.com/reedkihaddi/REST-API/db"
)

type Env struct {
	db     *database.DB
	router *mux.Router
}

var env Env

func NewRouter(user, password, dbname string) *mux.Router {
	conn := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname)
	var err error
	dbcon, err := sql.Open("postgres", conn)
	// postgres://%s:%s@localhost/%s?sslmode=disable
	if err != nil {
		log.Fatal(err)
	}
	env.db, err = database.New(dbcon)
	if err != nil {
		log.Fatal(err)
	}
	env.router = mux.NewRouter()
	env.initRoutes()
	return env.router
}

