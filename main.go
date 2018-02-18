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
//     Host: localhost:8000
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
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}

type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

func Logger(r *http.Request) {
    log.Printf(
        "\t%s\t%s",
        r.Method,
        r.RequestURI,
    )
}
// @SubApi People [/people]
// @SubApi Allows you access to different features of the persons, name, address, etc [/people]

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
func GetPeople(w http.ResponseWriter, r *http.Request) {
    Logger(r);
    json.NewEncoder(w).Encode(people)
}

// @Title Get Person Information
// @Description Get Person Information
// @Accept json
// @Param id path int true &quot;Person ID&quot;
// @Success 200 {object} string &quot;Success&quot;
// @Failure 401 {object} string &quot;Access denied&quot;
// @Failure 404 {object} string &quot;Not Found&quot;
// @Resource /people
// @Router /v1/people/:id [get]
func GetPerson(w http.ResponseWriter, r *http.Request) {
    Logger(r);
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
        }
    }
}

// our main function
func main() {
    people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
    people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
    people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

    router := mux.NewRouter()
    sub := router.PathPrefix("/v1").Subrouter()
    sub.HandleFunc("/people", GetPeople).Methods("GET")
    sub.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

