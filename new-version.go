package main

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/coreos/core-admin/update/types"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

var cmdNewVersion = &Command{
	UsageLine: "new-version -k [key] -a [app-id] -v [version] -t [track] -p [url path] [filename]",
	Short:     "update the version database for a given file",
	Long: `
Takes a file path and some meta data and update the information used in the datastore.
	`,
}

var versionD = cmdNewVersion.Flag.Bool("d", false, "dry run, print out the xml payload")
var versionK = cmdNewVersion.Flag.String("k", "", "api key for the admin user")

var versionA = cmdNewVersion.Flag.String("a", "", "application id")
var versionV = cmdNewVersion.Flag.String("v", "", "version ")
var versionT = cmdNewVersion.Flag.String("t", "", "track")
var versionP = cmdNewVersion.Flag.String("p", "", "url path")
var versionM = cmdNewVersion.Flag.String("m", "", "metadata filename")
var versionC = cmdNewVersion.Flag.Bool("c", false, "canary version")

func init() {
	cmdNewVersion.Run = runNewVersion
}

func readMetadata(filename string, pkg *types.Package) {
	metadata, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(metadata, pkg)
	if err != nil {
		panic(err)
	}
}

func calculateHashes(filename string, pkg *types.Package) {
	var (
		writers []io.Writer
		hashes  []hash.Hash
	)

	push := func(h hash.Hash) {
		writers = append(writers, h)
		hashes = append(hashes, h)
	}

	push(sha256.New())
	push(sha1.New())

	in, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	io.Copy(io.MultiWriter(writers...), in)

	formatHash := func(hash hash.Hash) string {
		return base64.StdEncoding.EncodeToString(hash.Sum(nil))
	}

	pkg.Sha256Sum = formatHash(hashes[0])
	pkg.Sha1Sum = formatHash(hashes[1])
}

func newVersionRequestBody(args []string) []byte {
	appId := *versionA
	version := *versionV
	track := *versionT
	path := *versionP
	metadata := *versionM
	canary := *versionC

	if appId == "" || version == "" || track == "" || path == "" || metadata == "" {
		panic("one of the required fields was not present\n")
	}

	if len(args) != 1 {
		panic("update file name not provided\n")
	}

	file := args[0]
	fileBase := filepath.Base(file)
	fi, err := os.Stat(file)
	if err != nil {
		panic(err)
	}

	fileSize := strconv.FormatInt(fi.Size(), 10)

	app := types.App{Id: appId, Version: version, Track: track, IsCanary: canary}
	pkg := types.Package{Name: fileBase, Size: fileSize, Path: path}
	ver := types.Version{App: &app, Package: &pkg}
	calculateHashes(file, ver.Package)
	readMetadata(metadata, ver.Package)

	raw, err := xml.MarshalIndent(ver, "", " ")
	if err != nil {
		panic(err)
	}

	body := []byte(xml.Header)
	return append(body, raw...)
}

func runNewVersion(cmd *Command, args []string) {
	dryRun := *versionD
	key := *versionK

	if dryRun == false && key == "" {
		panic("key or dry-run required")
	}

	body := newVersionRequestBody(args)
	adminURL, _ := url.Parse(updateURL.String())
	adminURL.Path = "/admin/v1/version"

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
