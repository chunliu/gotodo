package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
