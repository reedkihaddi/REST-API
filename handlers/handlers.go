package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	database "github.com/reedkihaddi/REST-API/db"
	"github.com/reedkihaddi/REST-API/models"
)

// func HelloHandler(db *database.DB) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// fmt.Printf("%+v\n",emp)
// 		//p := &models.Product{}
// 		//p.ID = 1
// 		//fmt.Printf("%+v\n",p)
// 		//db.DeleteProduct(&models.Product{ID: 1, Name: "Wlll", Price: 11.3})
// 		//x, _ := db.GetProducts(5, 1)
// 		db.CreateProduct(&models.Product{ID: 1, Name: "Wll", Price: 25.3})
// 		// Write it back to the client.
// 		fmt.Printf("Hello!")
// 	})
// }

//GetProduct is the handler to get product given the product ID.
func GetProduct(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		p := &models.Product{}
		p.ID = id
		if err := db.GetProduct(p); err != nil {
			switch err {
			case sql.ErrNoRows:
				respondWithError(w, http.StatusNotFound, "Product not found")
			default:
				respondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		respondWithJSON(w, http.StatusOK, p)
	})
}

//CreateProduct creates a product and inserts into the database.
func CreateProduct(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var p *models.Product

		err := dec.Decode(&p)
		if err != nil {
			var syntaxError *json.SyntaxError
			var unmarshalTypeError *json.UnmarshalTypeError

			switch {
			// Catch any syntax errors.
			case errors.As(err, &syntaxError):
				msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
				http.Error(w, msg, http.StatusBadRequest)

			// Catch any type errors.
			case errors.As(err, &unmarshalTypeError):
				msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
				http.Error(w, msg, http.StatusBadRequest)

			// An io.EOF error is returned by Decode() if the request body is empty.
			case errors.Is(err, io.EOF):
				msg := "Request body must not be empty"
				http.Error(w, msg, http.StatusBadRequest)

			// A 500 Internal Server Error response.
			default:
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		defer r.Body.Close()
		if err := db.CreateProduct(p); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, p)
	})
}

//UpdateProduct updates the product in the database.
func UpdateProduct(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}

		var p *models.Product
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&p); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()
		p.ID = id
		if err := db.UpdateProduct(p); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, p)
	})
}

//DeleteProduct deletes the product from the database.
func DeleteProduct(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}
		p := &models.Product{}
		p.ID = id
		if err := db.DeleteProduct(p); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	})
}

//ListProducts lists all the products from the database.
func ListProducts(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// count, _ := strconv.Atoi(r.FormValue("count"))
		// start, _ := strconv.Atoi(r.FormValue("start"))
	
		// if start < 0 {
		// 	start = 0
		// }
		// if count < 0 {
		// 	count = 0
		// }
		products, err := db.ListProducts()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	
		respondWithJSON(w, http.StatusOK, products)
	})
}

// Respond with error in case something goes wrong.
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Send back JSON.
func respondWithJSON(w http.ResponseWriter, code int, p interface{}) {
	response, _ := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Hello(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there."))
	})
}
