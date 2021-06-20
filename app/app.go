package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"

	l "github.com/annabkr/paydayz/utils/logger"
)

var app *App
type App struct { 
	port   string
	router *mux.Router
	pool *pgxpool.Pool
}

func Initialize() *App {
	app = &App{
		router: mux.NewRouter(),
		port:   ":3000",
	}
	app.initializeDB()

	return app
} 

func (a *App) initializeDB(){ 
	p, err := pgxpool.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil || p == nil {
		l.Err(fmt.Sprintf("unable to connect to postgres pool: %s", err.Error()))
		os.Exit(1)
	}
	// defer func(){
	// 	l.Info("Closing pool")
	// 	p.Close()
	// }()
	
	l.Info("Initialized PostgreSQL pool")
	a.pool = p 
}

func (a App) GetRouter() *mux.Router {
	return a.router
}

func (a App) Run() {
	l.Info(fmt.Sprintf("Listening on port %s", a.port))
	log.Fatal(http.ListenAndServe(a.port, a.router))
}

func GetPool() *pgxpool.Pool {
	return app.pool
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
