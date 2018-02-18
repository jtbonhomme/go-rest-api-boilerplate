package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	db "github.com/jtbonhomme/go-rest-api-boilerplate/db"
)

// GetPeople is an httpHandler for route GET /people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /people people listPersons
	//
	// Lists all people.
	//
	// This will show all available recorded people.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key: read, write
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       403: validationError
	json.NewEncoder(w).Encode(db.Get())
}

// GetPerson is an httpHandler for route GET /people/:id
func GetPerson(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /people/{id} people listPerson
	//
	// Lists person with id {id}.
	//
	// This will show the record of an identified person.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Schemes: http, https
	//
	//     Security:
	//       api_key: read, write
	//       oauth: read, write
	//
	//     Responses:
	//       default: genericError
	//       200: someResponse
	//       403: validationError
	params := mux.Vars(r)
	for _, item := range db.Get() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}
