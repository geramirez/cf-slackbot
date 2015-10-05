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

	bot := cfslackbot.InitBot()
	fmt.Println("bot ready, ^C exits")

	for {
		// get each incoming message
		m, err := bot.GetMessage()
		if err != nil {
			log.Fatal(err)
		}
		// process each message
		if bot.Process(&m) {
			bot.PostMessage(m)
		}
	}
}
