package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Channel struct {
	Id string `json:"id"`
}

type ConversationsList struct {
	Ok       bool      `json:"ok"`
	Channels []Channel `json:"channels"`
}

type PostMessage struct {
	Token   string `json:"token"`
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

func main() {
	token := "xoxb-1392098059300-1385963107650-2tIYynqnoisg8oiJDHTNQ8HF"
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

	channelId := conversations.Channels[0].Id

	postUrl := "https://slack.com/api/chat.postMessage"
	newMessage := PostMessage{token, "hello from go", channelId}
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
