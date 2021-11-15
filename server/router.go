package main

import "github.com/gorilla/mux"

// RESTFUL handler name
// GetIssues
// GetIssue
// DeleteIssue
// UpdateIssue
// CreateIssue
func (a *App) NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/issues", a.GetIssuesHandler).Methods("GET")
	router.Use(ResponseHeaderHandler)

	return router

}
