package main

import (
	"github.com/julienschmidt/httprouter"
)

// Route represents the model for httprouter handles
type Route struct {
	Method  string
	Pattern string
	Handle  httprouter.Handle
}

// Routes is a slice for all routes in the app
type Routes []Route

func newRoutes() *httprouter.Router {
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
	Route{
		Method:  "POST",
		Pattern: "/todo",
		Handle:  createItem,
	},
}
