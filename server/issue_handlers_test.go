package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIssuesHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/issues", nil)
	resp := httptest.NewRecorder()
	GetIssuesHandler(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "is equal")
}
