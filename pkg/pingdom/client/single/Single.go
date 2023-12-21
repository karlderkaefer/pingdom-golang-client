// Package single provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package single

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for DNSType.
const (
	Dns DNSType = "dns"
)

// Defines values for HTTPType.
const (
	Http HTTPType = "http"
)

// Defines values for HTTPCustomType.
const (
	Httpcustom HTTPCustomType = "httpcustom"
)

// Defines values for IMAPType.
const (
	Imap IMAPType = "imap"
)

// Defines values for POP3Type.
const (
	Pop3 POP3Type = "pop3"
)

// Defines values for SMTPType.
const (
	Smtp SMTPType = "smtp"
)

// Defines values for TCPType.
const (
	Tcp TCPType = "tcp"
)

// Defines values for UDPType.
const (
	Udp UDPType = "udp"
)

// DNS defines model for DNS.
type DNS struct {
	// Expectedip (dns specific) Expected ip
	Expectedip string `json:"expectedip"`

	// Host Target host
	Host string `json:"host"`

	// Ipv6 Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`

	// Nameserver (dns specific) Nameserver
	Nameserver string `json:"nameserver"`

	// ProbeFilters Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are 'EU', 'NA', 'APAC' and 'LATAM'. For example, "region: NA".
	ProbeFilters *int `json:"probe_filters,omitempty"`

	// Probeid Probe identifier
	Probeid *int `json:"probeid,omitempty"`

	// ResponsetimeThreshold Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
	ResponsetimeThreshold *int    `json:"responsetime_threshold,omitempty"`
	Type                  DNSType `json:"type"`
}

// DNSType defines model for DNS.Type.
type DNSType string

// HTTP defines model for HTTP.
type HTTP struct {
	// Auth (http specific) Username and password for target HTTP authentication.
	Auth *string `json:"auth,omitempty"`

	// Encryption (http specific) Connection encryption
	Encryption *bool `json:"encryption,omitempty"`

	// Host Target host
	Host string `json:"host"`

	// Ipv6 Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`

	// Port (http specific) Target port
	Port *int `json:"port,omitempty"`

	// Postdata (http specific) Data that should be posted to the web page, for example submission data for a sign-up or login form. The data needs to be formatted in the same way as a web browser would send it to the web server
	Postdata *string `json:"postdata,omitempty"`

	// ProbeFilters Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are 'EU', 'NA', 'APAC' and 'LATAM'.
	ProbeFilters *int `json:"probe_filters,omitempty"`

	// Probeid Probe identifier
	Probeid *int `json:"probeid,omitempty"`

	// RequestheaderX (http specific) Custom HTTP header name. Replace {X} with a number unique for each header argument.
	RequestheaderX *string `json:"requestheader{X},omitempty"`

	// ResponsetimeThreshold Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
	ResponsetimeThreshold *int `json:"responsetime_threshold,omitempty"`

	// Shouldcontain (http specific) Target site should contain this string
	Shouldcontain *string `json:"shouldcontain,omitempty"`

	// Shouldnotcontain (http specific) Target site should NOT contain this string
	Shouldnotcontain *string `json:"shouldnotcontain,omitempty"`

	// SslDownDaysBefore (http specific) Treat the target site as down if a certificate expires within the given number of days. This parameter will be ignored if `verify_certificate` is set to `false`.
	SslDownDaysBefore *int     `json:"ssl_down_days_before,omitempty"`
	Type              HTTPType `json:"type"`

	// Url (http specific) Target path on server
	Url *string `json:"url,omitempty"`

	// VerifyCertificate (http specific) Treat target site as down if an invalid/unverifiable certificate is found.
	VerifyCertificate *bool `json:"verify_certificate,omitempty"`
}

// HTTPType defines model for HTTP.Type.
type HTTPType string

