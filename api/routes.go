package api

import (
	"net/http"
)

type ApiRoute struct {
	Method  string
	Pattern string
	Handler func(w http.ResponseWriter, r *http.Request) error
}

func GetApiRoutes() []ApiRoute {
	return []ApiRoute{
		GetRoot,
		GetRecord,
		PostRecord,
	}
}

var GetRoot = ApiRoute{
	"GET",
	"/",
	root,
}

var GetRecord = ApiRoute{
	"GET",
	"/record",
	getRecord,
}

var PostRecord = ApiRoute{
	"POST",
	"/record",
	postRecord,
}