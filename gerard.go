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
	msg, n := readMessage(websocket)
	log.Printf("Received: %s", string(msg[:n]))
}

// GetWssURL : returns a wss URL for Slack
func GetWssURL() string {
	loginURL := getLoginURL()
	reader := getLoginRequestReader(loginURL)
	LoginJSON := getLoginJSONFromReader(reader)
	closeLoginRequestReader(reader)
	return LoginJSON.URL
}

func connectWebsocket(url string) *websocket.Conn {
	origin := "http://localhost/"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	return ws
}

func getLoginURL() string {
	apikey := os.Getenv("ROLLMOPS_SLACK_API_KEY")
	return "http://slack.com/api/rtm.start?token=" + apikey
}

func getLoginRequestReader(loginURL string) io.ReadCloser {
	resp, err := http.Get(loginURL)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body
}

func getLoginJSONFromReader(reader io.ReadCloser) *HTTPSlackResponse {
	loginStruct := new(HTTPSlackResponse)
	err := json.NewDecoder(reader).Decode(loginStruct)
	if err != nil {
		log.Fatal(err)
	}
	return loginStruct
}

func closeLoginRequestReader(reader io.ReadCloser) {
	defer reader.Close()
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
