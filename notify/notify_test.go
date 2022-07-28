package notify

import (
	"regexp"
	"testing"
)

func TestSendingWithOutToken(t *testing.T) {

	token_env := "LINE_NOTIFY_TOKEN"
	want := regexp.MustCompile(`\b` + token_env + `\b`)
	opt := SendOpts{
		Token:   "",
		Prefix:  "PREFIX",
		Message: "Test",
		Verbose: true,
	}

	msg := Send(opt)

	if !want.MatchString(msg) {
		t.Fatalf(`Should tell user to set %#q instead of %v`, want, msg)
	}
}
