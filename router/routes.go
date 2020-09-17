package router

import (
	"net/http"

	"github.com/reedkihaddi/REST-API/logging"

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

//var logger = log.New(os.Stdout, "", log.Ldate | log.Ltime)

// Contains all the routes.
func (env *Env) initRoutes() {

	routes = []Route{

		Route{
			Name:        "GetProduct",
			Method:      "GET",
			Pattern:     "/product/{id:[0-9]+}",
			HandlerFunc: handlers.WithMetrics(logging.Log, handlers.GetProduct(env.DB)),
		},
		Route{
			Name:        "GetProdusct",
			Method:      "GET",
			Pattern:     "/",
			HandlerFunc: handlers.WithMetrics(logging.Log, handlers.Hello(env.DB)),
		},
		Route{
			Name:        "CreateProduct",
			Method:      "POST",
			Pattern:     "/product",
			HandlerFunc: handlers.WithMetrics(logging.Log, handlers.CreateProduct(env.DB)),
		},
		Route{
			Name:        "UpdateProduct",
			Method:      "PUT",
			Pattern:     "/product/{id:[0-9]+}",
			HandlerFunc: handlers.WithMetrics(logging.Log, handlers.UpdateProduct(env.DB)),
		},
		Route{
			Name:        "DeleteProduct",
			Method:      "DELETE",
			Pattern:     "/product/{id:[0-9]+}",
			HandlerFunc: handlers.WithMetrics(logging.Log, handlers.DeleteProduct(env.DB)),
		},
		Route{
			Name:    "ListProduct",
			Method:  "GET",
			Pattern: "/products",
			//HandlerFunc: handlers.ListProducts(env.DB),
			HandlerFunc: handlers.WithMetrics(logging.Log, handlers.ListProducts(env.DB)),
		},
	}

	// Register the routes.
	for _, route := range routes {
		env.Router.
			Handle(route.Pattern, route.HandlerFunc).
			Name(route.Name).Methods(route.Method)
	}

}
