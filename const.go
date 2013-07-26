package main

import (
	"fmt"
	"net/url"
	"os"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"encoding/pem"
)

// Default URL for the update API, overriden with CORE_UPDATE_URL
// TODO BP: use api.core-os.net instead but need to override CA certs
var updateURL = url.URL{Scheme: "https", Host: "api.core-os.net"}
var tlsTransport = &http.Transport {
	DisableCompression: true,
}
var coreosInternetAuthPath = "./certs/CoreOS_Internet_Authority.pem"
var coreosNetworkAuthPath = "./certs/CoreOS_Network_Authority.pem"

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
	pool := x509.NewCertPool()
	coreosInternetAuth, err := ioutil.ReadFile(coreosInternetAuthPath)
	if err != nil {
		fmt.Printf("Error: couldn't parse Coreos_Internet_Auth: %s", err.Error())
		os.Exit(-1)
	}
	coreosNetworkAuth, err := ioutil.ReadFile(coreosNetworkAuthPath)
	if err != nil {
		fmt.Printf("Error: couldn't parse Coreos_Network_Auth: %s", err.Error())
		os.Exit(-1)
	}
	coreosInetPemBlock, _ := pem.Decode(coreosInternetAuth)
	if coreosInetPemBlock == nil {
		fmt.Printf("Error: No PEM data found in CoreOS_Internet_Auth")
		os.Exit(-1)
	}

	coreosNetPemBlock, _ := pem.Decode(coreosNetworkAuth)
	if coreosNetPemBlock == nil {
		fmt.Printf("Error: No PEM data found in CoreOS_Network_Auth")
		os.Exit(-1)
	}

	coreosInetAuthCert, err := x509.ParseCertificate(coreosInetPemBlock.Bytes)
	if err != nil {
		fmt.Printf("Error: invalid Internet Auth Cert: %s", err.Error())
		os.Exit(-1)
	}
	coreosNetAuthCert, err := x509.ParseCertificate(coreosNetPemBlock.Bytes)
	if err != nil {
		fmt.Printf("Error: invalid Network Auth Cert: %s", err.Error())
		os.Exit(-1)
	}
	pool.AddCert(coreosNetAuthCert)
	pool.AddCert(coreosInetAuthCert)
	tlsTransport.TLSClientConfig = &tls.Config{RootCAs: pool}
}
