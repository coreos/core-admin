package types

import (
	"encoding/xml"
)

type App struct {
	XMLName xml.Name `xml:"app"`
	Id      string   `xml:"id,attr"`
	Version string   `xml:"version,attr"`
	Track   string   `xml:"track,attr"`
}

type Package struct {
	XMLName   xml.Name `xml:"package"`
	Name      string   `xml:"name,attr"`      // Package filename
	Size      string   `xml:"size,attr"`      // Size of the file (in bytes)
	Path      string   `xml:"path,attr"`      // Path from the root to the file
	Sha1Sum   string   `xml:"sha1sum,attr"`   // SHA-1 hash of the file
	Sha256Sum string   `xml:"sha256sum,attr"` // Sha-256 hash of the file (extension)
	Required  bool     `xml:"required,attr"`
}
