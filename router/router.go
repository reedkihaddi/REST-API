package router

import (
	"github.com/reedkihaddi/REST-API/models"
	"github.com/reedkihaddi/REST-API/db"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)


func NewRouter(user, password, dbname string) *mux.Router {
	conn := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname)
	var err error
	db, err := sql.Open("postgres", conn)
	// postgres://%s:%s@localhost/%s?sslmode=disable
	if err != nil {
		log.Fatal(err)
	}
	stuffDB, err := database.New(db)
    if err != nil {
        log.Fatal(err)
    }
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello!")
		stuffDB.CreateProduct(&models.Product{ID:1,Name:"Wll",Price:11.3})
	})
	return r
}

