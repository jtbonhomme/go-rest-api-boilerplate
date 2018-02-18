package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	db "github.com/jtbonhomme/go-rest-api-boilerplate/db"
)

// IDParam is used to identify a person
//
// swagger:parameters listPerson
type IDParam struct {
	// The ID of a person
	//
	// in: path
	// required: true
	ID int64 `json:"id"`
}

// GetPeople is an httpHandler for route GET /people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /people people listPeople
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
	//     Responses:
	//       default: genericError
	//       200: okResponse
	//       404: notFoundError
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
	//     Params:
	//       id: IDParam
	//
	//     Responses:
	//       default: genericError
	//       200: okResponse
	//       404: notFoundError
	params := mux.Vars(r)
	for _, item := range db.Get() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}
