package router

import (
	"net/http"

	"github.com/reedkihaddi/REST-API/handlers"
)

// Route type description.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.Handler
}

//List of routes.
var routes []Route

// Contains all the routes.
func (env *Env) initRoutes() {

	routes = []Route{

		Route{
			Name:        "GetProduct",
			Method:      "GET",
			Pattern:     "/product/{id:[0-9]+}",
			HandlerFunc: handlers.GetProduct(env.db),
		},
		Route{
			Name:        "GetProdusct",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: handlers.Hello(env.db),
		},
		Route{
			Name:        "CreateProduct",
			Method:      "POST",
			Pattern:     "/product",
			HandlerFunc: handlers.CreateProduct(env.db),
		},
		Route{
			Name:        "UpdateProduct",
			Method:      "PUT",
			Pattern:     "/product/{id:[0-9]+}",
			HandlerFunc: handlers.UpdateProduct(env.db),
		},
		Route{
			Name:        "DeleteProduct",
			Method:      "DELETE",
			Pattern:     "/product",
			HandlerFunc: handlers.DeleteProduct(env.db),
		},
		Route{
			Name:        "ListProduct",
			Method:      "GET",
			Pattern:     "/products/",
			HandlerFunc: handlers.ListProducts(env.db),
		},
	}

	// Register the routes.
	for _, route := range routes {
		env.router.
			Handle(route.Pattern, route.HandlerFunc).
			Name(route.Name).Methods(route.Method)
	}

}
