package data

import (
	"encoding/xml"
	"fmt"
)

// BodyDecode data body of response object to data object type.
func BodyDecode(body []byte, data interface{}) (interface{}, error) {
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("data.bodyDecode: %w", err)
	}

	return data, nil
}
