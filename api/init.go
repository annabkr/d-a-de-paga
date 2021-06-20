package api

import (
	"github.com/gorilla/mux"

	s "github.com/annabkr/paydayz/app/server"
)

func Initialize(r *mux.Router){
	routes := GetApiRoutes()
	for _, route := range routes {
		RegisterRoute(r, route.Method, route.Pattern, route.Handler)
	}
}

func RegisterRoute(r *mux.Router, method string, pattern string, handler s.HandlerFunc) {
	r.Handle(pattern, handler).Methods(method)
}
