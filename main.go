package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        int    `json:"id"`
	FULL_NAME string `json:"name"`
}

var persons []Person

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(1)
}

func getPersons(w http.ResponseWriter, r *http.Request) {
	fmt.Println(persons)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func byId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range persons {
		if "1" == params["id"] {
			fmt.Println(item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person1 Person
	_ = json.NewDecoder(r.Body).Decode(&person1)
	person1.ID = rand.Intn(10000)
	persons = append(persons, person1)
	json.NewEncoder(w).Encode(person1)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range persons {
		if "1" == params["id"] {
			item.FULL_NAME = "updated name"
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	fmt.Println("Hello To REST/API")
	// Init Router
	router := mux.NewRouter()

	persons = append(persons, Person{ID: 1, FULL_NAME: "PHP"})
	persons = append(persons, Person{ID: 2, FULL_NAME: "LARAVEL"})
	persons = append(persons, Person{ID: 3, FULL_NAME: "NODE"})
	persons = append(persons, Person{ID: 4, FULL_NAME: "PYTHON"})
	persons = append(persons, Person{ID: 5, FULL_NAME: "GO"})

	// EndPoint
	router.HandleFunc("/api", greeting).Methods("GET")
	router.HandleFunc("/api/persons", getPersons).Methods("GET")
	router.HandleFunc("/api/persons/{id}", byId).Methods("GET")
	router.HandleFunc("/api/persons", createPerson).Methods("POST")
	router.HandleFunc("/api/persons/{id}", updatePerson).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
