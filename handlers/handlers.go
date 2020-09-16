package handlers

import (
	"fmt"
	"net/http"

	database "github.com/reedkihaddi/REST-API/db"
)

func HelloHandler(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Printf("%+v\n",emp)
		//p := &models.Product{}
		//p.ID = 1
		//fmt.Printf("%+v\n",p)
		//db.DeleteProduct(&models.Product{ID: 1, Name: "Wlll", Price: 11.3})
		//x, _ := db.GetProducts(5, 1)
		//db.CreateProduct(&models.Product{ID: 1, Name: "Wll", Price: 11.3})
		// Write it back to the client.
		fmt.Printf("Hello!")
	})
}
