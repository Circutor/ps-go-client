package powerstudioapi

import (
	"net/http"
	"ps-go-client/internal/business/model"
	"ps-go-client/internal/business/sys/data"
	"ps-go-client/internal/business/sys/errors"
	"ps-go-client/internal/business/sys/httpRequest"
	"ps-go-client/internal/business/sys/powerStudio"
)

// PowerStudio methods power studio API.
type PowerStudio struct {
	Request httprequest.Request
	Host    string
}

// NewPowerStudio creates a new PowerStudioAPI interface.
func NewPowerStudio(host string) PowerStudio {
	request := httprequest.NewHTTPRequest()

	return PowerStudio{
		Request: &request,
		Host:    host,
	}
}

// PowerStudioAPI contains methods power studio API.
type PowerStudioAPI interface {
	PsAllDevices() (*model.Devices, error)
	PsDeviceInfo(parameters []map[string]interface{}) (*model.DevicesInfo, error)
}

// PsAllDevices get all devices from power studio.
func (ps *PowerStudio) PsAllDevices() (*model.Devices, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIAllDevices

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, nil)
	if err != nil {
		return &model.Devices{}, err
	}

	if statusCode != http.StatusOK {
		return &model.Devices{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &model.Devices{})
	if err != nil {
		return &model.Devices{}, err
	}

	return body.(*model.Devices), nil
}

// PsDeviceInfo get a devices information from power studio.
func (ps *PowerStudio) PsDeviceInfo(parameters []map[string]interface{}) (*model.DevicesInfo, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIDevicesInfo

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, parameters)
	if err != nil {
		return &model.DevicesInfo{}, err
	}

	if statusCode != http.StatusOK {
		return &model.DevicesInfo{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &model.DevicesInfo{})
	if err != nil {
		return &model.DevicesInfo{}, err
	}

	return body.(*model.DevicesInfo), nil
}
