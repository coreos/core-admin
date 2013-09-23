package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/coreos/core-admin/admin/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var cmdCanary = &Command{
	UsageLine: "canary [bootid]",
	Short:     "add a canary to the database",
	Long: `
Takes a bootid and adds it to the canary machine list
	`,
}

var canaryK = cmdCanary.Flag.String("k", "", "api key for the admin user")
var canaryD = cmdCanary.Flag.Bool("d", false, "dry run, print out the xml payload")

func init() {
	cmdCanary.Run = runCanary
}

func newCanaryRequestBody(bootid string) []byte {
	cm := types.CanaryMachine{BootId: fmt.Sprintf("{%s}", bootid)}

	raw, err := xml.MarshalIndent(cm, "", " ")
	if err != nil {
		panic(err)
	}

	body := []byte(xml.Header)
	return append(body, raw...)
}

func runCanary(cmd *Command, args []string) {
	id := args[0]
	dryRun := *canaryD
	key := *canaryK

	if dryRun == false && key == "" {
		panic("key or dry-run required")
	}

	body := newCanaryRequestBody(id)
	adminURL, _ := url.Parse(updateURL.String())
	adminURL.Path = "/admin/v1/canaryMachine"

	req, _ := http.NewRequest("POST", adminURL.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "text/xml")
	req.SetBasicAuth("admin", key)

	if dryRun || *debug {
		req.Write(os.Stdout)
	}

	if dryRun {
		return
	}

	client := &http.Client{
		Transport: tlsTransport,
	}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, _ = ioutil.ReadAll(resp.Body)
	os.Stdout.Write(body)
	fmt.Printf("\n")

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("Error: bad return code %s\n", resp.Status))
	}

	return
}
