package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carlqt/jira-ya/models"
)

// Get /issues returns all Issues
func GetIssuesHandler(w http.ResponseWriter, r *http.Request) {
	queryType := r.URL.Query().Get("type")
	issues, err := models.AllIssues()
	issues = issues.FilterByType(queryType)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(issues)
	}
}
