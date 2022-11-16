package powerstudioapi_test

import (
	"errors"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/circutor/ps-go-client/internal/business/sys/httpRequest/mocks"
	psErrors "github.com/circutor/ps-go-client/pkg/errors"
	psAPI "github.com/circutor/ps-go-client/pkg/powerStudioAPI"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	methodMock               = "NewRequest"
	fileAllDevices           = "../../internal/business/sampleEntities/allDevices.xml"
	fileDevicesInfoParam0    = "../../internal/business/sampleEntities/deviceInfoParam0.xml"
	fileDevicesInfoParam1    = "../../internal/business/sampleEntities/deviceInfoParam1.xml"
	fileDevicesInfoParam2    = "../../internal/business/sampleEntities/deviceInfoParam2.xml"
	fileDevicesSelectionInfo = "../../internal/business/sampleEntities/devicesSelectionInfo.xml"
	fileVarInfoParam0        = "../../internal/business/sampleEntities/varInfoParam0.xml"
	fileVarInfoParamID1      = "../../internal/business/sampleEntities/varInfoParamId1.xml"
	fileVarInfoParamID2      = "../../internal/business/sampleEntities/varInfoParamId2.xml"
	fileVarInfoParamVar1     = "../../internal/business/sampleEntities/varInfoParamVar1.xml"
	fileVarInfoParamVar2     = "../../internal/business/sampleEntities/varInfoParamVar2.xml"
	fileVarValueParam0       = "../../internal/business/sampleEntities/varValueParam0.xml"
	fileVarValueParamID1     = "../../internal/business/sampleEntities/varValueParamId1.xml"
	fileVarValueParamID2     = "../../internal/business/sampleEntities/varValueParamId2.xml"
	fileVarValueParamVar1    = "../../internal/business/sampleEntities/varValueParamVar1.xml"
	fileVarValueParamVar2    = "../../internal/business/sampleEntities/varValueParamVar2.xml"
	fileRecordsParamVar0     = "../../internal/business/sampleEntities/recordsParamVar0.xml"
	fileRecordsParamVar1     = "../../internal/business/sampleEntities/recordsParamVar1.xml"
	fileRecordsParamVar2     = "../../internal/business/sampleEntities/recordsParamVar2.xml"
)

func TestAllDevices(t *testing.T) {
	t.Parallel()

	t.Logf("Given the need to call AllDevices method API.")
	{
		uri := "http://localhost/services/user/devices.xml"

		t.Logf("\tWhen a correct api call.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileAllDevices)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			devices, err := ps.PsAllDevices()

			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(devices.ID), 0)
		}

		t.Logf("\tWhen it fails because data unmarchal error.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock
			devices, err := ps.PsAllDevices()

			assert.Error(t, err)
			assert.Equal(t, true, errors.Is(err, io.EOF))
			assert.Nil(t, devices)
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			uri := "http://10.10.10.10/services/user/devices.xml"

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusUnauthorized, nil)

			ps.Request = mock
			devices, err := ps.PsAllDevices()

			assert.Nil(t, devices)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrUnauthorizedPowerStudioAPI))

			mock = new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock
			devices, err = ps.PsAllDevices()

			assert.Nil(t, devices)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrPowerStudioAPI))
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")
			devices, err := ps.PsAllDevices()

			assert.Nil(t, devices)
			assert.Error(t, err)
		}
	}
}

