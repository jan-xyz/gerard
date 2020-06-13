package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoginURL(t *testing.T) {
	testAPIToken := "test"
	expectedURL := "http://slack.com/api/rtm.start?token=" + testAPIToken
	os.Setenv("ROLLMOPS_SLACK_API_KEY", testAPIToken)
	url := getLoginURL()
	assert.Equal(t, expectedURL, url)
}
