package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
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

	if r.Header.Get("Content-Type") == "application/xml" {
		// set response content-type header
		w.Header().Add("Content-Type", "application/xml")

		err := xml.NewEncoder(w).Encode(customers)
		if err != nil {
			return
		}
	} else {
		// set response content-type header
		w.Header().Add("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(customers)
		if err != nil {
			return
		}
	}

}
