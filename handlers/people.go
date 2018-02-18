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
	// This will show all recorded people.
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
	//       200: peopleResponse
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Get())
}

// GetPerson is an httpHandler for route GET /people/{id}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /people/{id} people listPerson
	//
	// Lists person from their id.
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
	//       200: personResponse
	//       404: jsonError
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _, item := range db.Get() {
		if item.ID == params["id"] {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(item); err != nil {
				panic(err)
			}
			return
		}
	}
	// If we didn't find it, 404
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonError{Message: "Not Found"}); err != nil {
		panic(err)
	}
}
