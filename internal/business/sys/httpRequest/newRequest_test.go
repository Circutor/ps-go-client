package httprequest_test

import (
	"net/http"
	"testing"

	httpRequest "github.com/circutor/ps-go-client/internal/business/sys/httpRequest"
	"github.com/stretchr/testify/assert"
)

func TestErrorRequest(t *testing.T) {
	t.Parallel()

	r := httpRequest.HTTPRequest{}

	_, status, err := r.NewRequest("GET\000", "/", nil, nil)

	assert.Error(t, err)
	assert.Equal(t, status, http.StatusInternalServerError)

	_, status, err = r.NewRequest("GET", "/", nil, nil)

	assert.Error(t, err)
	assert.Equal(t, status, http.StatusInternalServerError)

	r = httpRequest.NewHTTPRequest("", "")

	_, status, err = r.NewRequest("GET\000", "/", nil, nil)

	assert.Error(t, err)
	assert.Equal(t, status, http.StatusInternalServerError)

	_, status, err = r.NewRequest("GET", "/", nil, nil)

	assert.Error(t, err)
	assert.Equal(t, status, http.StatusInternalServerError)
}

func TestRequestContentQuery(t *testing.T) {
	t.Parallel()

	r := httpRequest.HTTPRequest{}

	query := []map[string]interface{}{
		{
			"status": false,
		},
	}

	_, status, _ := r.NewRequest("GET", "/", nil, query)
	assert.Equal(t, status, http.StatusInternalServerError)
}