func TestDeviceInfo(t *testing.T) {
	t.Parallel()

	t.Logf("Given the need to call DeviceInfo method API.")
	{
		uri := "http://localhost/services/user/deviceInfo.xml"

		t.Logf("\tWhen a correct api call witch 0 parameters.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileDevicesInfoParam0)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			devices, err := ps.PsDeviceInfo(nil)

			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(devices.Device), 0)
		}

		t.Logf("\tWhen a correct api call witch 1 parameters.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileDevicesInfoParam1)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"id": "cvm-e3-mini"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			devices, err := ps.PsDeviceInfo([]string{"cvm-e3-mini"})

			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(devices.Device), 1)
		}

		t.Logf("\tWhen a correct api call witch 2 parameters.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileDevicesInfoParam2)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"id": "cvm-e3-mini"},
				{"id": "TCPRS1-firmware"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			devices, err := ps.PsDeviceInfo([]string{"cvm-e3-mini", "TCPRS1-firmware"})

			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(devices.Device), 2)
		}

		t.Logf("\tWhen it fails because data unmarchal error.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock
			devices, err := ps.PsDeviceInfo(nil)

			assert.Nil(t, devices)
			assert.Equal(t, true, errors.Is(err, io.EOF))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			uri := "http://10.10.10.10/services/user/deviceInfo.xml"

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock
			devices, err := ps.PsDeviceInfo(nil)

			assert.Nil(t, devices)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrPowerStudioAPI))

			mock = new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusUnauthorized, nil)

			ps.Request = mock
			devices, err = ps.PsDeviceInfo(nil)

			assert.Nil(t, devices)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrUnauthorizedPowerStudioAPI))
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			devices, err := ps.PsDeviceInfo(nil)

			assert.Nil(t, devices)
			assert.Error(t, err)
		}
	}
}

func TestDevicesSelectionInfo(t *testing.T) {
	t.Parallel()

	t.Logf("Given the need to call DevicesSelectionInfo method API.")
	{
		uri := "http://localhost/services/devices/devicesSelectionInfo.xml"

		t.Logf("\tWhen a correct api call.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileDevicesSelectionInfo)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			devicesSelectionInfo, err := ps.PsDevicesSelectionInfo()

			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(devicesSelectionInfo.Devices.Device), 0)
		}

		t.Logf("\tWhen it fails because data unmarchal error.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock
			devicesSelectionInfo, err := ps.PsDevicesSelectionInfo()

			assert.Nil(t, devicesSelectionInfo)
			assert.Equal(t, true, errors.Is(err, io.EOF))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			uri := "http://10.10.10.10/services/devices/devicesSelectionInfo.xml"

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock
			devicesSelectionInfo, err := ps.PsDevicesSelectionInfo()

			assert.Nil(t, devicesSelectionInfo)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrPowerStudioAPI))

			mock = new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusUnauthorized, nil)

			ps.Request = mock
			devicesSelectionInfo, err = ps.PsDevicesSelectionInfo()

			assert.Nil(t, devicesSelectionInfo)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrUnauthorizedPowerStudioAPI))
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			devicesSelectionInfo, err := ps.PsDevicesSelectionInfo()

			assert.Nil(t, devicesSelectionInfo)
			assert.Error(t, err)
		}
	}
}

func TestVarInfo(t *testing.T) {
	t.Parallel()

	t.Logf("Given the need to call VarInfo method API.")
	{
		uri := "http://localhost/services/user/varInfo.xml"

		t.Logf("\tWhen a correct api call witch 0 parameters.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarInfoParam0)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarInfo(nil, nil)

			assert.Nil(t, err)
			assert.Equal(t, 0, len(vars.Var))
		}

		t.Logf("\tWhen a correct api call witch 1 parameter type id")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarInfoParamID1)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"id": "cvm-e3-mini"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarInfo([]string{"cvm-e3-mini"}, nil)

			assert.Nil(t, err)
			assert.Equal(t, 657, len(vars.Var))
		}

		t.Logf("\tWhen a correct api call witch 2 parameter type id")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarInfoParamID2)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"id": "cvm-e3-mini"},
				{"id": "TCPRS1-firmware"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarInfo([]string{"cvm-e3-mini", "TCPRS1-firmware"}, nil)

			assert.Nil(t, err)
			assert.Equal(t, 2082, len(vars.Var))
		}

		t.Logf("\tWhen a correct api call witch 1 parameter type var")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarInfoParamVar1)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"var": "cvm-e3-mini.AE1"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarInfo(nil, []string{"cvm-e3-mini.AE1"})

			assert.Nil(t, err)
			assert.Equal(t, 1, len(vars.Var))
		}

		t.Logf("\tWhen a correct api call witch 2 parameter type var")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarInfoParamVar2)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"var": "cvm-e3-mini.AE1"},
				{"var": "cvm-e3-mini.AE1B"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarInfo(nil, []string{"cvm-e3-mini.AE1", "cvm-e3-mini.AE1B"})

			assert.Nil(t, err)
			assert.Equal(t, 2, len(vars.Var))
		}

		t.Logf("\tWhen it fails because data unmarchal error.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarInfo(nil, nil)

			assert.Nil(t, vars)
			assert.Equal(t, true, errors.Is(err, io.EOF))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			uri := "http://10.10.10.10/services/user/varInfo.xml"

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock

			vars, err := ps.PsVarInfo(nil, nil)

			assert.Nil(t, vars)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrPowerStudioAPI))

			mock = new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusUnauthorized, nil)

			ps.Request = mock

			vars, err = ps.PsVarInfo(nil, nil)

			assert.Nil(t, vars)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrUnauthorizedPowerStudioAPI))
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			vars, err := ps.PsVarInfo(nil, nil)

			assert.Nil(t, vars)
			assert.Error(t, err)
		}
	}
}

