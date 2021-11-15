package jira

import (
	"fmt"
	"net/url"
	"os"
)

type JiraConfig struct {
	Username    string
	AccessToken string
	Version     string
}

func DefaultJiraConfig() *JiraConfig {
	config := &JiraConfig{
		Username:    os.Getenv("USERNAME"),
		AccessToken: os.Getenv("ACCESS_TOKEN"),
		Version:     "latest",
	}

	return config
}

func (c *JiraConfig) URL() *url.URL {
	scheme := "https"
	baseURL := "sephora-asia.atlassian.net"
	stringUrl := fmt.Sprintf("%s://%s/rest/api/%s", scheme, baseURL, c.Version)

	url, _ := url.Parse(stringUrl)
	return url
}
