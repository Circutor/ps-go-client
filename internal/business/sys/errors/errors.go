package errors

import "github.com/pkg/errors"

var (
	ErrHTTPRequestCreate   = "error in create request"
	ErrHTTPRequestSend     = "error in send request"
	ErrHTTPRequestReadBody = "error in read response body"

	ErrBodyDecode = "error in data response body"

	ErrPowerStudioAPI        = errors.New("Error in PowerStudio API")
	ErrPowerStudioParameters = errors.New("Error in PowerStudio parameters")
)
