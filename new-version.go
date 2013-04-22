package reup

import (
	"fmt"
	"hash"
	"crypto/sha1"
	"crypto/sha256"
	"os"
	"io"
	"reup/web"
)

var cmdNewVersion = &Command{
	UsageLine: "new-version -k [key] -a [app-id] -v [version] -t [track] -u [url path] [filename]",
	Short:     "update the version database for a given file",
	Long: `
Takes a file path and some meta data and update the information used in the datastore.
	`,
}

func init() {
	cmdNewVersion.Run = runNewVersion
}

func calculateHashes(filename string, pkg *web.Package) (names []string, hashes []hash.Hash)  {
	var (
		writers []io.Writer
	)

	push := func(name string, h hash.Hash) {
		writers = append(writers, h)
		hashes = append(hashes, h)
		names = append(names, name)
	}

	push("sha256sum", sha256.New())
	push("sha1sum", sha1.New())

	in, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(io.MultiWriter(writers...), in)

	return names, hashes
}

func runNewVersion(cmd *Command, args []string) {
	app := web.App{}
	pkg := web.Package{}
	_, hashes := calculateHashes(args[0], &pkg)

	for i := 0; i < len(hashes); i++ {
		fmt.Printf("%v %v\n", app, pkg)
	}

	return
}
