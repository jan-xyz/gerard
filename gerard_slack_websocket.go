package main

import (
	"log"
	"net/http"

	//	"golang.org/x/net/websocket"
	"github.com/gorilla/websocket"
)

// ConnectWebsocket : returns a Websocket connection that can be written/read
func ConnectWebsocket(slackURL string) *websocket.Conn {
	origin := "http://localhost/"
	//conn, err := net.Dial("tcp", parseUrlToDial(slackUrl))
	headers := http.Header(map[string][]string{"origin": []string{origin}})
	conn, resp, err := websocket.DefaultDialer.Dial(slackURL, headers)
	if resp.StatusCode != http.StatusOK {
		log.Fatal("response from slack not ok")
	}
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func parseURLToDial(url string) string {
	return ""
}

func sendMessage(message string, ws *websocket.Conn) {
	if err := ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		log.Fatal(err)
	}
}

func readMessage(ws *websocket.Conn) []byte {
	msgType, msg, err := ws.ReadMessage()
	log.Printf("received message type: %d, content: %s", msgType, string(msg))
	if err != nil {
		log.Fatal(err)
	}
	return msg
}
