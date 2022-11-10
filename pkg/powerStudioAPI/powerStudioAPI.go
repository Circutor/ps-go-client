package powerstudioapi

import (
	"net/http"
	"time"

	models "github.com/circutor/ps-go-client/pkg/models"

	"github.com/circutor/ps-go-client/internal/business/sys/data"
	"github.com/circutor/ps-go-client/internal/business/sys/errors"
	httpRequest "github.com/circutor/ps-go-client/internal/business/sys/httpRequest"
	"github.com/circutor/ps-go-client/internal/business/sys/powerStudio"
)

// PowerStudio methods power studio API.
type PowerStudio struct {
	Request  httpRequest.Request
	Host     string
	Username string
	Password string
}

// NewPowerStudio creates a new PowerStudioAPI interface.
func NewPowerStudio(host, username, password string) PowerStudio {
	request := httpRequest.NewHTTPRequest(username, password)

	return PowerStudio{
		Request:  &request,
		Host:     host,
		Username: username,
		Password: password,
	}
}

//go:generate mockery --name PowerStudioAPI --structname PowerStudioAPIMock --filename PowerStudioAPIMock.go

// PowerStudioAPI contains methods power studio API.
type PowerStudioAPI interface {
	PsAllDevices() (*models.Devices, error)
	PsDeviceInfo(ids []string) (*models.DevicesInfo, error)
	PsDevicesSelectionInfo() (*models.DevicesSelectionInfo, error)
	PsVarInfo(ids, vars []string) (*models.VarInfo, error)
	PsVarValue(ids, vars []string) (*models.Values, error)
	PsRecords(begin, end time.Time, period int, vars []string) (*models.RecordGroup, error)
}

// PsAllDevices get all devices from power studio.
func (ps *PowerStudio) PsAllDevices() (*models.Devices, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIAllDevices

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, nil)
	if err != nil {
		return &models.Devices{}, err
	}

	if statusCode != http.StatusOK {
		return &models.Devices{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &models.Devices{})
	if err != nil {
		return &models.Devices{}, err
	}

	return body.(*models.Devices), nil
}

// PsDeviceInfo get a devices information from power studio.
func (ps *PowerStudio) PsDeviceInfo(ids []string) (*models.DevicesInfo, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIDevicesInfo

	parameters := powerstudio.ParseParameters(ids, nil)

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, parameters)
	if err != nil {
		return &models.DevicesInfo{}, err
	}

	if statusCode != http.StatusOK {
		return &models.DevicesInfo{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &models.DevicesInfo{})
	if err != nil {
		return &models.DevicesInfo{}, err
	}

	return body.(*models.DevicesInfo), nil
}

// PsDevicesSelectionInfo get a devices selection information from power studio.
func (ps *PowerStudio) PsDevicesSelectionInfo() (*models.DevicesSelectionInfo, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIDevicesSelectionInfo

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, nil)
	if err != nil {
		return &models.DevicesSelectionInfo{}, err
	}

	if statusCode != http.StatusOK {
		return &models.DevicesSelectionInfo{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &models.DevicesSelectionInfo{})
	if err != nil {
		return &models.DevicesSelectionInfo{}, err
	}

	return body.(*models.DevicesSelectionInfo), nil
}

// PsVarInfo get an information variables from power studio.
//
// If parameter content `ids` return all variables from device.
//
// If parameter content `vars` return variables from the device it belongs to.
func (ps *PowerStudio) PsVarInfo(ids, vars []string) (*models.VarInfo, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIVarInfo

	parameters := powerstudio.ParseParameters(ids, vars)

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, parameters)
	if err != nil {
		return &models.VarInfo{}, err
	}

	if statusCode != http.StatusOK {
		return &models.VarInfo{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &models.VarInfo{})
	if err != nil {
		return &models.VarInfo{}, err
	}

	return body.(*models.VarInfo), nil
}

// PsVarValue get a value variables from power studio.
//
// If parameter content `ids` return all values of variables from device.
//
// If parameter content `vars` return  value of variables from the device it belongs to.
func (ps *PowerStudio) PsVarValue(ids, vars []string) (*models.Values, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIVarValue

	parameters := powerstudio.ParseParameters(ids, vars)

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, parameters)
	if err != nil {
		return &models.Values{}, err
	}

	if statusCode != http.StatusOK {
		return &models.Values{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &models.Values{})
	if err != nil {
		return &models.Values{}, err
	}

	return body.(*models.Values), nil
}

// PsRecords get a records values from power studio.
func (ps *PowerStudio) PsRecords(begin, end time.Time, period int, vars []string) (*models.RecordGroup, error) {
	uri := powerstudio.HTTTP + ps.Host + powerstudio.URIRecord

	parameters := powerstudio.ParseParameters(nil, vars)

	if begin.IsZero() || end.IsZero() {
		return &models.RecordGroup{}, errors.ErrPowerStudioParameters
	}

	parameters = append(parameters, map[string]interface{}{"begin": powerstudio.ParseDateToPsFormat(begin)})
	parameters = append(parameters, map[string]interface{}{"end": powerstudio.ParseDateToPsFormat(end)})

	if period > 0 {
		parameters = append(parameters, map[string]interface{}{"period": period})
	}

	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, parameters)
	if err != nil {
		return &models.RecordGroup{}, err
	}

	if statusCode != http.StatusOK {
		return &models.RecordGroup{}, errors.ErrPowerStudioAPI
	}

	body, err := data.BodyDecode(resBody, &models.RecordGroup{})
	if err != nil {
		return &models.RecordGroup{}, err
	}

	return body.(*models.RecordGroup), nil
}
