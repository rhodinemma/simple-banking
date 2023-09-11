package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// define routes
	http.HandleFunc("/greet", greet)

	// starting the server
	log.Fatalln(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello Rhodin GO......")
	if err != nil {
		return
	}
}
