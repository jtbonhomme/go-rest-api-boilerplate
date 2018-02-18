package db

import (
	model "github.com/jtbonhomme/go-rest-api-boilerplate/model"
)

var people []model.Person

// Insert
func Insert(person model.Person) {
	people = append(people, person)
}

// Insert
func Get() []model.Person {
	return people
}
