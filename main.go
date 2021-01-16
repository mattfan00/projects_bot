package main

import (
	"context"
	"fmt"
	"net/http"

	"bot/database"
	"bot/router"
)

func main() {
	db := database.Init()
	defer db.Disconnect(context.TODO())

	router := router.CreateRouter()

	fmt.Println("server ready at port :8080")
	http.ListenAndServe(":8080", router)
}
