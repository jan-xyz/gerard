package gerard

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

func main() {
	Connect()
}

// Connect : connects to a websocket
func Connect() {
	origin := "http://localhost/"
	url := LoginToSlack()
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, newerr := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(newerr)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}

// LoginToSlack : returns a wss URL for Slack
func LoginToSlack() string {
	return "ws://echo.websocket.org/"
}
