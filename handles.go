package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// fmt.Fprint(w, "Welcome to the todo list app!\n")
	t, err := template.ParseFiles("pages/home.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	if err := t.Execute(w, nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Add necessary header to the response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todoRepo.Items); err != nil {
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

	_, t := todoRepo.Find(i)
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

func readBody(r *http.Request) (Todo, error) {
	b, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var t Todo
	err = json.Unmarshal(b, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := readBody(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	t = todoRepo.Add(t)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		panic(err)
	}
}

func update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	t, err := readBody(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	i, todo := todoRepo.Find(id)
	if todo == (Todo{}) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		todoRepo.Update(i, t)
		w.WriteHeader(http.StatusNoContent)
	}
}

func delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	i, todo := todoRepo.Find(id)
	if todo == (Todo{}) {
		w.WriteHeader(http.StatusNotFound)
	} else {
		todoRepo.Delete(i)
		w.WriteHeader(http.StatusNoContent)
	}
}
