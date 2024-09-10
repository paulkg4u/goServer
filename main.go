package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)



type Founder struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Email string `json:"email"`
	Company string `json:"company"`
}

var founders []Founder



func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from GoServer 👋🏻")
}
func formHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var founder Founder
	json.NewDecoder(r.Body).Decode(&founder)
	founder.Age = founder.Age + 1
	founders = append(founders, founder)
	json.NewEncoder(w).Encode(founders)
}


func main() {
	r := mux.NewRouter()
	founders = append(founders, Founder{Name: "Paul", Age: 20, Email: "paul@example.com", Company: "Google"})
	r.HandleFunc("/", greetingsHandler).Methods("GET")
	r.HandleFunc("/form", formHandler).Methods("POST")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}


// Gorilla/mux is archived right now, I would recommend everyone to use another router, like go-chi.io which can be used with (almost) the same code as the article!
