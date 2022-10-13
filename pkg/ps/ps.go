package ps

import (
	"fmt"
	"ps-go-client/internal/business/sys/httpRequest"
)

// CallAPI .
func CallAPI() {
	r := httprequest.NewHTTPRequest()

	res, code, err := r.NewRequest("GET", "http://10.200.200.116/services/user/devices.xml", nil, nil)

	fmt.Println(res)
	fmt.Println(string(res))
	fmt.Println(code)
	fmt.Println(err)
}
