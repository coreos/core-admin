package main

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"regexp"
	"code.google.com/p/gopass"
	"encoding/json"
)

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// newDevClient Returns a new oauth2 authenticated *http.Client which works in appengine dev environments
// TODO: add if the paths are nil then look at datastore
func NewToken()	(*Token, error){
	request_args := url.Values{"response_type": {"code"}, "client_id": {clientId}, "redirect_uri": {"urn:ietf:wg:oauth:2.0:oob"}, "scope": {scope}, }
	goto_url := fmt.Sprintf("%s?%s", "https://accounts.google.com/o/oauth2/auth", request_args.Encode())

	fmt.Println("Goto this url: ", goto_url)
	code, err := gopass.GetPass("Paste the code here: ")
	if err != nil {
		fmt.Printf("ERR with code %v", err)
	}
	token_args := url.Values{"code": {code}, "client_id": {clientId}, "client_secret": {clientSecret}, "redirect_uri": {"urn:ietf:wg:oauth:2.0:oob"}, "grant_type": {"authorization_code"}, }
	resp, err := http.PostForm("https://accounts.google.com/o/oauth2/token", token_args)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	token := new(Token)
	err = json.Unmarshal(body, token)
	
	return token, err
	
}

func AdminTest(tok *Token) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v?%v", "http://api.core-os.net/admin/v1/metrics", tok.AccessToken), nil)
	if err != nil {
		fmt.Println(err)
	}
	//req.Header.Set("Authorization", "Bearer "+tok.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(body))
}
