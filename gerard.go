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
	websocket := connectWebsocket(wssurl)
	for {
		msg, n := readMessage(websocket)
		log.Printf("Received: %s", string(msg[:n]))
	}
}

func connectWebsocket(url string) *websocket.Conn {
	origin := "http://localhost/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	return ws
}

func sendMessage(message string, ws *websocket.Conn) {
	if _, err := ws.Write([]byte(message)); err != nil {
		log.Fatal(err)
	}
}

func readMessage(ws *websocket.Conn) ([]byte, int) {
	var msg = make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	return msg, n
}
