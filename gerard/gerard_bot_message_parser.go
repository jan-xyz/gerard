package gerard

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

type reconnectURL struct {
	URL string `json:"url"`
}

// RTMSlackObject : structure for identifying Slack JSON messages
type RTMSlackObject struct {
	Type string `json:"type"`
}

// ParseMessage : Parses message to decide usage
func ParseMessage(msg []byte, data *Data) {
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
		for _, user := range data.Users {
			if user.ID == presenceChange.User {
				user.Presence = presenceChange.Presence
				log.Printf("%s is now %s", user.Name, user.Presence)
			}
		}
	} else if proto.Type == "hello" {
		log.Print("Successfully logged in.")
	} else if proto.Type == "reconnect_url" {
		urlContainer := new(reconnectURL)
		err = json.Unmarshal(msg, urlContainer)
		if err != nil {
			log.Fatal(err)
		}
		data.URL = urlContainer.URL
		log.Printf("New reconnection URL set: %s", data.URL)

	} else {
		log.Printf("Received: %s", string(msg))
	}
}
