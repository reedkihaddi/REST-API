package handlers

import (
	"encoding/json"
	"net/http"

	database "github.com/reedkihaddi/REST-API/db"
)

func writeJSONResponse(w http.ResponseWriter, code int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}

// Check is for the health of API
func Check(db *database.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbUp := db.Check()
		if dbUp {
			data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
			writeJSONResponse(w, http.StatusOK, data)
		} else {
			data, _ := json.Marshal(healthCheckResponse{Status: "Database unaccessible"})
			writeJSONResponse(w, http.StatusServiceUnavailable, data)
		}
	})
}
