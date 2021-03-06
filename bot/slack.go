/*
Package containing methods for connecting to slack api, reading messages,
and responding.

Code forked from `https://github.com/rapidloop/mybot`
*/

package cfslackbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
	"os"

	"golang.org/x/net/websocket"
)

var counter uint64

// These two structures represent the response of the Slack API rtm.start.
// Only some fields are included. The rest are ignored by json.Unmarshal.
type responseRtmStart struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	Url   string       `json:"url"`
	Self  responseSelf `json:"self"`
}
type responseSelf struct {
	Id string `json:"id"`
}

// These are the messages read off and written into the websocket. Since this
// struct serves as both read and write, we include the "Id" field which is
// required only for writing.
type Message struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

// slackStart does a rtm.start, and returns a websocket URL and user ID. The
// websocket URL can be used to initiate an RTM session.
func start(token string) (wsurl, id string, err error) {
	url := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", token)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("API request failed with code %d", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}
	var respObj responseRtmStart
	err = json.Unmarshal(body, &respObj)
	if err != nil {
		return
	}

	if !respObj.Ok {
		err = fmt.Errorf("Slack error: %s", respObj.Error)
		return
	}

	wsurl = respObj.Url
	id = respObj.Self.Id
	return
}

// Starts a websocket-based Real Time API session and return the websocket
// and the ID of the (bot-)user whom the token belongs to.
func NewSlackConnection() (*websocket.Conn, string) {
	// Check if the keys exist
	slack_key := os.Getenv("SLACK_KEY")
	if slack_key == "" {
		fmt.Fprintf(os.Stderr, "SLACK_KEY missing from env")
		os.Exit(1)
	}
	wsurl, id, err := start(slack_key)
	if err != nil {
		log.Fatal(err)
	}

	ws, err := websocket.Dial(wsurl, "", "https://api.slack.com/")
	if err != nil {
		log.Fatal(err)
	}
	return ws, id
}

// Method for getting messages from the websocket connection
func (bot *Bot) GetMessage() (m Message, err error) {
	err = websocket.JSON.Receive(bot.Connection, &m)
	return
}

// Methods for posting messages back into the slack
func (bot *Bot) PostMessage(m Message) error {
	m.Id = atomic.AddUint64(&counter, 1)
	return websocket.JSON.Send(bot.Connection, m)
}
