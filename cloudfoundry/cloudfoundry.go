/* This package allows the bot to interface with the cloudfoundry api */

package cloudfoundry

import (
	"fmt"
	"strings"

	"github.com/ramirezg/cf-slackbot/slack"
)

func Process(m *slack.Message) bool {
	fmt.Println(m)
	if m.Type == "message" && strings.HasPrefix(m.Text, "cf ") {
		m.Text = "Now executing `" + m.Text + "`"
		return true
	}
	return false
}
