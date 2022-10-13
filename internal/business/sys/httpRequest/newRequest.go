package httprequest

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"ps-go-client/internal/business/sys/errors"
)

// NewRequest generate request.
func (r *HTTPRequest) NewRequest(method, url string, body io.Reader,
	query map[string]interface{},
) ([]byte, int, error) {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("%s : %w", errors.ErrHTTPRequestCreate, err)
	}

	if query != nil {
		addQueryParameters(req, query)
	}

	resp, err := makeRequest(req)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("%w", err)
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	respBody, err := getBody(resp)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("%w", err)
	}

	return respBody, resp.StatusCode, nil
}
