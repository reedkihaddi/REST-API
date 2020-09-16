package main

import (
	"net/http"

	"github.com/reedkihaddi/REST-API/router"
)

func main() {
	router := router.NewRouter("saurabh","therock01","testdb")
	http.ListenAndServe(":8080", router)
}
