package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/annabkr/paydayz/api"
	l "github.com/annabkr/paydayz/utils/logger"
)

type App struct {
	router *mux.Router
	port   string
	db     *pgxpool.Pool
}

func Initialize() *App {
	db, err := pgxpool.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		l.Err(fmt.Sprintf("unable to connect to database: %s", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	app := &App{
		router: mux.NewRouter(),
		port:   ":3000",
		db:     db,
	}

	app.initializeAPI()
	return app
}

func (a App) initializeAPI() {
	routes := api.GetApiRoutes()
	for _, route := range routes {
		a.RegisterRoute(route.Method, route.Pattern, route.Handler)
	}
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

func (a App) GetDatabase() *pgxpool.Pool {
	return a.db
}
