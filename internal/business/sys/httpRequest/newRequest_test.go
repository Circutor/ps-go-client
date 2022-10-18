package httprequest_test

import (
	httpRequest "github.com/circutor/ps-go-client/internal/business/sys/httpRequest"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorRequest(t *testing.T) {
	t.Parallel()

	r := httpRequest.HTTPRequest{}

	_, status, _ := r.NewRequest("GET", "/", nil, nil)
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
