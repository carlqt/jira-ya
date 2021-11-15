package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carlqt/jira-ya/jira"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetIssuesHandler(t *testing.T) {
	godotenv.Load(".env")
	mockApp := App{
		JiraConfig: jira.DefaultJiraConfig(),
	}

	req, _ := http.NewRequest("GET", "/issues", nil)
	resp := httptest.NewRecorder()
	mockApp.GetIssuesHandler(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "is equal")
}
