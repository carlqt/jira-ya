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

// https://sephora-asia.atlassian.net/browse/EPS-676
type Issue struct {
	Id          string `json:"id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Assignee    string `json:"asignee"`
	Type        string `json:"type"` //EPS or SE
	Key         string `json:"key"`
	Link        string `json:"link"`
}

func AllIssues() ([]Issue, error) {
	var issues []Issue
	jiraIssues, err := jira.GetIssues()

	if err != nil {
		return issues, err
	}

	for _, v := range jiraIssues.Issues {
		issue := Issue{Type: "EPS"}
		issue.Id = v.Id
		issue.Key = v.Key
		issue.Description = v.Fields.Description
		issue.Summary = v.Fields.Summary
		issue.Assignee = v.Fields.Assignee.DisplayName
		issue.Link = v.Link()

		issues = append(issues, issue)
	}

	return issues, err
}

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
		issues, err := AllIssues()

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
