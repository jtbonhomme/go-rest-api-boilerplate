//go:generate swagger generate spec

// Package classification People API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:3000
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Jean-Thierry Bonhomme <jtbonhomme@gmail.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	db "github.com/jtbonhomme/go-rest-api-boilerplate/db"
	model "github.com/jtbonhomme/go-rest-api-boilerplate/models"
	"github.com/jtbonhomme/go-rest-api-boilerplate/router"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler

	return handleCORS(handler)
}

// our main function
func main() {
	db.Insert(model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "City X", State: "State X"}})
	db.Insert(model.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &model.Address{City: "City Z", State: "State Y"}})
	db.Insert(model.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", setupGlobalMiddleware(router)))
}
