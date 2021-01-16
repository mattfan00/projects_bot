package helpers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"bot/models"

	"fmt"
	"github.com/joho/godotenv"
)

func GetChannelId() string {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	getUrl := fmt.Sprintf("https://slack.com/api/conversations.list?token=%s", os.Getenv("SLACK_TOKEN"))

	resp, err := http.Get(getUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var conversations models.ConversationsList
	err = json.Unmarshal(body, &conversations)
	if err != nil {
		panic(err)
	}

	return conversations.Channels[0].Id
}

func PostMessage(channelId string, message string) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	token := os.Getenv("SLACK_TOKEN")
	postUrl := "https://slack.com/api/chat.postMessage"
	newMessage := models.Message{token, message, channelId}
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

	//fmt.Println(result)
}
