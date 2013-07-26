package main
import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"io/ioutil"
)

var cmdPing = &Command{
	UsageLine: "ping",
	Short: "Test certs with a ping to c10n",
}

func init() {
	cmdPing.Run = runPing
}

func runPing(cmd *Command, args []string) {
	client := &http.Client{
		Transport: tlsTransport,
	}
	resp, err := client.PostForm("https://api.core-os.net/v1/c10n/group", url.Values{"c10n_url": {"test.net"}, "ip_list": {"test"}})
	if err != nil {
		fmt.Printf("Error: Could not ping c10n service: %v", err.Error())
		os.Exit(-1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error: Error reading response body: %v", err.Error())
		os.Exit(-1)
	}
	fmt.Printf("%v\n%v", resp, string(body))
}
