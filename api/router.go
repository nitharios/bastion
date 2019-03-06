package router

import (
	"net/http"

	"github.com/gorilla/mux"
	v0 "github.com/nitharios/bastion/api/v0"
	log "github.com/nitharios/bastion/logger"
)

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range v0.Router {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = log.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
