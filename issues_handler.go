package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/carlqt/jira-ya/jira"
)

// https://sephora-asia.atlassian.net/browse/EPS-676
// Response schema
type Issue struct {
	Id          string `json:"id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Assignee    string `json:"asignee"`
	Type        string `json:"type"` //EPS or SE
	Key         string `json:"key"`
	Link        string `json:"link"`
}

func AllIssues(issueType jira.IssueType) ([]Issue, error) {
	var issues []Issue
	jiraIssues, err := jira.GetIssues(issueType)

	if err != nil {
		return issues, err
	}

	for _, v := range jiraIssues.Issues {
		issue := Issue{Type: string(issueType)}
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

// Get /issues returns all Issues
func GetIssues() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		issueType := getIssueType(r.URL.Query())
		issues, err := AllIssues(issueType)

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

func getIssueType(query url.Values) jira.IssueType {
	typeParam := query.Get("type")

	if typeParam != "" {
		upcaseTypeParam := strings.ToUpper(typeParam)
		return jira.IssueType(upcaseTypeParam)
	} else {
		return jira.EPS
	}
}
