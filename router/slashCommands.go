package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"bot/slack"
)

type Project struct {
	Name      string
	CreatedBy string
}

var Projects []Project

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

	Projects = append(Projects, newProject)

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

	var text string
	for _, project := range Projects {
		str := fmt.Sprintf("%s - <@%s>", project.Name, project.CreatedBy)
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
