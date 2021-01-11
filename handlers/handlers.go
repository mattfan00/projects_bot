package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bot/helpers"
	"bot/models"
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

	// var newReq map[string]interface{}

	var newReq models.SlackRequest

	if err := json.Unmarshal(body, &newReq); err != nil {
		panic(err)
	}

	/*
		// fmt.Println(newReq["token"])
		b, err := json.MarshalIndent(newReq, "", "  ")
		fmt.Println(string(b))
	*/

	if newReq.Challenge != "" {
		fmt.Printf("this is a challenge")
		// fmt.Fprintf(w, "%s", newReq.Challenge)
	} else {
		fmt.Printf("not a challenge")
		fmt.Printf("%s", newReq.Event.Text)
		helpers.PostMessage(newReq.Event.Channel, "hey dude")
	}
}
