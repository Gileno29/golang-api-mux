package person

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID      string   `json:"id"`
	Name    string   `json:"Name"`
	Address *Address `json:"adress"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"State"`
}

var People []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(People)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range People {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	People = append(People, person)
	json.NewEncoder(w).Encode(People)

}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range People {
		if item.ID == params["id"] {
			People = append(People[:index], People[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(People)

	}
}
