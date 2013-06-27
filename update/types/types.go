// Package types is where exportable API structures go. This is so we can share
// the XML marshalling and unmarshalling with the services.
package types

import (
	"encoding/xml"
)

type Version struct {
	XMLName xml.Name `xml:"version"`
	App     *App     `xml:"app"`
	Package *Package `xml:"package"`
}

type App struct {
	XMLName xml.Name `xml:"app"`
	Id      string   `xml:"id,attr"`
	Version string   `xml:"version,attr"`
	Track   string   `xml:"track,attr"`
}

type Package struct {
	XMLName              xml.Name `xml:"package"`
	Name                 string   `xml:"name,attr"`      // Package filename
	Size                 string   `xml:"size,attr"`      // Size of the file (in bytes)
	Path                 string   `xml:"path,attr"`      // Path from the root to the file
	Sha1Sum              string   `xml:"sha1sum,attr"`   // SHA-1 hash of the file
	Sha256Sum            string   `xml:"sha256sum,attr"` // Sha-256 hash of the file (extension)
	Required             bool     `xml:"required,attr"`
	MetadataSignatureRsa string   `xml:"MetadataSignatureRsa,attr,omitempty" json:"metadata_signature_rsa"`
	MetadataSize         string   `xml:"MetadataSize,attr,omitempty" json:"metadata_size"`
}
