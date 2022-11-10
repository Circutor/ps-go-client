package powerstudioapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/circutor/ps-go-client/internal/business/sys/data"
	httpRequest "github.com/circutor/ps-go-client/internal/business/sys/httpRequest"
	"github.com/circutor/ps-go-client/internal/business/sys/powerStudio"
	"github.com/circutor/ps-go-client/pkg/errors"
	models "github.com/circutor/ps-go-client/pkg/models"
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
	res, err := ps.requestPs(powerstudio.HTTTP+ps.Host+powerstudio.URIAllDevices, nil)
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsAllDevices: %w", err)
	}

	body, err := data.BodyDecode(res, &models.Devices{})
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsAllDevices: %w", err)
	}

	return body.(*models.Devices), nil
}

// PsDeviceInfo get a devices information from power studio.
func (ps *PowerStudio) PsDeviceInfo(ids []string) (*models.DevicesInfo, error) {
	parameters := powerstudio.ParseParameters(ids, nil)

	res, err := ps.requestPs(powerstudio.HTTTP+ps.Host+powerstudio.URIDevicesInfo, parameters)
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsDeviceInfo: %w", err)
	}

	body, err := data.BodyDecode(res, &models.DevicesInfo{})
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsDeviceInfo: %w", err)
	}

	return body.(*models.DevicesInfo), nil
}

// PsDevicesSelectionInfo get a devices selection information from power studio.
func (ps *PowerStudio) PsDevicesSelectionInfo() (*models.DevicesSelectionInfo, error) {
	res, err := ps.requestPs(powerstudio.HTTTP+ps.Host+powerstudio.URIDevicesSelectionInfo, nil)
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsDevicesSelectionInfo: %w", err)
	}

	body, err := data.BodyDecode(res, &models.DevicesSelectionInfo{})
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsDevicesSelectionInfo: %w", err)
	}

	return body.(*models.DevicesSelectionInfo), nil
}

// PsVarInfo get an information variables from power studio.
//
// If parameter content `ids` return all variables from device.
//
// If parameter content `vars` return variables from the device it belongs to.
func (ps *PowerStudio) PsVarInfo(ids, vars []string) (*models.VarInfo, error) {
	parameters := powerstudio.ParseParameters(ids, vars)

	res, err := ps.requestPs(powerstudio.HTTTP+ps.Host+powerstudio.URIVarInfo, parameters)
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsVarInfo: %w", err)
	}

	body, err := data.BodyDecode(res, &models.VarInfo{})
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsVarInfo: %w", err)
	}

	return body.(*models.VarInfo), nil
}

// PsVarValue get a value variables from power studio.
//
// If parameter content `ids` return all values of variables from device.
//
// If parameter content `vars` return  value of variables from the device it belongs to.
func (ps *PowerStudio) PsVarValue(ids, vars []string) (*models.Values, error) {
	parameters := powerstudio.ParseParameters(ids, vars)

	res, err := ps.requestPs(powerstudio.HTTTP+ps.Host+powerstudio.URIVarValue, parameters)
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsVarValue: %w", err)
	}

	body, err := data.BodyDecode(res, &models.Values{})
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsVarValue: %w", err)
	}

	return body.(*models.Values), nil
}

// PsRecords get a records values from power studio.
func (ps *PowerStudio) PsRecords(begin, end time.Time, period int, vars []string) (*models.RecordGroup, error) {
	parameters := powerstudio.ParseParameters(nil, vars)

	if begin.IsZero() || end.IsZero() {
		return nil, fmt.Errorf("powerstudioapi.PsRecords: %w", errors.ErrPowerStudioParameters)
	}

	parameters = append(parameters, map[string]interface{}{"begin": powerstudio.ParseDateToPsFormat(begin)})
	parameters = append(parameters, map[string]interface{}{"end": powerstudio.ParseDateToPsFormat(end)})

	if period > 0 {
		parameters = append(parameters, map[string]interface{}{"period": period})
	}

	res, err := ps.requestPs(powerstudio.HTTTP+ps.Host+powerstudio.URIRecord, parameters)
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsRecords: %w", err)
	}

	body, err := data.BodyDecode(res, &models.RecordGroup{})
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.PsRecords: %w", err)
	}

	return body.(*models.RecordGroup), nil
}

// requestPs call methods power studio.
func (ps *PowerStudio) requestPs(uri string, query []map[string]interface{}) ([]byte, error) {
	resBody, statusCode, err := ps.Request.NewRequest("GET", uri, nil, query)
	if err != nil {
		return nil, fmt.Errorf("powerstudioapi.requestPs: %w", err)
	}

	if statusCode != http.StatusOK {
		if statusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("powerstudioapi.requestPs: %w", errors.ErrUnauthorizedPowerStudioAPI)
		}

		return nil, fmt.Errorf("powerstudioapi.requestPs: %w", errors.ErrPowerStudioAPI)
	}

	return resBody, nil
}
