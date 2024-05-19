package main

import (
	"log"
	"net/http"

	"github.com/Gileno29/golang-api-mux/person"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	p := person.Person{ID: "1", Name: "Gileno", Address: &person.Address{City: "Natal", State: "RN"}}
	p2 := person.Person{ID: "2", Name: "Paulina", Address: &Address{City: "Natal", State: "RN"}}
	var people []person.Person

	people = append(people, p)
	people = append(people, p2)

	router.HandleFunc("/contato", person.getPeople()).Methods("GET")

	router.HandleFunc("/contato/{id}", getPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
