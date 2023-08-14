package httprequest

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// NewRequest generate request.
func (r *HTTPRequest) NewRequest(method, url string, body io.Reader,
	query []map[string]interface{},
) ([]byte, int, error) {
	ctx := context.Background()

	start := time.Now()

	log.Printf("PowerStudio started    method: %s  url: %s\n", method, url)

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

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			//nolint: forbidigo
			fmt.Println(err)
		}
	}(resp.Body)

	respBody, err := getBody(resp)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("httprequest.NewRequest: %w", err)
	}

	log.Printf("PowerStudio completed  method: %s  url: %s  statusCode: %d  exec_time: %f\n",
		method, url, resp.StatusCode, float64(time.Since(start))/1000000)

	return respBody, resp.StatusCode, nil
}
