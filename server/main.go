package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/carlqt/jira-ya/jira"
	"github.com/carlqt/jira-ya/models"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type App struct {
	Issues interface {
		AllIssues() (models.Issues, error)
	}
}

func (a *App) Start() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	port := os.Getenv("PORT")
	corsOptions := cors.New(cors.Options{
		AllowedMethods: []string{"OPTIONS", "GET", "POST"},
		AllowedHeaders: []string{"Content-Type"},
	})

	loggedRouter := handlers.LoggingHandler(os.Stdout, a.NewRouter())

	log.Printf("starting at port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), corsOptions.Handler(loggedRouter))
}

func NewApp() *App {
	app := new(App)
	app.Issues = models.IssueModel{Config: jira.DefaultJiraConfig()}

	return app
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	} else {
		app := NewApp()
		app.Start()
	}
}
