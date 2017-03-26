package main

import (
	"log"

	"golang.org/x/net/websocket"
)

// ConnectWebsocket : returns a Websocket connection that can be written/read
func ConnectWebsocket() *websocket.Conn {
	origin := "http://localhost/"
	ws, err := websocket.Dial(SlackData.URL, "", origin)
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
