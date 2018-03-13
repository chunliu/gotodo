package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func getAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Add necessary header to the response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todoItems); err != nil {
		panic(err)
	}
}

func getByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// input param must be an int
	i, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t := RepoFindTodo(i)
	if t == (Todo{}) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/todo/", getAll)
	router.GET("/todo/:id", getByID)

	log.Fatal(http.ListenAndServe(":8080", router))
}
