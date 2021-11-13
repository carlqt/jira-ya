package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func (a *App) Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	port := os.Getenv("PORT")
	corsOptions := cors.New(cors.Options{
		// AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	})

	loggedRouter := handlers.LoggingHandler(os.Stdout, a.Router)

	log.Printf("starting at port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), corsOptions.Handler(loggedRouter))
}

func NewApp() *App {
	app := new(App)
	app.Router = NewRouter()

	return app
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	} else {
		app := NewApp()
		app.Start()
	}
}
