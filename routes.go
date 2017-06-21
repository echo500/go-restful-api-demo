package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Rountes []Route

func NewRounter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Rountes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", TodoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", TodoShow},
}
