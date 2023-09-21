package controller

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "CreatePerson")
}

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetPersonById")
}

func GetPersonByTerm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetPersonByTerm")
}

func GetPersonCount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetPersonCount")
}
