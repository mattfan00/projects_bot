package router

import (
	"bot/handlers"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Path("/slack").Methods("POST").HandlerFunc(handlers.SlackHandler)
	router.Path("/").Methods("GET").HandlerFunc(handlers.IndexHandler)

	router.Path("/slack/add-project").Methods("POST").HandlerFunc(handlers.AddProject)
	router.Path("/slack/all-projects").Methods("POST").HandlerFunc(handlers.AllProjects)
	router.Path("/slack/delete-project").Methods("POST").HandlerFunc(handlers.DeleteProject)

	return router
}
