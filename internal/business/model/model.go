package model

import "encoding/xml"

// Devices identify the main XML as a response to the device list request.
type Devices struct {
	XMLName xml.Name `xml:"devices"`
	Text    string   `xml:",chardata"`
	ID      []string `xml:"id"`
}

// DevicesInfo identify  the main XML  as  a  response  to  the  device  information request.
type DevicesInfo struct {
	XMLName xml.Name `xml:"devices"`
	Text    string   `xml:",chardata"`
	Device  []struct {
		Text            string   `xml:",chardata"`
		ID              string   `xml:"id"`
		Type            string   `xml:"type"`
		TypeDescription string   `xml:"typeDescription"`
		Var             []string `xml:"var"`
		SerialNumber    string   `xml:"serialNumber"`
		Modules         struct {
			Text   string `xml:",chardata"`
			Module struct {
				Text  string `xml:",chardata"`
				StNum string `xml:"stNum"`
				Model string `xml:"model"`
			} `xml:"module"`
		} `xml:"modules"`
	} `xml:"device"`
}
