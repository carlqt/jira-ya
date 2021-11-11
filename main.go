package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func (a *App) Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	corsOptions := cors.New(cors.Options{
		// AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	})

	loggedRouter := handlers.LoggingHandler(os.Stdout, a.Router)

	log.Println("starting at port 8000")
	http.ListenAndServe(":8000", corsOptions.Handler(loggedRouter))
}

func NewApp() *App {
	app := new(App)
	app.Router = NewRouter()

	return app
}

func main() {
	app := NewApp()
	app.Start()
}
