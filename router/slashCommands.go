package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"bot/database"
	"bot/slack"
	"go.mongodb.org/mongo-driver/bson"
)

type Project struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
	CreatedBy   string `bson:"created_by"`
}

func testSlash(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testing this handler")
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	for key, value := range r.Form {
		fmt.Printf("%s: %s\n", key, value)
	}

	newMessage := slack.SlashMessage{
		ResponseType: "ephemeral",
		Text:         "whats up my dude\n not much",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newMessage)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add a new project")
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	newProject := Project{
		Name:      r.FormValue("text"),
		CreatedBy: r.FormValue("user_id"),
	}

	createdProject, err := database.Db.Collection("projects").InsertOne(context.TODO(), newProject)
	if err != nil {
		panic(err)
	}
	fmt.Println(createdProject)

	newMessage := slack.SlashMessage{
		ResponseType: "ephemeral",
		Text:         "Created a new project!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newMessage)
}

func allProjects(w http.ResponseWriter, r *http.Request) {
	fmt.Println("view all projects")

	cur, err := database.Db.Collection("projects").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var text string
	for cur.Next(context.TODO()) {
		var project Project
		err := cur.Decode(&project)
		if err != nil {
			panic(err)
		}
		fmt.Println(project)
		str := fmt.Sprintf("*%s* - <@%s>", project.Name, project.CreatedBy)
		text = text + "\n" + str
	}

	newMessage := slack.SlashMessage{
		ResponseType: "ephemeral",
		Text:         text,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newMessage)
}
