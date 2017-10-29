package tests

import (
	"os"
	"testing"
	"github.com/jan-xyz/gerard/gerard"
)

func TestGetLoginURL(t *testing.T) {
	testAPIToken := "test"
	expectedURL := "http://slack.com/api/rtm.start?token=" + testAPIToken
	os.Setenv("ROLLMOPS_SLACK_API_KEY", testAPIToken)
	url := gerard.GetLoginURL()
	if url != expectedURL {
		t.Fail()
	}
}
