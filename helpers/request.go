package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func NewPostRequest(url string, body interface{}) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	token := os.Getenv("SLACK_TOKEN")

	reqBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	var result map[string]interface{}

	json.NewDecoder(res.Body).Decode(&result)

}
