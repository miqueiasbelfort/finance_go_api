package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func ConfigRouter(r *mux.Router) *mux.Router {
	routers := userRouters

	for _, route := range routers {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
