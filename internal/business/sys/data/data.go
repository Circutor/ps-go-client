package data

import (
	"encoding/xml"
	"fmt"

	"github.com/circutor/ps-go-client/internal/business/sys/errors"
)

// BodyDecode data body of response object to data object type.
func BodyDecode(body []byte, data interface{}) (interface{}, error) {
	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("%s : %w", errors.ErrBodyDecode, err)
	}

	return data, nil
}
