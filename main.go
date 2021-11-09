package main

import (
	"fmt"

	"github.com/carlqt/jira-ya/jira"
)

func main() {
	issues, err := jira.GetIssues()
	if err != nil {
		panic(err)
	}

	fmt.Println(issues)
}
