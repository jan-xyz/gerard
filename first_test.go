package gerard

import "testing"

func TestConnect(t *testing.T) {
	Connect()
}

func TestLoginToSlack(t *testing.T) {
	test := LoginToSlack()
	if test != "ws://echo.websocket.org/" {
		t.Fail()
	}
}
