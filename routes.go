package main

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

type Routes []Route

func NewRoutes() *httprouter.Router {
	router := httprouter.New()
	for _, r := range routes {
		router.Handle(r.Method, r.Pattern, r.Handle)
	}

	return router
}

var routes = Routes{
	Route{
		Method:  "GET",
		Pattern: "/",
		Handle:  index,
	},
	Route{
		Method:  "GET",
		Pattern: "/todo",
		Handle:  getAll,
	},
	Route{
		Method:  "GET",
		Pattern: "/todo/:id",
		Handle:  getByID,
	},
}
