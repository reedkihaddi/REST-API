package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/reedkihaddi/REST-API/pkg/logging"

	"github.com/gorilla/mux"
	database "github.com/reedkihaddi/REST-API/pkg/db"
	"github.com/reedkihaddi/REST-API/pkg/models"
	"go.uber.org/zap"
)

// GetProduct godoc
// @Description Get a product by ID
// @Tags product
// @Produce  json
// @Param id path int true "Get product"
// @Success 200 {object} models.Product
// @Failure 404 {object} models.HTTPError{error=string}
// @Failure 500 {object} models.HTTPError{error=string}
// @Router /product/{id} [get]
func GetProduct(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			logging.Log.Error("Couldn't convert id to int.")
			respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
			return
		}
		p := &models.Product{}
		p.ID = id
		if err := db.GetProduct(p); err != nil {
			switch err {
			case sql.ErrNoRows:
				logging.Log.Infof("Product with id:%s requested but not found.", vars["id"])
				respondWithError(w, http.StatusNotFound, "Product not found")
			default:
				logging.Log.Infof("Product with id:%s requested but not found.", vars["id"])
				respondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		respondWithJSON(w, http.StatusOK, p)
	})
}

// CreateProduct godoc
// @Description Create a product
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Create a product"
// @Success 201 {object} models.Product
// @Failure 404 {object} models.HTTPError{error=string}
// @Failure 500 {object} models.HTTPError{error=string}
// @Router /product [post]
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
				logging.Log.Info("Error in creating product")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		defer r.Body.Close()
		if err := db.CreateProduct(p); err != nil {
			logging.Log.Infof("Couldn't insert a product into database. error: %s", err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, p)
	})
}

// UpdateProduct godoc
// @Description Update a product by ID
// @Tags product
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Update a product"
// @Param id path int true "Update product"
// @Success 200 {object} models.Product
// @Failure 404 {object} models.HTTPError{error=string}
// @Failure 500 {object} models.HTTPError{error=string}
// @Router /product/{id} [put]
func UpdateProduct(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			logging.Log.Info("Couldn't convert id to int.")
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
			logging.Log.Infof("Couldn't update a product with id:%s. error: %s", vars["id"], err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, p)
	})
}

// DeleteProduct godoc
// @Description Delete a product by ID
// @Tags product
// @Param id path int true "Delete product"
// @Success 200 {object} models.HTTPOK
// @Failure 404 {object} models.HTTPError{error=string}
// @Failure 500 {object} models.HTTPError{error=string}
// @Router /product/{id} [delete]
func DeleteProduct(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			logging.Log.Info("Couldn't convert id to int.")
			respondWithError(w, http.StatusBadRequest, "Invalid product ID")
			return
		}
		p := &models.Product{}
		p.ID = id
		if err := db.DeleteProduct(p); err != nil {
			logging.Log.Infof("Couldn't delete a product with id:%s. error: %s", vars["id"], err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	})
}

// ListProducts godoc
// @Description Lists all the products
// @Tags product
// @Security ApiKeyAuth
// @Success 200 "OK"
// @Failure 404 {object} models.HTTPError{error=string}
// @Failure 500 {object} models.HTTPError{error=string}
// @Router /products [get]
func ListProducts(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products, err := db.ListProducts()
		if err != nil {
			logging.Log.Infof("Couldn't list products. error: %s", err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, products)
	})
}

// WithMetrics returns the time taken for the request.
func WithMetrics(l *zap.SugaredLogger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		began := time.Now()
		next.ServeHTTP(w, r)
		l.Infof("%s %s took %s", r.Method, r.URL, time.Since(began))
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

// Hello is just hello?
func Hello(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello there."))
	})
}
