package httprequest

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// NewRequest generate request.
func (r *HTTPRequest) NewRequest(method, url string, body io.Reader,
	query []map[string]interface{},
) ([]byte, int, error) {
	ctx := context.Background()

	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("httprequest.NewRequest.NewRequestWithContext: %w", err)
	}

	if query != nil {
		addQueryParameters(req, query)
	}

	resp, err := makeRequest(req, r.Username, r.Password)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("httprequest.NewRequest: %w", err)
	}

	defer resp.Body.Close()

	respBody, err := getBody(resp)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("httprequest.NewRequest: %w", err)
	}

	r.Logger.Info("request ", method, " ", url, "statusCode: ", resp.StatusCode, " exec_time ",
		float64(time.Since(start))/1000000, " ms")

	return respBody, resp.StatusCode, nil
}
