package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bot/slack"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from the index")
}

func slackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reached the slack route")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// var newReq map[string]interface{}

	var newReq slack.SlackRequest

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
		slack.PostMessage(newReq.Event.Channel, "hey dude")
	}
}
