package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bot/database"
	"bot/helpers"
	"bot/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from the index")
}

func SlackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reached the slack route")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var newReq models.SlackRequest

	if err := json.Unmarshal(body, &newReq); err != nil {
		panic(err)
	}

	if newReq.Challenge != "" {
		fmt.Printf("this is a challenge")
		// fmt.Fprintf(w, "%s", newReq.Challenge)
	} else {
		fmt.Printf("not a challenge")
		fmt.Printf("%s", newReq.Event.Text)
		helpers.PostMessage(newReq.Event.Channel, "hey dude")
	}
}

func InteractiveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("interactive")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	payload := r.FormValue("payload")

	var newReq interface{}
	if err := json.Unmarshal([]byte(payload), &newReq); err != nil {
		panic(err)
	}

	messageType := newReq.(map[string]interface{})["type"]

	if messageType == "view_submission" {
		inputs := newReq.(map[string]interface{})["view"].(map[string]interface{})["state"].(map[string]interface{})["values"].(map[string]interface{})
		newProject := models.Project{
			CreatedBy: newReq.(map[string]interface{})["user"].(map[string]interface{})["id"].(string),
		}
		for _, v := range inputs {
			if v.(map[string]interface{})["name"] != nil {
				newProject.Name = v.(map[string]interface{})["name"].(map[string]interface{})["value"].(string)
			} else if v.(map[string]interface{})["url"] != nil {
				newProject.Url = v.(map[string]interface{})["url"].(map[string]interface{})["value"].(string)
			} else if v.(map[string]interface{})["description"] != nil {
				newProject.Description = v.(map[string]interface{})["description"].(map[string]interface{})["value"].(string)
			}
		}

		_, err := database.Db.Collection("projects").InsertOne(context.TODO(), newProject)
		if err != nil {
			panic(err)
		}
	} else if messageType == "block_actions" {
		action := newReq.(map[string]interface{})["actions"].([]interface{})[0].(map[string]interface{})
		responseUrl := newReq.(map[string]interface{})["response_url"].(string)

		if action["value"] == "delete" {
			id, _ := primitive.ObjectIDFromHex(action["action_id"].(string))
			_, err := database.Db.Collection("projects").DeleteOne(context.TODO(), bson.M{"_id": id})
			if err != nil {
				panic(err)
			}
			fmt.Println("delete success!")

			newMessage := map[string]interface{}{
				"replace_original": "true",
				"text":             "Successfully deleted!",
			}

			helpers.NewPostRequest(responseUrl, newMessage)
		}

	}
}
