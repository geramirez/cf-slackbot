package cfslackbot

import (
	"fmt"
	"strings"
)

func (bot *Bot) routeMessage(message *Message) bool {
	// Split the string
	words := strings.Fields(message.Text)
	// Check commands that have 3 fields
	if len(words) == 3 {
		// Search for the command
		switch words[1] {
		case "app":
			message.Text = bot.getAppStatus(words[2])
		default:
			message.Text = "Sorry I couldn't understand what you wanted"
		}
		return true
	}
	return false
}

// Method for processing an incomming messages from slack
func (bot *Bot) Process(message *Message) bool {
	fmt.Println(message)
	if message.Type == "message" && strings.HasPrefix(message.Text, "cf ") {
		return bot.routeMessage(message)
	}
	return false
}
