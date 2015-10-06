/*
Script for starting the bot server
*/

package main

import (
	"fmt"
	"log"

	"github.com/ramirezg/cf-slackbot/bot"
)

func main() {
	// Start the bot server
	bot := cfslackbot.InitBot()
	fmt.Println("Bot Ready")
	for {
		// Get each incoming message
		message, err := bot.GetMessage()
		if err != nil {
			log.Fatal(err)
		}
		// Process each message and respond if
		// trigger is set
		if bot.Process(&message) {
			bot.PostMessage(message)
		}
	}
}
