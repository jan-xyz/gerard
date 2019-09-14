package main

import (
	"log"
)

func main() {
	connect()
}

func connect() {
	slackData := startRTM()
	for _, user := range slackData.Users {
		log.Printf("User: %s (%s) is %s", user.Name, user.ID, user.Presence)
	}
	slackConnection := connectWebsocket(slackData.URL)
	for {
		msg := readMessage(slackConnection)
		ParseMessage(msg, slackData)
	}
}
