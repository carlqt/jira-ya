package main

import "github.com/gorilla/mux"

// RESTFUL handler name
// GetIssues
// GetIssue
// DeleteIssue
// UpdateIssue
// CreateIssue
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/issues", GetIssuesHandler).Methods("GET")
	router.Use(ResponseHeaderHandler)

	return router

}
