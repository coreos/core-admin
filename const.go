package main

import (
	"fmt"
	"net/url"
	"os"
)

// Default URL for the update API, overriden with CORE_UPDATE_URL
// TODO BP: use api.core-os.net instead but need to override CA certs
var updateURL = url.URL{Scheme: "https", Host: "core-api.appspot.com"}

func init() {
	// Setup the updateURL if the environment variable is set
	envApi := os.Getenv("CORE_UPDATE_URL")
	if envApi != "" {
		envUrl, err := url.Parse(envApi)
		if err != nil {
			fmt.Printf("Error: couldn't parse CORE_UPDATE_URL: %s", err.Error())
			os.Exit(-1)
		}
		updateURL = *envUrl
	}
}
