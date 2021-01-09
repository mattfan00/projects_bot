package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"bot/database"
	"bot/slack"
	"bot/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type Project struct {
	Name        string `bson:"name"`
	Url         string `bson:"url"`
	Description string `bson:"description"`
	CreatedBy   string `bson:"created_by"`
}

func addProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add a new project")
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	text := r.FormValue("text")
	fmt.Println(text)

	// -n "name" -u url -d "description"
	args := utils.GetArgs(text)

	newProject := Project{
		Name:        args["Name"],
		Url:         args["Url"],
		Description: args["Description"],
		CreatedBy:   r.FormValue("user_id"),
	}

	createdProject, err := database.Db.Collection("projects").InsertOne(context.TODO(), newProject)
	if err != nil {
		panic(err)
	}

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
		str := fmt.Sprintf("*%s* - <@%s>\n%s\n%s\n", project.Name, project.CreatedBy, project.Url, project.Description)
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
