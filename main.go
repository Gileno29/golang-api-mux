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
	p2 := person.Person{ID: "2", Name: "Paulina", Address: &person.Address{City: "Natal", State: "RN"}}

	person.People = append(person.People, p)
	person.People = append(person.People, p2)

	router.HandleFunc("/contato", person.GetPeople).Methods("GET")

	router.HandleFunc("/contato/{id}", person.GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", person.CreatePerson).Methods("POST")
	router.HandleFunc("/contato/{id}", person.DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
