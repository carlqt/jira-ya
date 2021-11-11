package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/issues", GetIssues()).Methods("GET")
	router.Use(ResponseHeaderHandler)

	return router

}
