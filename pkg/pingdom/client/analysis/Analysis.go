// Package analysis provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package analysis

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

// AnalysisAnalysisid defines model for analysis_analysisid.
type AnalysisAnalysisid = int

// AnalysisCheckid defines model for analysis_checkid.
type AnalysisCheckid = int

// AnalysisFrom defines model for analysis_from.
type AnalysisFrom = int

// AnalysisLimit defines model for analysis_limit.
type AnalysisLimit = int

// AnalysisOffset defines model for analysis_offset.
type AnalysisOffset = int

// AnalysisRespAttrs defines model for analysis_resp_attrs.
type AnalysisRespAttrs struct {
	Analysis *[]struct {
		// Id Analysis id
		Id *int `json:"id,omitempty"`

		// Timeconfirmtest Time of the confirmation test that performed the error analysis
		Timeconfirmtest *int `json:"timeconfirmtest,omitempty"`

		// Timefirsttest Time of test that initiated the confirmation test
		Timefirsttest *int `json:"timefirsttest,omitempty"`
	} `json:"analysis,omitempty"`
}

// AnalysisTo defines model for analysis_to.
type AnalysisTo = int

// GetAnalysisCheckidParams defines parameters for GetAnalysisCheckid.
type GetAnalysisCheckidParams struct {
	// Limit Limits the number of returned results to the specified quantity.
	Limit *AnalysisLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Offset for listing. (Requires limit.)
	Offset *AnalysisOffset `form:"offset,omitempty" json:"offset,omitempty"`

	// From Return only results with timestamp of first test greater or equal to this value. Format is UNIX timestamp.
	From *AnalysisFrom `form:"from,omitempty" json:"from,omitempty"`

	// To Return only results with timestamp of first test less or equal to this value. Format is UNIX timestamp. Default: current timestamp
	To *AnalysisTo `form:"to,omitempty" json:"to,omitempty"`
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
	// GetAnalysisCheckid request
	GetAnalysisCheckid(ctx context.Context, checkid AnalysisCheckid, params *GetAnalysisCheckidParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetAnalysisCheckidAnalysisid request
	GetAnalysisCheckidAnalysisid(ctx context.Context, checkid AnalysisCheckid, analysisid AnalysisAnalysisid, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAnalysisCheckid(ctx context.Context, checkid AnalysisCheckid, params *GetAnalysisCheckidParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnalysisCheckidRequest(c.Server, checkid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetAnalysisCheckidAnalysisid(ctx context.Context, checkid AnalysisCheckid, analysisid AnalysisAnalysisid, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAnalysisCheckidAnalysisidRequest(c.Server, checkid, analysisid)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAnalysisCheckidRequest generates requests for GetAnalysisCheckid
func NewGetAnalysisCheckidRequest(server string, checkid AnalysisCheckid, params *GetAnalysisCheckidParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "checkid", runtime.ParamLocationPath, checkid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/analysis/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Limit != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "limit", runtime.ParamLocationQuery, *params.Limit); err != nil {
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

		if params.Offset != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "offset", runtime.ParamLocationQuery, *params.Offset); err != nil {
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

		if params.From != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "from", runtime.ParamLocationQuery, *params.From); err != nil {
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

		if params.To != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "to", runtime.ParamLocationQuery, *params.To); err != nil {
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

// NewGetAnalysisCheckidAnalysisidRequest generates requests for GetAnalysisCheckidAnalysisid
func NewGetAnalysisCheckidAnalysisidRequest(server string, checkid AnalysisCheckid, analysisid AnalysisAnalysisid) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "checkid", runtime.ParamLocationPath, checkid)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "analysisid", runtime.ParamLocationPath, analysisid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/analysis/%s/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
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
	// GetAnalysisCheckidWithResponse request
	GetAnalysisCheckidWithResponse(ctx context.Context, checkid AnalysisCheckid, params *GetAnalysisCheckidParams, reqEditors ...RequestEditorFn) (*GetAnalysisCheckidResponse, error)

	// GetAnalysisCheckidAnalysisidWithResponse request
	GetAnalysisCheckidAnalysisidWithResponse(ctx context.Context, checkid AnalysisCheckid, analysisid AnalysisAnalysisid, reqEditors ...RequestEditorFn) (*GetAnalysisCheckidAnalysisidResponse, error)
}

type GetAnalysisCheckidResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AnalysisRespAttrs
}

// Status returns HTTPResponse.Status
func (r GetAnalysisCheckidResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAnalysisCheckidResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAnalysisCheckidAnalysisidResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetAnalysisCheckidAnalysisidResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAnalysisCheckidAnalysisidResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAnalysisCheckidWithResponse request returning *GetAnalysisCheckidResponse
func (c *ClientWithResponses) GetAnalysisCheckidWithResponse(ctx context.Context, checkid AnalysisCheckid, params *GetAnalysisCheckidParams, reqEditors ...RequestEditorFn) (*GetAnalysisCheckidResponse, error) {
	rsp, err := c.GetAnalysisCheckid(ctx, checkid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnalysisCheckidResponse(rsp)
}

// GetAnalysisCheckidAnalysisidWithResponse request returning *GetAnalysisCheckidAnalysisidResponse
func (c *ClientWithResponses) GetAnalysisCheckidAnalysisidWithResponse(ctx context.Context, checkid AnalysisCheckid, analysisid AnalysisAnalysisid, reqEditors ...RequestEditorFn) (*GetAnalysisCheckidAnalysisidResponse, error) {
	rsp, err := c.GetAnalysisCheckidAnalysisid(ctx, checkid, analysisid, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAnalysisCheckidAnalysisidResponse(rsp)
}

// ParseGetAnalysisCheckidResponse parses an HTTP response from a GetAnalysisCheckidWithResponse call
func ParseGetAnalysisCheckidResponse(rsp *http.Response) (*GetAnalysisCheckidResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAnalysisCheckidResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AnalysisRespAttrs
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetAnalysisCheckidAnalysisidResponse parses an HTTP response from a GetAnalysisCheckidAnalysisidWithResponse call
func ParseGetAnalysisCheckidAnalysisidResponse(rsp *http.Response) (*GetAnalysisCheckidAnalysisidResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAnalysisCheckidAnalysisidResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}