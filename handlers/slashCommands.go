package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"bot/database"
	"bot/helpers"
	"bot/models"

	"go.mongodb.org/mongo-driver/bson"
)

func AddProject(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add a new project")
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	/*
		text := r.FormValue("text")

		// -n "name" -u url -d "description"
		args, help := helpers.GetArgs(text)

		var messageText string
		if !help {
			newProject := models.Project{
				Name:        args.Name,
				Url:         args.Url,
				Description: args.Description,
				CreatedBy:   r.FormValue("user_id"),
			}

			_, err := database.Db.Collection("projects").InsertOne(context.TODO(), newProject)
			if err != nil {
				panic(err)
			}

			messageText = "Created a new project!"
		} else {
			messageText = "*Example usage: * \n /add-project -n \"Project name\" -u https://www.github.com -d \"Project description\" \n\n *Options:* \n-n, --name=\tName of your project. Required.\n-u, --url=\tGitHub URL of project. \n-d, --description=\tA short description of the project."
		}

		newMessage := models.SlashMessage{
			ResponseType: "ephemeral",
			Text:         messageText,
		}
	*/

	b := []byte(fmt.Sprintf(`{
	"trigger_id": "%s",
	"view": {
		"title": {
			"type": "plain_text",
			"text": "Add a new project",
			"emoji": true
		},
		"submit": {
			"type": "plain_text",
			"text": "Submit",
			"emoji": true
		},
		"type": "modal",
		"close": {
			"type": "plain_text",
			"text": "Cancel",
			"emoji": true
		},
		"blocks": [
			{
				"type": "input",
				"element": {
					"type": "plain_text_input",
					"action_id": "name"
				},
				"label": {
					"type": "plain_text",
					"text": "Name (Required)",
					"emoji": true
				}
			},
			{
				"type": "input",
				"element": {
					"type": "plain_text_input",
					"action_id": "url"
				},
				"label": {
					"type": "plain_text",
					"text": "URL",
					"emoji": true
				}
			},
			{
				"type": "input",
				"element": {
					"type": "plain_text_input",
					"multiline": true,
					"action_id": "description"
				},
				"label": {
					"type": "plain_text",
					"text": "Description",
					"emoji": true
				}
			}
		]
	}
	}`, r.FormValue("trigger_id")))

	var newMessage interface{}
	if err := json.Unmarshal(b, &newMessage); err != nil {
		panic(err)
	}

	helpers.NewPostRequest("https://slack.com/api/views.open", newMessage)
}

func AllProjects(w http.ResponseWriter, r *http.Request) {
	fmt.Println("view all projects")

	cur, err := database.Db.Collection("projects").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var text string
	for cur.Next(context.TODO()) {
		var project models.Project
		err := cur.Decode(&project)
		if err != nil {
			panic(err)
		}
		str := fmt.Sprintf("*%s* - <@%s>\n%s\n%s\n", project.Name, project.CreatedBy, project.Url, project.Description)
		text = text + "\n" + str
	}

	newMessage := models.SlashMessage{
		ResponseType: "ephemeral",
		Text:         text,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newMessage)
}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	newMessage := map[string]interface{}{
		"response_type": "ephemeral",
		"blocks": []interface{}{
			map[string]interface{}{
				"type": "section",
				"text": map[string]interface{}{
					"type": "mrkdwn",
					"text": "*Choose a project to delete:*",
				},
			},
			map[string]interface{}{
				"type":     "actions",
				"elements": []interface{}{},
			},
		},
	}

	cur, err := database.Db.Collection("projects").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	for cur.Next(context.TODO()) {
		var project models.Project
		err := cur.Decode(&project)
		if err != nil {
			panic(err)
		}

		newElement := map[string]interface{}{
			"type": "button",
			"text": map[string]interface{}{
				"type": "plain_text",
				"text": project.Name,
			},
			"value":     "delete",
			"action_id": project.Id,
		}

		a := newMessage["blocks"].([]interface{})[1].(map[string]interface{})["elements"].([]interface{})

		newMessage["blocks"].([]interface{})[1].(map[string]interface{})["elements"] = append(a, newElement)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newMessage)
}
