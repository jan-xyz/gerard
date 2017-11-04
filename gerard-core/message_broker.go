package gerard_core

import (
	"encoding/json"
	"log"
)

type slackJson struct {
	Type      string `json:"type"`
	Channel   string `json:"channel"`
	User      string `json:"user"`
	Text      string `json:"text"`
	Timestamp string `json:"ts"`
	Presence  string `json:"presence"`
	URL       string `json:"url"`
}

// ParseMessage : Parses message to decide usage
func ParseMessage(msg []byte, data *Data) {
	slackUnmarshal := new(slackJson)
	err := json.Unmarshal(msg, slackUnmarshal)
	if err != nil {
		log.Fatal(err)
	}
	if slackUnmarshal.Type == "presence_change" {
		for _, user := range data.Users {
			if user.ID == slackUnmarshal.User {
				user.Presence = slackUnmarshal.Presence
				log.Printf("%s is now %s", user.Name, user.Presence)
			}
		}
	} else if slackUnmarshal.Type == "hello" {
		log.Print("Successfully logged in.")
	} else if slackUnmarshal.Type == "reconnect_url" {
		data.URL = slackUnmarshal.URL
		log.Printf("New reconnection URL set: %s", data.URL)

	} else {
		log.Printf("Received: %s", string(msg))
	}
}