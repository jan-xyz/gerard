package gerard_core

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// Data : structure that should fit the JSON getting back
type Data struct {
	URL      string    `json:"url"`
	OK       bool      `json:"ok"`
	Users    []User    `json:"users"`
	Team     team      `json:"team"`
	Channels []channel `json:"channels"`
	Groups   []group   `json:"groups"`
	Error    string    `json:"error"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Presence string `json:"presence"`
}

type team struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// StartRTM : Login to Slack via rtm.start
func StartRTM() *Data {
	loginURL := GetLoginURL()
	reader := getLoginRequestReader(loginURL)
	data := getLoginJSONFromReader(reader)
	closeLoginRequestReader(reader)
	return data
}

func GetLoginURL() string {
	apikey := os.Getenv("ROLLMOPS_SLACK_API_KEY")
	if apikey == "" {
		log.Fatal("no API key found.")
	}
	return "http://slack.com/api/rtm.start?token=" + apikey
}

func getLoginRequestReader(loginURL string) io.ReadCloser {
	resp, err := http.Get(loginURL)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body
}

func getLoginJSONFromReader(reader io.ReadCloser) *Data {
	loginStructBytes := new(bytes.Buffer)
	loginStructBytes.ReadFrom(reader)
	s := loginStructBytes.String()
	loginStruct := new(Data)
	err := json.Unmarshal([]byte(s), loginStruct)
	if err != nil {
		log.Fatal(err)
	}
	return loginStruct
}

func closeLoginRequestReader(reader io.ReadCloser) {
	defer reader.Close()
}
