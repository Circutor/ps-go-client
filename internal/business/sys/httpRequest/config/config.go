package config

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/icholy/digest"
	"golang.org/x/net/publicsuffix"
)

const timeout = 20

// customTransport transport configuration.
type customTransport struct {
	http.RoundTripper
}

// CreateHTTPClient created Http client configuration.
func CreateHTTPClient(username, password string) http.Client {
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       timeout * time.Second,
		Transport:     newTransport(http.DefaultTransport.(*http.Transport)),
		Jar:           newJar(),
	}

	if username != "" || password != "" {
		client.Transport = &digest.Transport{Username: username, Password: password}
	}

	return client
}

// newTransport configuration Transport.
func newTransport(upstream *http.Transport) *customTransport {
	// TLS config
	tlsConfig := new(tls.Config)
	tlsConfig.InsecureSkipVerify = true

	upstream.TLSClientConfig = tlsConfig

	return &customTransport{upstream}
}

// newJar configuration cookieJar.
func newJar() *cookiejar.Jar {
	// All users of cookiejar.
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	return jar
}
