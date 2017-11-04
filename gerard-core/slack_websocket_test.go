package gerard_core

import (
	"testing"
)

func TestGetLoginURL(t *testing.T) {
	testAPIToken := "test"
	expectedURL := "http://slack.com/api/rtm.start?token=" + testAPIToken
	url := GetLoginURL(testAPIToken)
	if url != expectedURL {
		t.Fail()
	}
}
