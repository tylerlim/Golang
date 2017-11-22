package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// The person Type (more like an object)
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

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// our main function
func main() {
    router := mux.NewRouter()
    people = append(people, Person{ID: "1", Firstname: "Tyler", Lastname: "Lim", Address: &Address{City: "Bayan Baru", State: "Pulau Pinang"}})
    router.HandleFunc("/people", GetPeople).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}