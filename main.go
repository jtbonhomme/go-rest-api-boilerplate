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
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//     oauth2:
//         type: oauth2
//         authorizationUrl: /oauth2/auth
//         tokenUrl: /oauth2/token
//         in: header
//         scopes:
//           bar: foo
//         flow: accessCode
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	db "github.com/jtbonhomme/go-rest-api-boilerplate/db"
	handler "github.com/jtbonhomme/go-rest-api-boilerplate/handlers"
	model "github.com/jtbonhomme/go-rest-api-boilerplate/model"
)

// Logger is a gorilla/mux middleware to add log to the API
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}

// @SubApi People [/people]
// @SubApi Allows you access to different features of the persons, name, address, etc [/people]

// our main function
func main() {
	db.Insert(model.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &model.Address{City: "City X", State: "State X"}})
	db.Insert(model.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &model.Address{City: "City Z", State: "State Y"}})
	db.Insert(model.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	// When StrictSlash == true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa.
	router := mux.NewRouter().StrictSlash(true)
	router.Use(Logger)
	sub := router.PathPrefix("/v1").Subrouter()
	sub.HandleFunc("/people", handler.GetPeople).Methods("GET")
	sub.HandleFunc("/people/{id}", handler.GetPerson).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
