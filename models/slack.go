package models

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

type SlashMessage struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}