// HTTPCustom defines model for HTTP-Custom.
type HTTPCustom struct {
	// Additionalurls (httpcustom specific) ;-separated list of addidional URLs with hostname included.
	Additionalurls *string `json:"additionalurls,omitempty"`

	// Auth (httpcustom specific) Username and password for target HTTP authentication.
	Auth *string `json:"auth,omitempty"`

	// Encryption (httpcustom specific) Connection encryption
	Encryption *bool `json:"encryption,omitempty"`

	// Host Target host
	Host string `json:"host"`

	// Ipv6 Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`

	// Port (httpcustom specific) Target port
	Port *int `json:"port,omitempty"`

	// ProbeFilters Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are 'EU', 'NA', 'APAC' and 'LATAM'.
	ProbeFilters *int `json:"probe_filters,omitempty"`

	// Probeid Probe identifier
	Probeid *int `json:"probeid,omitempty"`

	// ResponsetimeThreshold Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
	ResponsetimeThreshold *int           `json:"responsetime_threshold,omitempty"`
	Type                  HTTPCustomType `json:"type"`

	// Url (httpcustom specific) Target path to XML file on server
	Url string `json:"url"`
}

// HTTPCustomType defines model for HTTPCustom.Type.
type HTTPCustomType string

// IMAP defines model for IMAP.
type IMAP struct {
	// Encryption (imap specific) Connection encryption
	Encryption *bool `json:"encryption,omitempty"`

	// Host Target host
	Host string `json:"host"`

	// Ipv6 Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`

	// Port (imap specific) Target port
	Port *int `json:"port,omitempty"`

	// ProbeFilters Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are 'EU', 'NA', 'APAC' and 'LATAM'.
	ProbeFilters *int `json:"probe_filters,omitempty"`

	// Probeid Probe identifier
	Probeid *int `json:"probeid,omitempty"`

	// ResponsetimeThreshold Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
	ResponsetimeThreshold *int `json:"responsetime_threshold,omitempty"`

	// Stringtoexpect (imap specific) String to expect in response
	Stringtoexpect *string  `json:"stringtoexpect,omitempty"`
	Type           IMAPType `json:"type"`
}

// IMAPType defines model for IMAP.Type.
type IMAPType string

// POP3 defines model for POP3.
type POP3 struct {
	// Encryption (pop3 specific) Connection encryption
	Encryption *bool `json:"encryption,omitempty"`

	// Host Target host
	Host string `json:"host"`

	// Ipv6 Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`

	// Port (pop3 specific) Target port
	Port *int `json:"port,omitempty"`

	// ProbeFilters Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are 'EU', 'NA', 'APAC' and 'LATAM'.
	ProbeFilters *int `json:"probe_filters,omitempty"`

	// Probeid Probe identifier
	Probeid *int `json:"probeid,omitempty"`

	// ResponsetimeThreshold Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
	ResponsetimeThreshold *int `json:"responsetime_threshold,omitempty"`

	// Stringtoexpect (pop3 specific) String to expect in response
	Stringtoexpect *string  `json:"stringtoexpect,omitempty"`
	Type           POP3Type `json:"type"`
}

// POP3Type defines model for POP3.Type.
type POP3Type string

// SMTP defines model for SMTP.
type SMTP struct {
	// Auth (smtp specific) Username and password for target HTTP authentication.
	Auth *string `json:"auth,omitempty"`

	// Encryption (smtp specific) Connection encryption
	Encryption *bool `json:"encryption,omitempty"`

	// Host Target host
	Host string `json:"host"`

	// Ipv6 Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`

	// Port (smtp specific) Target port
	Port *int `json:"port,omitempty"`

	// ProbeFilters Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are 'EU', 'NA', 'APAC' and 'LATAM'.
	ProbeFilters *int `json:"probe_filters,omitempty"`

	// Probeid Probe identifier
	Probeid *int `json:"probeid,omitempty"`

	// ResponsetimeThreshold Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
	ResponsetimeThreshold *int `json:"responsetime_threshold,omitempty"`

	// Stringtoexpect (smtp specific) String to expect in response
	Stringtoexpect *string  `json:"stringtoexpect,omitempty"`
	Type           SMTPType `json:"type"`
}

// SMTPType defines model for SMTP.Type.
type SMTPType string

