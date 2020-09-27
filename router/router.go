package router

import (
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Path("/slack").Methods("POST").HandlerFunc(slackHandler)
	router.Path("/").Methods("GET").HandlerFunc(indexHandler)

	router.Path("/slack/test").Methods("POST").HandlerFunc(testSlash)
	router.Path("/slack/add-project").Methods("POST").HandlerFunc(addProject)
	router.Path("/slack/all-projects").Methods("POST").HandlerFunc(allProjects)

	return router
}
