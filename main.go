package main

import (
	"log"
	"net/http"
)

func main() {
	// to add : url functionality
	// to add : color functionality
	log.Fatal(http.ListenAndServe(":8080", nil))
}
