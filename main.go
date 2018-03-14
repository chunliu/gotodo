package main

import (
	"log"
	"net/http"
)

func main() {
	router := newRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
