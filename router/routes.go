package router

import (
	"net/http"

	handler "github.com/jtbonhomme/go-rest-api-boilerplate/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetPeople",
		"GET",
		"/people",
		handler.GetPeople,
	},
	Route{
		"GetPerson",
		"GET",
		"/people/{id}",
		handler.GetPerson,
	},
}
