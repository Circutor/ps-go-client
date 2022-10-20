package powerstudioapi_test

import (
	"github.com/circutor/ps-go-client/internal/business/sys/errors"
	"github.com/circutor/ps-go-client/internal/business/sys/httpRequest/mocks"
	powerStudioAPI "github.com/circutor/ps-go-client/pkg/powerStudioAPI"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	methodMock            = "NewRequest"
	fileAllDevices        = "../../internal/business/sampleEntities/allDevices.xml"
	fileDevicesInfoParam0 = "../../internal/business/sampleEntities/deviceInfoParam0.xml"
	fileDevicesInfoParam1 = "../../internal/business/sampleEntities/deviceInfoParam1.xml"
	fileDevicesInfoParam2 = "../../internal/business/sampleEntities/deviceInfoParam2.xml"
	fileVarInfoParam0     = "../../internal/business/sampleEntities/varInfoParam0.xml"
	fileVarInfoParamId1   = "../../internal/business/sampleEntities/varInfoParamId1.xml"
	fileVarInfoParamId2   = "../../internal/business/sampleEntities/varInfoParamId2.xml"
	fileVarInfoParamVar1  = "../../internal/business/sampleEntities/varInfoParamVar1.xml"
	fileVarInfoParamVar2  = "../../internal/business/sampleEntities/varInfoParamVar2.xml"
	fileVarValueParam0    = "../../internal/business/sampleEntities/varValueParam0.xml"
	fileVarValueParamId1  = "../../internal/business/sampleEntities/varValueParamId1.xml"
	fileVarValueParamId2  = "../../internal/business/sampleEntities/varValueParamId2.xml"
	fileVarValueParamVar1 = "../../internal/business/sampleEntities/varValueParamVar1.xml"
	fileVarValueParamVar2 = "../../internal/business/sampleEntities/varValueParamVar2.xml"
	fileRecordsParamVar0  = "../../internal/business/sampleEntities/recordsParamVar0.xml"
	fileRecordsParamVar1  = "../../internal/business/sampleEntities/recordsParamVar1.xml"
	fileRecordsParamVar2  = "../../internal/business/sampleEntities/recordsParamVar2.xml"
)

