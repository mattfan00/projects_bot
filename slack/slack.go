package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fmt"
)

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
