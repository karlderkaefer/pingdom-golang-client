package client

import (
	"net"
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for the Pingdom API
	DefaultBaseURL = "https://api.pingdom.com/api/3.1"
)

func NewDefaultApiTokenClient(token string) *http.Client {
	
	// Create a new HTTP client
	httpClient := &http.Client{
		// Set the timeout for requests to 10 seconds
		Timeout: 10 * time.Second,
	}

	httpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// Add the Authorization header to the HTTP client
	httpClient.Transport = &authTransport{
		Transport: httpClient.Transport,
		Token:     token,
	}
	return httpClient
}

// authTransport is a custom transport that adds the Authorization header to the request
type authTransport struct {
    Transport http.RoundTripper
    Token     string
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    req.Header.Set("Authorization", "Bearer " + t.Token)
    return t.Transport.RoundTrip(req)
}
