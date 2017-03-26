package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// HTTPSlackResponse : structure that should fit the JSON getting back
type HTTPSlackResponse struct {
	URL      string `json:"url"`
	OK       bool   `json:"ok"`
	Users    []user `json:"users"`
	Teams    team   `json:"team"`
	Channels []team `json:"channels"`
	Groups   []team `json:"groups"`
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
	loginStructBytes := new(bytes.Buffer)
	loginStructBytes.ReadFrom(reader)
	s := loginStructBytes.String()
	log.Print(s)
	loginStruct := new(HTTPSlackResponse)
	err := json.Unmarshal([]byte(s), loginStruct)
	// err := json.NewDecoder(reader).Decode(loginStruct)
	if err != nil {
		log.Fatal(err)
	}
	return loginStruct
}

func closeLoginRequestReader(reader io.ReadCloser) {
	defer reader.Close()
}
