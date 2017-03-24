package gerard

import (
	"os"
	"testing"
)

//func TestConnect(t *testing.T) {
//	Connect()
//}

func TestLoginToSlack(t *testing.T) {
	GetWssURL()
}

func TestGetLoginURL(t *testing.T) {
	testAPIToken := "test"
	expectedURL := "http://slack.com/api/rtm.start?token=" + testAPIToken
	os.Setenv("ROLLMOPS_SLACK_API_KEY", testAPIToken)
	url := getLoginURL()
	if url != expectedURL {
		t.Fail()
	}
}
