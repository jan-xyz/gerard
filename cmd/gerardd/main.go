package main

import (
	"github.com/jan-xyz/gerard/gerard-core"
	"os"
	"log"
)

func main() {
	apikey := os.Getenv("ROLLMOPS_SLACK_API_KEY")
	if apikey == "" {
		log.Fatal("no API key found.")
	}
	gerard_core.Connect(apikey)
}
