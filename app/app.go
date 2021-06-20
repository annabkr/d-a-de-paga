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

var pool *pgxpool.Pool
type App struct {
	router *mux.Router
	port   string 
}

func Initialize() *App {
	app := &App{
		router: mux.NewRouter(),
		port:   ":3000",
	}
	app.initializeDB()
	app.initializeAPI()
	return app
}

func (a App) initializeDB(){
	p, err := pgxpool.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		l.Err(fmt.Sprintf("unable to connect to postgres pool: %s", err.Error()))
		os.Exit(1)
	}
	defer pool.Close()
	
	pool = p
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

func GetPool() *pgxpool.Pool {
	return pool
}

// Schema
// CREATE TABLE transactions(
//     id SERIAL PRIMARY KEY,
//     amount FLOAT NOT NULL,
//     source TEXT NOT NULL
// )

// id  | amount |  source  
// -----+--------+----------
//  123 |  100.5 | Paycheck
