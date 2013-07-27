package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/coreos/core-admin/certs"
	"net/http"
	"net/url"
	"os"
)

// Default URL for the update API, overriden with CORE_UPDATE_URL
// TODO BP: use api.core-os.net instead but need to override CA certs
var updateURL = url.URL{Scheme: "https", Host: "api.core-os.net"}
var tlsTransport = &http.Transport{
	DisableCompression: true,
}

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

	coreosInetPemBlock, _ := pem.Decode([]byte(certs.CoreOS_Internet_Authority_pem))
	if coreosInetPemBlock == nil {
		panic("Error: No PEM data found in CoreOS_Internet_Auth")
	}

	coreosNetPemBlock, _ := pem.Decode([]byte(certs.CoreOS_Internet_Authority_pem))
	if coreosNetPemBlock == nil {
		panic("Error: No PEM data found in CoreOS_Network_Auth")
	}

	coreosInetAuthCert, err := x509.ParseCertificate(coreosInetPemBlock.Bytes)
	if err != nil {
		panic(err.Error())
	}
	coreosNetAuthCert, err := x509.ParseCertificate(coreosNetPemBlock.Bytes)
	if err != nil {
		panic(err.Error())
	}
	pool.AddCert(coreosNetAuthCert)
	pool.AddCert(coreosInetAuthCert)
	tlsTransport.TLSClientConfig = &tls.Config{RootCAs: pool}
}
