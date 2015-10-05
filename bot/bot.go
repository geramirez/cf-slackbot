package cfslackbot

import (
	"fmt"
	"os"

	"golang.org/x/net/websocket"
)

// Struct that holds the websocket connection, the bot user id,
// and the Cloud Foundry token information.
type Bot struct {
	ID         string
	Connection *websocket.Conn
}

// Function for initializing the bot
func InitBot() *Bot {
	slack_key := os.Getenv("SLACK_KEY")
	if slack_key == "" {
		fmt.Fprintf(os.Stderr, "SLACK_KEY missing from env")
		os.Exit(1)
	}
	// start a websocket-based Real Time API session
	ws, id := Connect(slack_key)
	return &Bot{id, ws}
}
