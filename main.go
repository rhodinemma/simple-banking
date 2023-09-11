package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {
	// define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// starting the server
	log.Fatalln(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello Rhodin GO......")
	if err != nil {
		return
	}
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Rhodin", "Mukono", "110075"},
		{"Emmanuel", "Kampala", "110075"},
		{"Mariam", "Seeta", "110075"},
	}

	// set response content-type header
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(customers)
	if err != nil {
		return
	}
}
