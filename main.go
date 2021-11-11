package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	router := mux.NewRouter()
	router.Handle("/issues", GetIssues()).Methods("GET")
	router.Use(ResponseHeaderHandler)

	corsOptions := cors.New(cors.Options{
		// AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	})

	log.Println("starting at port 8000")
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	http.ListenAndServe(":8000", corsOptions.Handler(loggedRouter))

}
