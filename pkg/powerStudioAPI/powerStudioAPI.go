package powerstudioapi

import (
	"fmt"
	"ps-go-client/internal/business/sys/httpRequest"
)

// PowerStudio methods power studio API.
type PowerStudio struct {
	Request httprequest.HTTPRequest
	Host    string
}

// NewPowerStudio creates a new PowerStudioAPI interface.
func NewPowerStudio(host string) PowerStudio {
	return PowerStudio{
		Request: httprequest.NewHTTPRequest(),
		Host:    host,
	}
}

// PowerStudioAPI contains methods power studio API.
type PowerStudioAPI interface{}

// CallAPI .
func CallAPI() {
	r := httprequest.NewHTTPRequest()

	res, code, err := r.NewRequest("GET", "http://10.200.200.116/services/user/devices.xml", nil, nil)

	fmt.Println(res)
	fmt.Println(string(res))
	fmt.Println(code)
	fmt.Println(err)
}
