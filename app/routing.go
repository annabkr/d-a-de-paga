package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/annabkr/dia-de-paga/api"
	l "github.com/annabkr/dia-de-paga/utils/logger"
	"github.com/gorilla/mux"
)

type App struct {
	router *mux.Router
	port   string
}

func Initialize() *App {
	app := &App{
		router: mux.NewRouter(),
		port:   ":3000",
	}

	app.initializeAPI()
	return app
}

func (a App) GetRouter() *mux.Router {
	return a.router
}

func (a App) RegisterRoute(method string, pattern string, handler HandlerFunc) {
	a.GetRouter().Handle(pattern, handler).Methods(method)
}

func (a App) Run() {
	l.Info(fmt.Sprintf("Listening on port %s", a.port))
	log.Fatal(http.ListenAndServe(a.port, a.router))
}

func (a App) initializeAPI() {
	routes := api.GetApiRoutes()
	for _, route := range routes {
		a.RegisterRoute(route.Method, route.Pattern, route.Handler)
	}
}
