package router

import (
	"net/http"

	"github.com/reedkihaddi/REST-API/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.Handler
}

var routes []Route

func (env *Env) initRoutes() {


	routes = []Route{
		Route {
			Name: "Hello", 
			Method: "GET", 
			Pattern: "/", 
			HandlerFunc: handlers.HelloHandler(env.db),
		},
	}

	for _, route := range routes {
		env.router.
			Handle(route.Pattern, route.HandlerFunc).
			Name(route.Name).Methods(route.Method)
	}

}
