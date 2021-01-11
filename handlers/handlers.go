package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"bot/database"
	"bot/helpers"
	"bot/models"

	"github.com/joho/godotenv"
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

	action := newReq.(map[string]interface{})["actions"].([]interface{})[0].(map[string]interface{})
	responseUrl := newReq.(map[string]interface{})["response_url"].(string)

	if action["value"] == "delete" {
		id, _ := primitive.ObjectIDFromHex(action["action_id"].(string))
		_, err := database.Db.Collection("projects").DeleteOne(context.TODO(), bson.M{"_id": id})
		if err != nil {
			panic(err)
		}
		fmt.Println("delete success!")

		if err := godotenv.Load(); err != nil {
			panic(err)
		}
		token := os.Getenv("SLACK_TOKEN")
		newMessage := map[string]interface{}{
			"replace_original": "true",
			"text":             "Successfully deleted!",
		}
		reqBody, err := json.Marshal(newMessage)
		if err != nil {
			panic(err)
		}

		client := &http.Client{}
		r, err := http.NewRequest("POST", responseUrl, bytes.NewBuffer(reqBody))
		if err != nil {
			panic(err)
		}
		r.Header.Add("Content-Type", "application/json")
		r.Header.Add("Authorization", "Bearer "+token)

		res, err := client.Do(r)

		if err != nil {
			panic(err)
		}

		var result map[string]interface{}

		json.NewDecoder(res.Body).Decode(&result)

		fmt.Println(result)
	}
}
