package app

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.method).Path(route.path).Name(route.name).Handler(route.handler)
	}
	return router
}
