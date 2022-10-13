package model

import "encoding/xml"

// Devices identify the main XML as a response to he device list request.
type Devices struct {
	XMLName xml.Name `xml:"devices"`
	Text    string   `xml:",chardata"`
	ID      []string `xml:"id"`
}
