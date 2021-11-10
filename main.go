package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/carlqt/jira-ya/jira"
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

func GetIssues() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issues, err := jira.GetIssues()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(issues)
		}

	})
}

func ResponseHeaderHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
