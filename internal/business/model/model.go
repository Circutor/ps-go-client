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

// VarInfo identify  the main XML  as  a  response  to  the  request  for  information about variables.
type VarInfo struct {
	XMLName xml.Name `xml:"varInfo"`
	Text    string   `xml:",chardata"`
	Var     []struct {
		Text         string `xml:",chardata"`
		ID           string `xml:"id"`
		IDEx         string `xml:"idEx"`
		Title        string `xml:"title"`
		HasValue     string `xml:"hasValue"`
		HasLogger    string `xml:"hasLogger"`
		HasForced    string `xml:"hasForced"`
		SampleMode   string `xml:"sampleMode"`
		MeasureUnits string `xml:"measureUnits"`
		UnitsFactor  string `xml:"unitsFactor"`
		Decimals     string `xml:"decimals"`
		VarType      string `xml:"varType"`
		ValueInfo    struct {
			Text     string `xml:",chardata"`
			CtrlType string `xml:"ctrlType"`
			Type     string `xml:"type"`
		} `xml:"valueInfo"`
	} `xml:"var"`
}

// Values identify  the main XML  as  a  response  to  the  request  for  value about variables.
type Values struct {
	XMLName  xml.Name `xml:"values"`
	Text     string   `xml:",chardata"`
	Variable []struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		Value     string `xml:"value"`
		TextValue string `xml:"textValue"`
	} `xml:"variable"`
}

// RecordGroup identify  the main XML  as  a  response  to  the  request  for  record about variables.
type RecordGroup struct {
	XMLName xml.Name `xml:"recordGroup"`
	Text    string   `xml:",chardata"`
	Period  string   `xml:"period"`
	Record  []struct {
		Text     string `xml:",chardata"`
		DateTime string `xml:"dateTime"`
		Field    struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Value string `xml:"value"`
		} `xml:"field"`
		FieldComplex struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Value string `xml:"value"`
			Flags string `xml:"flags"`
		} `xml:"fieldComplex"`
		FieldARM struct {
			Text    string `xml:",chardata"`
			ID      string `xml:"id"`
			Element []struct {
				Text     string `xml:",chardata"`
				Harmonic string `xml:"harmonic"`
				Value    string `xml:"value"`
			} `xml:"element"`
		} `xml:"fieldARM"`
		FieldFO struct {
			Text    string `xml:",chardata"`
			ID      string `xml:"id"`
			Element []struct {
				Text  string `xml:",chardata"`
				Msec  string `xml:"msec"`
				Value string `xml:"value"`
			} `xml:"element"`
		} `xml:"fieldFO"`
		FieldEVQ struct {
			Text             string `xml:",chardata"`
			ID               string `xml:"id"`
			Value            string `xml:"value"`
			Phase            string `xml:"phase"`
			Duration         string `xml:"duration"`
			AverageValue     string `xml:"averageValue"`
			PreviousValue    string `xml:"previousValue"`
			EventType        string `xml:"eventType"`
			EndForced        string `xml:"endForced"`
			SemicycleVoltage []struct {
				Text  string `xml:",chardata"`
				Date  string `xml:"date"`
				Value string `xml:"value"`
			} `xml:"semicycleVoltage"`
		} `xml:"fieldEVQ"`
	} `xml:"record"`
}
