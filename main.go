package gerard

import (
	"encoding/json"
	"fmt"
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
	origin := "http://localhost/"
	url := GetWssURL()
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}

// GetWssURL : returns a wss URL for Slack
func GetWssURL() string {
	loginURL := getLoginURL()
	reader := getLoginRequestReader(loginURL)
	LoginJSON := getLoginJSONFromReader(reader)
	closeLoginRequestReader(reader)
	return LoginJSON.URL
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
	m := new(HTTPSlackResponse)
	err := json.NewDecoder(reader).Decode(m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func closeLoginRequestReader(reader io.ReadCloser) {
	defer reader.Close()
}
