package gerard

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

// ConnectWebsocket : returns a Websocket connection that can be written/read
func ConnectWebsocket(slackURL string) *websocket.Conn {
	origin := "http://localhost/"
	headers := http.Header(map[string][]string{"origin": []string{origin}})
	conn, resp, err := websocket.DefaultDialer.Dial(slackURL, headers)
	if resp.StatusCode != http.StatusSwitchingProtocols {
		log.Fatalf("response from slack not ok: %d", resp.StatusCode)
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

func ReadMessage(ws *websocket.Conn) []byte {
	_, msg, err := ws.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}
	return msg
}
