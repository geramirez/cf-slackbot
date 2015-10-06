package cfslackbot

import (
	"fmt"

	"golang.org/x/net/websocket"
)

// Struct that holds the websocket connection, the bot user id,
// and the Cloud Foundry token information.
type Bot struct {
	ID         string
	Connection *websocket.Conn
	CFToken *Token
}

// Function for initializing the bot
func InitBot() *Bot {
	// Start a websocket-based Real Time API session
	ws, id := NewSlackConnection()
	fmt.Println("Slack Connection Ready")
	// Initalize Cloud Foundry token
	token := NewCloudFoundryToken()
	fmt.Println("CloudFoundry Token Ready")
	return &Bot{id, ws, token}
}
