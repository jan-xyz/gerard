package main

import (
	"log"
	"github.com/jan-xyz/gerard/gerard"
	"github.com/gorilla/websocket"
)

func main() {
	Connect()
}

// SlackConnection : Can be used to read/write to slack
var slackConnection *websocket.Conn

// Connect : connects to a websocket
func Connect() {
	slackData := gerard.StartRTM()
	for _, user := range slackData.Users {
		log.Printf("User: %s (%s) is %s", user.Name, user.ID, user.Presence)
	}
	slackConnection = gerard.ConnectWebsocket(slackData.URL)
	for {
		msg := gerard.ReadMessage(slackConnection)
		gerard.ParseMessage(msg, slackData)
	}
}
