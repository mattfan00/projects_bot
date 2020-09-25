package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Channel struct {
	Id string `json:"id"`
}

type ConversationsList struct {
	Ok       bool      `json:"ok"`
	Channels []Channel `json:"channels"`
}

type Message struct {
	Token   string `json:"token"`
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

type SlackRequest struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
	Event     struct {
		Channel string `json:"channel"`
		Text    string `json:"text"`
	} `json:"event"`
}

var token = "xoxb-1392098059300-1385963107650-2tIYynqnoisg8oiJDHTNQ8HF"

func GetChannelId() string {
	getUrl := fmt.Sprintf("https://slack.com/api/conversations.list?token=%s", token)

	resp, err := http.Get(getUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var conversations ConversationsList
	err = json.Unmarshal(body, &conversations)
	if err != nil {
		panic(err)
	}

	return conversations.Channels[0].Id

}

func PostMessage(channelId string, message string) {
	postUrl := "https://slack.com/api/chat.postMessage"
	newMessage := Message{token, message, channelId}
	reqBody, err := json.Marshal(newMessage)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	r, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(reqBody))
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

	var newReq SlackRequest

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
		fmt.Fprintf(w, "%s", newReq.Challenge)
	} else {
		fmt.Printf("not a challenge")
		fmt.Printf("%s", newReq.Event.Text)
		PostMessage(newReq.Event.Channel, "hey dude")
	}

}

func main() {
	// PostMessage("no more messages")

	router := mux.NewRouter().StrictSlash(true)
	router.Path("/slack").Methods("POST").HandlerFunc(slackHandler)
	router.Path("/").Methods("GET").HandlerFunc(indexHandler)

	http.ListenAndServe(":8080", router)

}
