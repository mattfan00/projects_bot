package main

import (
	"net/http"

	"bot/router"
)

func main() {
	// PostMessage("no more messages")

	router := router.CreateRouter()

	http.ListenAndServe(":8080", router)

}
