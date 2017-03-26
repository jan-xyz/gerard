package main

import (
	"log"

	"golang.org/x/net/websocket"
)

func main() {
	Connect()
}

// SlackData : Holds the knowledge about the current session on Slack
var SlackData *Data

// Connect : connects to a websocket
func Connect() {
	wssurl := GetWssURL()
	for _, user := range SlackData.Users {
		log.Printf("User: %s (%s) is %s", user.Name, user.ID, user.Presence)
	}
	websocketConnection := ConnectWebsocket(wssurl)
	for {
		msg, n := readMessage(websocketConnection)
		log.Printf("Received: %s", string(msg[:n]))
	}
}
