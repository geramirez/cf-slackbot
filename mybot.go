/*

mybot - Illustrative Slack bot in Go

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ramirezg/cf-slackbot/cloudfoundry"
	"github.com/ramirezg/cf-slackbot/slack"
)

func main() {
	slackbot_key := os.Getenv("SLACKBOT_KEY")
	if slackbot_key == "" {
		fmt.Fprintf(os.Stderr, "SLACKBOT_KEY missing from env")
		os.Exit(1)
	}

	// start a websocket-based Real Time API session
	bot := slack.Connect(slackbot_key)
	fmt.Println("bot ready, ^C exits")

	for {
		// get each incoming message
		m, err := bot.GetMessage()
		if err != nil {
			log.Fatal(err)
		}
		// process each message
		if cloudfoundry.Process(&m) {
			bot.PostMessage(m)
		}
	}
}
