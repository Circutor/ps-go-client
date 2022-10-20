package powerstudioapi

import (
	"net/http"

	"github.com/circutor/ps-go-client/internal/business/model"
	"github.com/circutor/ps-go-client/internal/business/sys/data"
	"github.com/circutor/ps-go-client/internal/business/sys/errors"
	httpRequest "github.com/circutor/ps-go-client/internal/business/sys/httpRequest"
	"github.com/circutor/ps-go-client/internal/business/sys/powerStudio"
)

// PowerStudio methods power studio API.
type PowerStudio struct {
	Request httpRequest.Request
	Host    string
}

// NewPowerStudio creates a new PowerStudioAPI interface.
func NewPowerStudio(host string) PowerStudio {
	request := httpRequest.NewHTTPRequest()

	return PowerStudio{
		Request: &request,
		Host:    host,
	}
}

// PowerStudioAPI contains methods power studio API.
type PowerStudioAPI interface {
	PsAllDevices() (*model.Devices, error)
	PsDeviceInfo(ids []string) (*model.DevicesInfo, error)
	PsVarInfo(ids, vars []string) (*model.VarInfo, error)
	PsVarValue(ids, vars []string) (*model.Values, error)
	PsRecords(begin, end string, period int, vars []string) (*model.RecordGroup, error)
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
func (ps *PowerStudio) PsDeviceInfo(ids []string) (*model.DevicesInfo, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIDevicesInfo

	parameters := powerstudio.ParseParameters(ids, nil)

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

// PsVarInfo get an information variables from power studio.
//
// If parameter content `ids` return all variables from device.
//
// If parameter content `vars` return variables from the device it belongs to.
func (ps *PowerStudio) PsVarInfo(ids, vars []string) (*model.VarInfo, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIVarInfo

	parameters := powerstudio.ParseParameters(ids, vars)

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, parameters)
	if err != nil {
		return &model.VarInfo{}, err
	}

	if statusCode != http.StatusOK {
		return &model.VarInfo{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &model.VarInfo{})
	if err != nil {
		return &model.VarInfo{}, err
	}

	return body.(*model.VarInfo), nil
}

// PsVarValue get a value variables from power studio.
//
// If parameter content `ids` return all values of variables from device.
//
// If parameter content `vars` return  value of variables from the device it belongs to.
func (ps *PowerStudio) PsVarValue(ids, vars []string) (*model.Values, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIVarValue

	parameters := powerstudio.ParseParameters(ids, vars)

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, parameters)
	if err != nil {
		return &model.Values{}, err
	}

	if statusCode != http.StatusOK {
		return &model.Values{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &model.Values{})
	if err != nil {
		return &model.Values{}, err
	}

	return body.(*model.Values), nil
}

// PsRecords get a records values from power studio.
func (ps *PowerStudio) PsRecords(begin, end string, period int, vars []string) (*model.RecordGroup, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIRecord

	parameters := powerstudio.ParseParameters(nil, vars)

	if begin == "" || end == "" {
		return &model.RecordGroup{}, errors.ErrPowerStudioParameters
	}

	parameters = append(parameters, map[string]interface{}{"begin": begin})
	parameters = append(parameters, map[string]interface{}{"end": end})

	if period > 0 {
		parameters = append(parameters, map[string]interface{}{"period": period})
	}

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, parameters)
	if err != nil {
		return &model.RecordGroup{}, err
	}

	if statusCode != http.StatusOK {
		return &model.RecordGroup{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &model.RecordGroup{})
	if err != nil {
		return &model.RecordGroup{}, err
	}

	return body.(*model.RecordGroup), nil
}