// TCP defines model for TCP.
type TCP struct {
	// Host Target host
	Host string `json:"host"`

	// Ipv6 Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`

	// Port (tcp specific) Target port
	Port int `json:"port"`

	// ProbeFilters Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are 'EU', 'NA', 'APAC' and 'LATAM'.
	ProbeFilters *int `json:"probe_filters,omitempty"`

	// Probeid Probe identifier
	Probeid *int `json:"probeid,omitempty"`

	// ResponsetimeThreshold Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
	ResponsetimeThreshold *int `json:"responsetime_threshold,omitempty"`

	// Stringtoexpect (tcp specific) String to expect in response
	Stringtoexpect *string `json:"stringtoexpect,omitempty"`

	// Stringtosend (tcp specific) String to send
	Stringtosend *string `json:"stringtosend,omitempty"`
	Type         TCPType `json:"type"`
}

// TCPType defines model for TCP.Type.
type TCPType string

// UDP defines model for UDP.
type UDP struct {
	// Host Target host
	Host string `json:"host"`

	// Ipv6 Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`

	// Port (udp specific) Target port
	Port int `json:"port"`

	// ProbeFilters Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are 'EU', 'NA', 'APAC' and 'LATAM'.
	ProbeFilters *int `json:"probe_filters,omitempty"`

	// Probeid Probe identifier
	Probeid *int `json:"probeid,omitempty"`

	// ResponsetimeThreshold Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
	ResponsetimeThreshold *int `json:"responsetime_threshold,omitempty"`

	// Stringtoexpect (udp specific) String to expect in response
	Stringtoexpect *string `json:"stringtoexpect,omitempty"`

	// Stringtosend (udp specific) String to send
	Stringtosend *string `json:"stringtosend,omitempty"`
	Type         UDPType `json:"type"`
}

// UDPType defines model for UDP.Type.
type UDPType string

// SingleResp defines model for single_resp.
type SingleResp struct {
	Result *struct {
		// Probedesc Probe description
		Probedesc *string `json:"probedesc,omitempty"`

		// Probeid Probe identifier
		Probeid *int `json:"probeid,omitempty"`

		// Responsetime Response time in milliseconds
		Responsetime *int `json:"responsetime,omitempty"`

		// Status Test result status (up, down)
		Status *string `json:"status,omitempty"`

		// Statusdesc Short status description
		Statusdesc *string `json:"statusdesc,omitempty"`

		// Statusdesclong Long status description
		Statusdesclong *string `json:"statusdesclong,omitempty"`
	} `json:"result,omitempty"`
}

// GetSingleParams defines parameters for GetSingle.
type GetSingleParams struct {
	// QueryParameters Query Parameters to chose
	QueryParameters *struct {
		union json.RawMessage
	} `form:"Query Parameters,omitempty" json:"Query Parameters,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetSingle request
	GetSingle(ctx context.Context, params *GetSingleParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetSingle(ctx context.Context, params *GetSingleParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSingleRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetSingleRequest generates requests for GetSingle
func NewGetSingleRequest(server string, params *GetSingleParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/single")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.QueryParameters != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "Query Parameters", runtime.ParamLocationQuery, *params.QueryParameters); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetSingleWithResponse request
	GetSingleWithResponse(ctx context.Context, params *GetSingleParams, reqEditors ...RequestEditorFn) (*GetSingleResponse, error)
}

type GetSingleResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SingleResp
}

// Status returns HTTPResponse.Status
func (r GetSingleResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSingleResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetSingleWithResponse request returning *GetSingleResponse
func (c *ClientWithResponses) GetSingleWithResponse(ctx context.Context, params *GetSingleParams, reqEditors ...RequestEditorFn) (*GetSingleResponse, error) {
	rsp, err := c.GetSingle(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSingleResponse(rsp)
}

// ParseGetSingleResponse parses an HTTP response from a GetSingleWithResponse call
func ParseGetSingleResponse(rsp *http.Response) (*GetSingleResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSingleResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SingleResp
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
