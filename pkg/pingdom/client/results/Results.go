// Package results provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package results

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

// Defines values for ResultsRespAttrsResultsStatus.
const (
	Down        ResultsRespAttrsResultsStatus = "down"
	Unconfirmed ResultsRespAttrsResultsStatus = "unconfirmed"
	Unknown     ResultsRespAttrsResultsStatus = "unknown"
	Up          ResultsRespAttrsResultsStatus = "up"
)

// ResultsCheckid defines model for results_checkid.
type ResultsCheckid = int

// ResultsFrom defines model for results_from.
type ResultsFrom = int

// ResultsIncludeanalysis defines model for results_includeanalysis.
type ResultsIncludeanalysis = bool

// ResultsLimit defines model for results_limit.
type ResultsLimit = int

// ResultsMaxresponse defines model for results_maxresponse.
type ResultsMaxresponse = int

// ResultsMinresponse defines model for results_minresponse.
type ResultsMinresponse = int

// ResultsOffset defines model for results_offset.
type ResultsOffset = int

// ResultsProbes Default: all probes
type ResultsProbes = string

// ResultsRespAttrs defines model for results_resp_attrs.
type ResultsRespAttrs struct {
	// Activeprobes For your convenience, a list of used probes that produced the showed results
	Activeprobes *[]float32 `json:"activeprobes,omitempty"`
	Results      *[]struct {
		// Probeid Probe identifier
		Probeid *float32 `json:"probeid,omitempty"`

		// Responsetime Response time (in milliseconds) (Will be 0 if no response was received)
		Responsetime *float32 `json:"responsetime,omitempty"`

		// Status Result status
		Status *ResultsRespAttrsResultsStatus `json:"status,omitempty"`

		// Statusdesc Short status description
		Statusdesc *string `json:"statusdesc,omitempty"`

		// Statusdesclong Long status description
		Statusdesclong *string `json:"statusdesclong,omitempty"`

		// Time Time when test was performed. Format is UNIX timestamp
		Time *float32 `json:"time,omitempty"`
	} `json:"results,omitempty"`
}

// ResultsRespAttrsResultsStatus Result status
type ResultsRespAttrsResultsStatus string

// ResultsStatus Default: all statuses
type ResultsStatus = string

// ResultsTo defines model for results_to.
type ResultsTo = int

// GetResultsCheckidParams defines parameters for GetResultsCheckid.
type GetResultsCheckidParams struct {
	// To End of period. Format is UNIX timestamp. Default value is current timestamp.
	To *ResultsTo `form:"to,omitempty" json:"to,omitempty"`

	// From Start of period. Format is UNIX timestamp. Default value is 1 day prior to `to`.
	From *ResultsFrom `form:"from,omitempty" json:"from,omitempty"`

	// Probes Filter to only show results from a list of probes. Format is a comma separated list of probe identifiers
	Probes *ResultsProbes `form:"probes,omitempty" json:"probes,omitempty"`

	// Status Filter to only show results with specified statuses. Format is a comma separated list of (`down`, `up`, `unconfirmed`, `unknown`)
	Status *ResultsStatus `form:"status,omitempty" json:"status,omitempty"`

	// Limit Number of results to show (Will be set to 1000 if the provided value is greater than 1000)
	Limit *ResultsLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Number of results to skip (Max value is `43200`)
	Offset *ResultsOffset `form:"offset,omitempty" json:"offset,omitempty"`

	// Includeanalysis Attach available root cause analysis identifiers to corresponding results
	Includeanalysis *ResultsIncludeanalysis `form:"includeanalysis,omitempty" json:"includeanalysis,omitempty"`

	// Maxresponse Maximum response time (ms). If set, specified interval must not be larger than 31 days.
	Maxresponse *ResultsMaxresponse `form:"maxresponse,omitempty" json:"maxresponse,omitempty"`

	// Minresponse Minimum response time (ms). If set, specified interval must not be larger than 31 days.
	Minresponse *ResultsMinresponse `form:"minresponse,omitempty" json:"minresponse,omitempty"`
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
	// GetResultsCheckid request
	GetResultsCheckid(ctx context.Context, checkid ResultsCheckid, params *GetResultsCheckidParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetResultsCheckid(ctx context.Context, checkid ResultsCheckid, params *GetResultsCheckidParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetResultsCheckidRequest(c.Server, checkid, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetResultsCheckidRequest generates requests for GetResultsCheckid
func NewGetResultsCheckidRequest(server string, checkid ResultsCheckid, params *GetResultsCheckidParams) (*http.Request, error) {
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

	operationPath := fmt.Sprintf("/results/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

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

		if params.Probes != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "probes", runtime.ParamLocationQuery, *params.Probes); err != nil {
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

		if params.Status != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "status", runtime.ParamLocationQuery, *params.Status); err != nil {
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

		if params.Includeanalysis != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "includeanalysis", runtime.ParamLocationQuery, *params.Includeanalysis); err != nil {
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

		if params.Maxresponse != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maxresponse", runtime.ParamLocationQuery, *params.Maxresponse); err != nil {
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

		if params.Minresponse != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "minresponse", runtime.ParamLocationQuery, *params.Minresponse); err != nil {
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
	// GetResultsCheckidWithResponse request
	GetResultsCheckidWithResponse(ctx context.Context, checkid ResultsCheckid, params *GetResultsCheckidParams, reqEditors ...RequestEditorFn) (*GetResultsCheckidResponse, error)
}

type GetResultsCheckidResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ResultsRespAttrs
}

// Status returns HTTPResponse.Status
func (r GetResultsCheckidResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetResultsCheckidResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetResultsCheckidWithResponse request returning *GetResultsCheckidResponse
func (c *ClientWithResponses) GetResultsCheckidWithResponse(ctx context.Context, checkid ResultsCheckid, params *GetResultsCheckidParams, reqEditors ...RequestEditorFn) (*GetResultsCheckidResponse, error) {
	rsp, err := c.GetResultsCheckid(ctx, checkid, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetResultsCheckidResponse(rsp)
}

// ParseGetResultsCheckidResponse parses an HTTP response from a GetResultsCheckidWithResponse call
func ParseGetResultsCheckidResponse(rsp *http.Response) (*GetResultsCheckidResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetResultsCheckidResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ResultsRespAttrs
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
