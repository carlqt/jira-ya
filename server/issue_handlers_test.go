package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/carlqt/jira-ya/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type mockIssuesModel struct{}

func (m mockIssuesModel) AllIssues() (models.Issues, error) {
	var issues models.Issues

	issues = append(issues, models.Issue{})

	return issues, nil
}

func TestGetIssuesHandler(t *testing.T) {
	godotenv.Load(".env")
	mockApp := App{
		Issues: mockIssuesModel{},
	}

	req, _ := http.NewRequest("GET", "/issues", nil)
	resp := httptest.NewRecorder()
	mockApp.GetIssuesHandler(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "is equal")
}
