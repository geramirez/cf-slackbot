/* This package allows the bot to interface with the cloudfoundry api */

package cfslackbot

import (
	"fmt"
	"strings"
)

// Method for processing an incomming messages from slack
func (bot *Bot) Process(m *Message) bool {
	fmt.Println(m)
	if m.Type == "message" && strings.HasPrefix(m.Text, "cf ") {
		m.Text = "Now executing `" + m.Text + "`"
		return true
	}
	return false
}
