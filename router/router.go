package router

import (
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Path("/slack").Methods("POST").HandlerFunc(slackHandler)
	router.Path("/slack/test").Methods("POST").HandlerFunc(testHandler)
	router.Path("/").Methods("GET").HandlerFunc(indexHandler)

	return router
}
