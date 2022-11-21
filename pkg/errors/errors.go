package errors

import "github.com/pkg/errors"

var (
	ErrPowerStudioRequest         = errors.New("Error in PowerStudio Request")
	ErrPowerStudioAPI             = errors.New("Error in PowerStudio API")
	ErrUnauthorizedPowerStudioAPI = errors.New("Error Unauthorized in PowerStudio API")
	ErrPowerStudioParameters      = errors.New("Error in PowerStudio parameters")
)
