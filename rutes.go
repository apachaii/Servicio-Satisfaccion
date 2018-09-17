package main

import(
	"net/http"
	"github.com/gorilla/mux"
)
type Route struct {
	Name string
	Method string
	Pattern string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
		      Name(route.Name).		  
			  Path(route.Pattern).
			  Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes {

	Route {
		"SatisList",
		"Get",
		"/peliculas",
		SatisList,
	},

}