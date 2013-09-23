// Package types is where exportable API structures go. This is so we can share
// the XML marshalling and unmarshalling with the services.
package types

import (
	"encoding/xml"
)

// CanaryMachine is a machine that will get a new update on a track earlier than
// everyone else. This is a mechanism for helping us gain confidence on rolling
// out an update.
type CanaryMachine struct {
	XMLName xml.Name `xml:"canarymachine" datastore:"-" json:"-"`
	BootId string `xml:"bootid,attr"`
}
