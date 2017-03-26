package main

import (
	"encoding/json"
	"log"
)

type slackMessage struct {
	Channel   string `json:"channel"`
	User      string `json:"user"`
	Text      string `json:"text"`
	Timestamp string `json:"ts"`
}

type slackPresenceChange struct {
	User     string `json:"user"`
	Presence string `json:"presence"`
}

// RTMSlackObject : structure for identifying Slack JSON messages
type RTMSlackObject struct {
	Type string `json:"type"`
}

// ParseMessage : Parses message to decide usage
func ParseMessage(msg []byte) {
	proto := new(RTMSlackObject)
	err := json.Unmarshal(msg, proto)
	if err != nil {
		log.Fatal(err)
	}
	if proto.Type == "presence_change" {
		presenceChange := new(slackPresenceChange)
		err = json.Unmarshal(msg, presenceChange)
		if err != nil {
			log.Fatal(err)
		}
		for _, user := range SlackData.Users {
			if user.ID == presenceChange.User {
				user.Presence = presenceChange.Presence
				log.Printf("%s is now %s", user.Name, user.Presence)
			}
		}
	} else if proto.Type == "hello" {
		log.Print("Successfully logged in.")
	} else {
		log.Printf("Received: %s", string(msg))
	}
}
