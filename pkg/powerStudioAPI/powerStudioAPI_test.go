package powerstudioapi_test

import (
	"io"
	"os"
	"ps-go-client/internal/business/sys/errors"
	"ps-go-client/internal/business/sys/httpRequest/mocks"
	powerStudioAPI "ps-go-client/pkg/powerStudioAPI"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAllDevices(t *testing.T) {
	t.Parallel()

	t.Logf("Given the need to call AllDevices method API.")
	{
		t.Logf("\tWhen a correct api call.")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

			xmlFile, err := os.Open("../../internal/business/sampleEntities/allDevices.xml")
			require.NoError(t, err)

			byteValue, err := io.ReadAll(xmlFile)
			require.NoError(t, err)

			mock := new(mocks.RequestMock)
			mock.On("NewRequest", "GET", "http://localhost/services/user/devices.xml", nil, map[string]interface{}(nil)).
				Return(byteValue, 200, nil)

			ps.Request = mock

			devices, err := ps.PsAllDevices()

			assert.Nil(t, err)
			assert.GreaterOrEqual(t, len(devices.ID), 0)
		}

		t.Logf("\tWhen it fails because data unmarchal error.")
		{
			ps := powerStudioAPI.NewPowerStudio("localhost")

			mock := new(mocks.RequestMock)
			mock.On("NewRequest", "GET", "http://localhost/services/user/devices.xml", nil, map[string]interface{}(nil)).
				Return([]byte(""), 200, nil)

			ps.Request = mock

			devices, err := ps.PsAllDevices()

			assert.Error(t, err)
			assert.Equal(t, 0, len(devices.ID))
		}

		t.Logf("\tWhen it fails because power studio error.")
		{
			ps := powerStudioAPI.NewPowerStudio("10.10.10.10")

			mock := new(mocks.RequestMock)
			mock.On("NewRequest", "GET", "http://10.10.10.10/services/user/devices.xml", nil, map[string]interface{}(nil)).
				Return([]byte(""), 404, nil)

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
