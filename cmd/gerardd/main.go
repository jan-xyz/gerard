package main

import (
	"log"
	"github.com/jan-xyz/gerard/gerard-core"
	"github.com/gorilla/websocket"
)

func main() {
	Connect()
}

// SlackConnection : Can be used to read/write to slack
var slackConnection *websocket.Conn

// Connect : connects to a websocket
func Connect() {
	slackData := gerard_core.StartRTM()
	for _, user := range slackData.Users {
		log.Printf("User: %s (%s) is %s", user.Name, user.ID, user.Presence)
	}
	log.Print(slackData.Channels)
	log.Print(slackData.Users)
	slackConnection = gerard_core.ConnectWebsocket(slackData.URL)
	for {
		msg := gerard_core.ReadMessage(slackConnection)
		gerard_core.ParseMessage(msg, slackData)
	}
}
