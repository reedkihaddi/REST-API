package router

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/reedkihaddi/REST-API/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	database "github.com/reedkihaddi/REST-API/db"
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
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Printf("Hello!")
	// 	stuffDB.CreateProduct(&models.Product{ID: 1, Name: "Wll", Price: 11.3})
	// })
	r.Handle("/", handlers.HelloHandler(stuffDB))
	return r
}

// func hello(db *database.DB) {
// 	fmt.Printf("Hello!")
// 	db.CreateProduct(&models.Product{ID: 1, Name: "Wll", Price: 11.3})

// }

// func helloHandler(db *database.DB) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		db.CreateProduct(&models.Product{ID: 1, Name: "Wll", Price: 11.3})
// 		// Write it back to the client.
// 		fmt.Printf("Hello!")
// 	})
// }
