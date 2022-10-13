package httprequest_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"os"
	httprequest "ps-go-client/internal/business/sys/httpRequest"
	"ps-go-client/internal/business/sys/httpRequest/mocks"
	"testing"
)

func TestNewHTTPRequest(t *testing.T) {
	t.Parallel()

	xmlFile, err := os.Open("../../sampleEntities/allDevices.xml")
	require.NoError(t, err)

	byteValue, err := io.ReadAll(xmlFile)
	require.NoError(t, err)

	uri := "http://localhost/services/user/devices.xml"

	mock := new(mocks.RequestMock)
	mock.On("NewRequest", "POST", uri, nil, map[string]interface{}(nil)).Return(byteValue, 200, nil)

	reqHttp, codeHttp, err := mock.NewRequest("POST", uri, nil, nil)

	assert.Equal(t, nil, err)
	assert.Equal(t, string(byteValue), string(reqHttp))
	assert.Equal(t, 200, codeHttp)
}

func TestErrorRequest(t *testing.T) {
	t.Parallel()

	r := httprequest.HTTPRequest{}

	_, status, _ := r.NewRequest("GET", "/", nil, nil)
	assert.Equal(t, status, http.StatusInternalServerError)
}

func TestRequestContentQuery(t *testing.T) {
	t.Parallel()

	r := httprequest.HTTPRequest{}

	query := map[string]interface{}{
		"status": false,
	}

	_, status, _ := r.NewRequest("GET", "/", nil, query)
	assert.Equal(t, status, http.StatusInternalServerError)
}