func TestAllDevices(t *testing.T) {
	t.Parallel()

	t.Logf("Given the need to call AllDevices method API.")
	{
		uri := "http://localhost/services/user/devices.xml"

		t.Logf("\tWhen a correct api call.")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

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
			ps := powerStudioAPI.NewPowerStudio("localhost")

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock

			devices, err := ps.PsAllDevices()

			assert.Error(t, err)
			assert.Equal(t, 0, len(devices.ID))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			uri := "http://10.10.10.10/services/user/devices.xml"

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock

			devices, err := ps.PsAllDevices()

			assert.Equal(t, 0, len(devices.ID))
			assert.Equal(t, errors.ErrPowerStudioAPI, err)
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			devices, err := ps.PsAllDevices()

			assert.Equal(t, 0, len(devices.ID))
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
			ps := powerStudioAPI.NewPowerStudio("localhost")

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
			ps := powerStudioAPI.NewPowerStudio("localhost")

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
			ps := powerStudioAPI.NewPowerStudio("localhost")

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
			ps := powerStudioAPI.NewPowerStudio("localhost")

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock

			devices, err := ps.PsDeviceInfo(nil)

			assert.Error(t, err)
			assert.Equal(t, 0, len(devices.Device))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			uri := "http://10.10.10.10/services/user/deviceInfo.xml"

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock

			devices, err := ps.PsDeviceInfo(nil)

			assert.Equal(t, 0, len(devices.Device))
			assert.Equal(t, errors.ErrPowerStudioAPI, err)
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			devices, err := ps.PsDeviceInfo(nil)

			assert.Equal(t, 0, len(devices.Device))
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
			ps := powerStudioAPI.NewPowerStudio("localhost")

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
			ps := powerStudioAPI.NewPowerStudio("localhost")

			xmlFile, err := os.Open(fileVarInfoParamId1)
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
			ps := powerStudioAPI.NewPowerStudio("localhost")

			xmlFile, err := os.Open(fileVarInfoParamId2)
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
			ps := powerStudioAPI.NewPowerStudio("localhost")

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
			ps := powerStudioAPI.NewPowerStudio("localhost")

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
			ps := powerStudioAPI.NewPowerStudio("localhost")

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock

			vars, err := ps.PsVarInfo(nil, nil)

			assert.Error(t, err)
			assert.Equal(t, 0, len(vars.Var))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			uri := "http://10.10.10.10/services/user/varInfo.xml"

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}{}).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock

			vars, err := ps.PsVarInfo(nil, nil)

			assert.Equal(t, 0, len(vars.Var))
			assert.Equal(t, errors.ErrPowerStudioAPI, err)
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			vars, err := ps.PsVarInfo(nil, nil)

			assert.Equal(t, 0, len(vars.Var))
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
			ps := powerStudioAPI.NewPowerStudio("localhost")

			xmlFile, err := os.Open(fileVarValueParam0)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock

			vars, err := ps.PsVarValue(nil)

			assert.Nil(t, err)
			assert.Equal(t, 0, len(vars.Variable))
		}

		t.Logf("\tWhen a correct api call witch 1 parameter type id")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

			xmlFile, err := os.Open(fileVarValueParamId1)
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

			vars, err := ps.PsVarValue(parameters)

			assert.Nil(t, err)
			assert.Equal(t, 657, len(vars.Variable))
		}

		t.Logf("\tWhen a correct api call witch 2 parameter type id")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

			xmlFile, err := os.Open(fileVarValueParamId2)
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

			vars, err := ps.PsVarValue(parameters)

			assert.Nil(t, err)
			assert.Equal(t, 2082, len(vars.Variable))
		}

		t.Logf("\tWhen a correct api call witch 1 parameter type var")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

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

			vars, err := ps.PsVarValue(parameters)

			assert.Nil(t, err)
			assert.Equal(t, 1, len(vars.Variable))
		}

		t.Logf("\tWhen a correct api call witch 2 parameter type var")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

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

			vars, err := ps.PsVarValue(parameters)

			assert.Nil(t, err)
			assert.Equal(t, 2, len(vars.Variable))
		}

		t.Logf("\tWhen it fails because data unmarchal error.")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock

			vars, err := ps.PsVarValue(nil)

			assert.Error(t, err)
			assert.Equal(t, 0, len(vars.Variable))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			uri := "http://10.10.10.10/services/user/values.xml"

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, []map[string]interface{}(nil)).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock

			vars, err := ps.PsVarValue(nil)

			assert.Equal(t, 0, len(vars.Variable))
			assert.Equal(t, errors.ErrPowerStudioAPI, err)
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			vars, err := ps.PsVarValue(nil)

			assert.Equal(t, 0, len(vars.Variable))
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
			ps := powerStudioAPI.NewPowerStudio("localhost")

			xmlFile, err := os.Open(fileRecordsParamVar0)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameter := []map[string]interface{}{
				{"begin": "18102022"},
				{"end": "18102022"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock

			records, err := ps.PsRecords("18102022", "18102022", 0, nil)

			assert.Nil(t, err)
			assert.Equal(t, 0, len(records.Record))
		}

		t.Logf("\tWhen a correct api call witch 1 parameters type var.")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

			xmlFile, err := os.Open(fileRecordsParamVar1)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameter := []map[string]interface{}{
				{"var": "CVM-C5.VMX23"},
				{"begin": "18102022"},
				{"end": "18102022"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock

			records, err := ps.PsRecords("18102022", "18102022", 0, []map[string]interface{}{
				{"var": "CVM-C5.VMX23"},
			})

			assert.Nil(t, err)
			assert.Equal(t, 40, len(records.Record))
		}

		t.Logf("\tWhen a correct api call witch 2 parameters type var.")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

			xmlFile, err := os.Open(fileRecordsParamVar2)
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			parameter := []map[string]interface{}{
				{"var": "CVM-C5.VMX23"},
				{"var": "CVM-C5.AET1"},
				{"begin": "18102022"},
				{"end": "18102022"},
				{"period": 10},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return(byteValue, http.StatusOK, nil)

			ps.Request = mock

			records, err := ps.PsRecords("18102022", "18102022", 10, []map[string]interface{}{
				{"var": "CVM-C5.VMX23"},
				{"var": "CVM-C5.AET1"},
			})

			assert.Nil(t, err)
			assert.Equal(t, 40, len(records.Record))
		}

		t.Logf("\tWhen it fails because data unmarchal error.")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

			parameter := []map[string]interface{}{
				{"begin": "18102022"},
				{"end": "18102022"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return([]byte(""), http.StatusOK, nil)

			ps.Request = mock

			records, err := ps.PsRecords("18102022", "18102022", 0, nil)

			assert.Error(t, err)
			assert.Equal(t, 0, len(records.Record))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			uri := "http://10.10.10.10/services/user/records.xml"

			parameter := []map[string]interface{}{
				{"begin": "18102022"},
				{"end": "18102022"},
			}

			mock := new(mocks.RequestMock)
			mock.On(methodMock, http.MethodGet, uri, nil, parameter).
				Return([]byte(""), http.StatusNotFound, nil)

			ps.Request = mock

			records, err := ps.PsRecords("18102022", "18102022", 0, nil)

			assert.Equal(t, 0, len(records.Record))
			assert.Equal(t, errors.ErrPowerStudioAPI, err)
		}

		t.Logf("\tWhen it fails because there is timeout.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			records, err := ps.PsRecords("18102022", "18102022", 0, nil)

			assert.Equal(t, 0, len(records.Record))
			assert.Error(t, err)
		}

		t.Logf("\tWhen it fails because there is parameters .")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

			records, err := ps.PsRecords("18102022", "", 0, nil)

			assert.Equal(t, 0, len(records.Record))
			assert.Error(t, err)

			records, err = ps.PsRecords("", "19102022", 0, nil)

			assert.Equal(t, 0, len(records.Record))
			assert.Error(t, err)
		}
	}
}
