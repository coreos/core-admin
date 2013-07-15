package main

import (
	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/goauth2/oauth/jwt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// newDevClient Returns a new oauth2 authenticated *http.Client which works in appengine dev environments
// TODO: add if the paths are nil then look at datastore
func NewToken(pemFilePath string, clientSecretsPath string, s string) (*oauth.Token, error) {
	keybytes, err := ioutil.ReadFile(pemFilePath)
	if err != nil {
		return nil, err
	}

	secretbytes, err := ioutil.ReadFile(clientSecretsPath)
	if err != nil {
		return nil, err
	}

	var config struct {
		Web struct {
			ClientEmail string `json:"client_email"`
			ClientID    string `json:"client_id"`
			TokenURI    string `json:"token_uri"`
		}
	}

	err = json.Unmarshal(secretbytes, &config)
	if err != nil {
		return nil, err
	}

	tok := jwt.NewToken(config.Web.ClientEmail, s, keybytes)
	tok.ClaimSet.Aud = config.Web.TokenURI

	tokClient := &http.Client{}

	o, err := tok.Assert(tokClient)

	return o, err
}
