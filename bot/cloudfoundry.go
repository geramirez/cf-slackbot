/* This package allows the bot to interface with the cloudfoundry api */

package cfslackbot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type TokenRes struct {
	// Basic token struct that CF url returns
	AccessToken  string `json:"access_token"`
	Expires      int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	JTI          string `json:"jti"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

type Token struct {
	// Modified token struct with a time stamp to check if it's expired
	TokenRes
	CreatedTime int
}

type APIResponse struct {
	// Basic API struct used in CF api responses
	TotalResults int    `json:"total_results"`
	TotalPages   int    `json:"total_pages"`
	PrevUrl      string `json:"prev_url"`
	NextUrl      string `json:"next_url"`
}

type AppResponse struct {
	APIResponse
	Apps string `json:"resource"`
}

func configTokenRequest() *http.Request {
	// Configure a new token request
	token_url := fmt.Sprintf("https://uaa.%s/oauth/token", os.Getenv("API_URL"))
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", os.Getenv("CF_USERNAME"))
	data.Set("password", os.Getenv("CF_PASSWORD"))
	req, _ := http.NewRequest("POST", token_url, bytes.NewBufferString(data.Encode()))
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("authorization", "Basic Y2Y6")
	return req
}

func NewCloudFoundryToken() *Token {
	// Check if env variables are in place
	if os.Getenv("CF_USERNAME") == "" {
		fmt.Fprintf(os.Stderr, "CF_USERNAME missing from env")
		os.Exit(1)
	}
	if os.Getenv("CF_PASSWORD") == "" {
		fmt.Fprintf(os.Stderr, "CF_PASSWORD missing from env")
		os.Exit(1)
	}
	// Initalize a new token
	var token Token
	token.getToken()
	return &token
}

func (token *Token) getToken() {
	// Get a new token
	req := configTokenRequest()
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	if json.Unmarshal(body, token) != nil {
		fmt.Println("Error")
	}
	token.CreatedTime = int(time.Now().Unix())
}

func (token *Token) updateToken() {
	// Check if token has expired, if so updates the token
	if int(time.Now().Unix())-token.CreatedTime > token.Expires {
		// replace with a token refresher
		token.getToken()
	}
}

func (token *Token) request(req_url string) *http.Response {
	// Makes a request to the specific url with the token
	token.updateToken()
	req, _ := http.NewRequest("GET", req_url, nil)
	req.Header.Set("authorization", fmt.Sprintf("bearer %s", token.AccessToken))
	client := &http.Client{}
	res, _ := client.Do(req)
	return res
}

func (bot *Bot) getAppStatus(app string) string {
	// Get the status of a particular app
	req_url := fmt.Sprintf("https://api.%s%s", os.Getenv("API_URL"), "/v2/apps")
	res := bot.CFToken.request(req_url)

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	var response AppResponse
	if json.Unmarshal(body, &response) != nil {
		fmt.Println("Error")
	}
	return "trying to get data"
}