func TestVarValue(t *testing.T) {
	t.Parallel()

	t.Logf("Given the need to call VarValue method API.")
	{
		uri := "http://localhost/services/user/values.xml"

		t.Logf("\tWhen a correct api call witch 0 parameters.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarValueParam0)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarValue(nil, nil)

			assert.Nil(t, err)
			assert.Equal(t, 0, len(vars.Variable))
		}

		t.Logf("\tWhen a correct api call witch 1 parameter type id")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarValueParamID1)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"id": "cvm-e3-mini"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarValue([]string{"cvm-e3-mini"}, nil)

			assert.Nil(t, err)
			assert.Equal(t, 657, len(vars.Variable))
		}

		t.Logf("\tWhen a correct api call witch 2 parameter type id")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarValueParamID2)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"id": "cvm-e3-mini"},
				{"id": "TCPRS1-firmware"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarValue([]string{"cvm-e3-mini", "TCPRS1-firmware"}, nil)

			assert.Nil(t, err)
			assert.Equal(t, 2082, len(vars.Variable))
		}

		t.Logf("\tWhen a correct api call witch 1 parameter type var")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarValueParamVar1)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"var": "cvm-e3-mini.AE1"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarValue(nil, []string{"cvm-e3-mini.AE1"})

			assert.Nil(t, err)
			assert.Equal(t, 1, len(vars.Variable))
		}

		t.Logf("\tWhen a correct api call witch 2 parameter type var")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileVarValueParamVar2)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameters := []map[string]interface{}{
				{"var": "cvm-e3-mini.AE1"},
				{"var": "cvm-e3-mini.AE1B"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameters).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarValue(nil, []string{"cvm-e3-mini.AE1", "cvm-e3-mini.AE1B"})

			assert.Nil(t, err)
			assert.Equal(t, 2, len(vars.Variable))
		}

		t.Logf("\tWhen it fails because data unmarchal error.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock
			vars, err := ps.PsVarValue(nil, nil)

			assert.Nil(t, vars)
			assert.Equal(t, true, errors.Is(err, io.EOF))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			uri := "http://10.10.10.10/services/user/values.xml"

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock
			vars, err := ps.PsVarValue(nil, nil)

			assert.Nil(t, vars)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrPowerStudioAPI))

			mock = new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusUnauthorized, nil)

			ps.Request = mock
			vars, err = ps.PsVarValue(nil, nil)

			assert.Nil(t, vars)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrUnauthorizedPowerStudioAPI))
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			vars, err := ps.PsVarValue(nil, nil)

			assert.Nil(t, vars)
			assert.Error(t, err)
		}
	}
}

