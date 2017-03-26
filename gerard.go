package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	Connect()
}

// HTTPSlackResponse : structure that should fit the JSON getting back
type HTTPSlackResponse struct {
	URL string `json:"url"`
}

// Connect : connects to a websocket
func Connect() {
	wssurl := GetWssURL()
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
