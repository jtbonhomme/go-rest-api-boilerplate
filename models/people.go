package model

// Person description.
// swagger:model person
type Person struct {
	// ID of the person
	//
	// required: true
	ID string `json:"id,omitempty"`
	// Firstname of the person
	//
	// required: true
	Firstname string `json:"firstname,omitempty"`
	// Lastname of the person
	//
	// required: true
	Lastname string `json:"lastname,omitempty"`
	// Address of the person
	//
	// required: false
	Address *Address `json:"address,omitempty"`
}

// Address description
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}
