package router

import (
	"github.com/gorilla/mux"
)

// NewRouter builds and returns a new router from routes
func NewRouter() *mux.Router {
	// When StrictSlash == true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa.
	router := mux.NewRouter().StrictSlash(true)
	router.Use(Logger)
	sub := router.PathPrefix("/v1").Subrouter()

	for _, route := range routes {
		sub.
			HandleFunc(route.Pattern, route.HandlerFunc).
			Name(route.Name).
			Methods(route.Method)
	}

	return router
}
