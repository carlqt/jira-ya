package models

import (
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

func (issues Issues) FilterByType(t string) Issues {
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

func AllIssues(c *jira.JiraConfig) (Issues, error) {
	var issues Issues
	jiraIssues, err := jira.GetIssues(c)

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
