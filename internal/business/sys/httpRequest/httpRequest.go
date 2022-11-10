package httprequest

import (
	"fmt"
	"io"
	"net/http"

	"github.com/circutor/ps-go-client/internal/business/sys/httpRequest/config"
)

// Request interface created a new request.
type Request interface {
	NewRequest(method, url string, body io.Reader, query []map[string]interface{}) ([]byte, int, error)
}

//go:generate mockery --name Request --structname RequestMock --filename RequestMock.go

type HTTPRequest struct {
	Username string
	Password string
}

// NewHTTPRequest creates a new NewRequest interface.
func NewHTTPRequest(username, password string) HTTPRequest {
	return HTTPRequest{
		Username: username,
		Password: password,
	}
}

// addQueryParameters method aggregate queries in to the request.
func addQueryParameters(req *http.Request, queryParameters []map[string]interface{}) {
	query := req.URL.Query()

	for _, element := range queryParameters {
		for key, param := range element {
			query.Add(key, fmt.Sprintf("%v", param))
		}
	}

	req.URL.RawQuery = query.Encode()
}

// Method to make the request and return the response.
func makeRequest(req *http.Request, username, password string) (*http.Response, error) {
	client := config.CreateHTTPClient(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httprequest.makeRequest: %w", err)
	}

	return resp, nil
}

// getBody get content body.
func getBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("httprequest.getBody: %w", err)
	}

	return body, nil
}