func TestRecords(t *testing.T) {
	t.Parallel()

	t.Logf("Given the need to call Records method API.")
	{
		uri := "http://localhost/services/user/records.xml"

		t.Logf("\tWhen a correct api call witch 0 parameters type var.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileRecordsParamVar0)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameter := []map[string]interface{}{
				{"begin": "18102022000000"},
				{"end": "18102022000000"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock

			timeBegin := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)
			timeEnd := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)

			records, err := ps.PsRecords(timeBegin, timeEnd, 0, nil)

			assert.Nil(t, err)
			assert.Equal(t, 0, len(records.Record))
		}

		t.Logf("\tWhen a correct api call witch 1 parameters type var.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileRecordsParamVar1)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameter := []map[string]interface{}{
				{"var": "CVM-C5.VMX23"},
				{"begin": "18102022000000"},
				{"end": "18102022000000"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock

			timeBegin := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)
			timeEnd := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)

			records, err := ps.PsRecords(timeBegin, timeEnd, 0, []string{"CVM-C5.VMX23"})

			assert.Nil(t, err)
			assert.Equal(t, 40, len(records.Record))
		}

		t.Logf("\tWhen a correct api call witch 2 parameters type var.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			xmlFile, err := os.Open(fileRecordsParamVar2)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameter := []map[string]interface{}{
				{"var": "CVM-C5.VMX23"},
				{"var": "CVM-C5.AET1"},
				{"begin": "18102022000000"},
				{"end": "18102022000000"},
				{"period": 10},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock

			timeBegin := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)
			timeEnd := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)

			records, err := ps.PsRecords(timeBegin, timeEnd, 10, []string{"CVM-C5.VMX23", "CVM-C5.AET1"})

			assert.Nil(t, err)
			assert.Equal(t, 40, len(records.Record))
		}

		t.Logf("\tWhen it fails because data unmarchal error.")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			parameter := []map[string]interface{}{
				{"begin": "18102022000000"},
				{"end": "18102022000000"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock

			timeBegin := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)
			timeEnd := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)

			records, err := ps.PsRecords(timeBegin, timeEnd, 0, nil)

			assert.Nil(t, records)
			assert.Equal(t, true, errors.Is(err, io.EOF))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			uri := "http://10.10.10.10/services/user/records.xml"

			parameter := []map[string]interface{}{
				{"begin": "18102022000000"},
				{"end": "18102022000000"},
			}

			timeBegin := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)
			timeEnd := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock
			records, err := ps.PsRecords(timeBegin, timeEnd, 0, nil)

			assert.Nil(t, records)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrPowerStudioAPI))

			mock = new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return([]byte(""), http.StatusUnauthorized, nil)

			ps.Request = mock
			records, err = ps.PsRecords(timeBegin, timeEnd, 0, nil)

			assert.Nil(t, records)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrUnauthorizedPowerStudioAPI))
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := psAPI.NewPowerStudio("10.10.10.10", "", "")

			timeBegin := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)
			timeEnd := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)

			records, err := ps.PsRecords(timeBegin, timeEnd, 0, nil)

			assert.Nil(t, records)
			assert.Error(t, err)
		}

		t.Logf("\tWhen it fails because there is parameters .")
		{
			ps := psAPI.NewPowerStudio("localhost", "", "")

			timeBegin := time.Date(2022, 10, 18, 0, 0, 0, 0, time.UTC)
			timeEnd := time.Date(2022, 10, 19, 0, 0, 0, 0, time.UTC)

			records, err := ps.PsRecords(timeBegin, time.Time{}, 0, nil)

			assert.Nil(t, records)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrPowerStudioParameters))

			records, err = ps.PsRecords(time.Time{}, timeEnd, 0, nil)

			assert.Nil(t, records)
			assert.Equal(t, true, errors.Is(err, psErrors.ErrPowerStudioParameters))
		}
	}
}

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	t.Logf("Given the need to call load config method API.")
	{
		t.Logf("\tWhen a correct api call load config.")
		{
			const (
				host     = "0.0.0.0"
				user     = "user"
				password = "password"
			)

			ps := psAPI.NewPowerStudio("", "", "")

			ps.PsLoadConfig(host, user, password)

			assert.Equal(t, host, ps.Host)
			assert.Equal(t, user, ps.Username)
			assert.Equal(t, password, ps.Password)
		}
	}
}
