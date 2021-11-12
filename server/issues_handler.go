package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/carlqt/jira-ya/jira"
)

// https://sephora-asia.atlassian.net/browse/EPS-676
// Response schema
type Issue struct {
	Id          string `json:"id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
	Type        string `json:"type"` //EPS or SE
	Key         string `json:"key"`
	Link        string `json:"link"`
}

type Issues []Issue

func (issues Issues) filterByType(t string) Issues {
	var newIssues Issues

	if t == "" {
		return issues
	}

	for _, issue := range issues {
		if issue.Type == strings.ToUpper(t) {
			newIssues = append(newIssues, issue)
		}
	}

	return newIssues
}

func AllIssues() (Issues, error) {
	var issues Issues
	jiraIssues, err := jira.GetIssues()

	if err != nil {
		return issues, err
	}

	for _, v := range jiraIssues.Issues {
		k := strings.Split(v.Key, "-")
		issue := Issue{
			Id:          v.Id,
			Summary:     v.Fields.Summary,
			Description: v.Fields.Description,
			Assignee:    v.Fields.Assignee.DisplayName,
			Type:        k[0],
			Key:         v.Key,
			Link:        v.Link(),
		}

		issues = append(issues, issue)
	}

	return issues, err
}

// Get /issues returns all Issues
func GetIssues() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		queryType := r.URL.Query().Get("type")
		issues, err := AllIssues()
		issues = issues.filterByType(queryType)

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
