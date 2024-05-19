package main

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

var people []Person

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}
