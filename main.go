package main

import (
	"log"
)

func main() {
	Connect()
}

// Connect : connects to a websocket
func Connect() {
	slackData := StartRTM()
	for _, user := range slackData.Users {
		log.Printf("User: %s (%s) is %s", user.Name, user.ID, user.Presence)
	}
	slackConnection := ConnectWebsocket(slackData.URL)
	for {
		msg := ReadMessage(slackConnection)
		ParseMessage(msg, slackData)
	}
}
