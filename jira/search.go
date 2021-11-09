package jira

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SearchResponse struct {
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	StartAt    int     `json:"startAt"`
	Issues     []Issue `json:"issues"`
}

type Issue struct {
	Id   string `json:"id"`
	Key  string `json:"key"`
	Self string `json:"self"`
}

type SearchRequest struct {
	Jql        string   `json:"jql"`
	MaxResults int      `json:"maxResults"`
	Fields     []string `json:"fields"`
}

func GetIssues() (SearchResponse, error) {
	var searchResponse SearchResponse
	url := "https://sephora-asia.atlassian.net/rest/api/latest/search"

	requestBody := SearchRequest{
		Jql:        "project=EPS AND labels=SE AND resolution=Unresolved",
		Fields:     []string{"summary", "description", "status"},
		MaxResults: 5,
	}

	jsonBody, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("***REMOVED***", "***REMOVED***")

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
