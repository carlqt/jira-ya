package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type IssueType string

const (
	SE  IssueType = "SE"
	EPS IssueType = "EPS"
)

type SearchResponse struct {
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	StartAt    int     `json:"startAt"`
	Issues     []Issue `json:"issues"`
}

type Issue struct {
	Id     string     `json:"id"`
	Key    string     `json:"key"`
	Self   string     `json:"self"`
	Fields IssueField `json:"fields"`
}

type IssueField struct {
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	Assignee    Assignee `json:"assignee"`
}

type Assignee struct {
	Id          string `json:"accountId"`
	DisplayName string `json:"displayName"`
}

type SearchRequest struct {
	Jql        string   `json:"jql"`
	MaxResults int      `json:"maxResults"`
	Fields     []string `json:"fields"`
}

func (i *Issue) Link() string {
	link := fmt.Sprintf("https://sephora-asia.atlassian.net/browse/%s", i.Key)

	return link
}

func GetIssues(issueType IssueType) (SearchResponse, error) {
	var searchResponse SearchResponse
	req, _ := jiraRequest(issueType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return searchResponse, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return searchResponse, err
	}

	return searchResponse, nil
}

func jiraRequest(issueType IssueType) (*http.Request, error) {
	var jql string
	username := "***REMOVED***"
	accessToken := "***REMOVED***"
	url := "https://sephora-asia.atlassian.net/rest/api/latest/search"

	if issueType == SE {
		jql = seJql()
	} else {
		jql = epsJql()
	}

	requestBody := SearchRequest{
		Jql:        jql,
		Fields:     []string{"summary", "description", "status", "assignee"},
		MaxResults: 50,
	}
	jsonBody, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))

	if err != nil {
		return req, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, accessToken)

	return req, err
}

func seJql() string {
	return "project=SE AND sprint in openSprints()"
}

func epsJql() string {
	return "project=EPS AND labels=SE AND resolution=Unresolved"
}
