package main

import (
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

func main() {
	token := "xoxb-1392098059300-1385963107650-2tIYynqnoisg8oiJDHTNQ8HF"
	url := fmt.Sprintf("https://slack.com/api/conversations.list?token=%s", token)

	resp, err := http.Get(url)
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

	fmt.Println(conversations.Channels[0])

}
