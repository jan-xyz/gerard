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

// SlackConnection : Can be used to read/write to slack
var SlackConnection *websocket.Conn

// Connect : connects to a websocket
func Connect() {
	StartRTM()
	for _, user := range SlackData.Users {
		log.Printf("User: %s (%s) is %s", user.Name, user.ID, user.Presence)
	}
	SlackConnection = ConnectWebsocket()
	for {
		msg, n := readMessage(SlackConnection)
		ParseMessage(msg[:n])
	}
}
