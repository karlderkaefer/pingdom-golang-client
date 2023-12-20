// Package actions provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package actions

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

// Defines values for ActionsAlertsEntryActionsAlertsStatus.
const (
	Delivered    ActionsAlertsEntryActionsAlertsStatus = "delivered"
	Error        ActionsAlertsEntryActionsAlertsStatus = "error"
	Nocredits    ActionsAlertsEntryActionsAlertsStatus = "nocredits"
	Notdelivered ActionsAlertsEntryActionsAlertsStatus = "notdelivered"
	Sent         ActionsAlertsEntryActionsAlertsStatus = "sent"
)

// Defines values for ActionsAlertsEntryActionsAlertsVia.
const (
	Agcm    ActionsAlertsEntryActionsAlertsVia = "agcm"
	Apns    ActionsAlertsEntryActionsAlertsVia = "apns"
	Email   ActionsAlertsEntryActionsAlertsVia = "email"
	Sms     ActionsAlertsEntryActionsAlertsVia = "sms"
	Twitter ActionsAlertsEntryActionsAlertsVia = "twitter"
)

// ActionsAlertsEntry defines model for actions_alerts_entry.
type ActionsAlertsEntry struct {
	Actions *struct {
		// Alerts Alert entry
		Alerts *[]struct {
			// Charged True if your account was charged for this message
			Charged *string `json:"charged,omitempty"`

			// Checkid Identifier of alerted user
			Checkid *string `json:"checkid,omitempty"`

			// Messagefull Full message body
			Messagefull *string `json:"messagefull,omitempty"`

			// Messageshort Short description of message
			Messageshort *string `json:"messageshort,omitempty"`

			// Sentto Target address, phone number etc
			Sentto *string `json:"sentto,omitempty"`

			// Status Alert status
			Status *ActionsAlertsEntryActionsAlertsStatus `json:"status,omitempty"`

			// Time Time of alert generation. Format UNIX time
			Time *string `json:"time,omitempty"`

			// Userid Identifier of alerted user
			Userid *string `json:"userid,omitempty"`

			// Username Name of alerted user
			Username *string `json:"username,omitempty"`

			// Via Alert medium. apns - iphone, agcm - android
			Via *ActionsAlertsEntryActionsAlertsVia `json:"via,omitempty"`
		} `json:"alerts,omitempty"`
	} `json:"actions,omitempty"`
}

// ActionsAlertsEntryActionsAlertsStatus Alert status
type ActionsAlertsEntryActionsAlertsStatus string

// ActionsAlertsEntryActionsAlertsVia Alert medium. apns - iphone, agcm - android
type ActionsAlertsEntryActionsAlertsVia string

// ActionsCheckids defines model for actions_checkids.
type ActionsCheckids = string

// ActionsFrom defines model for actions_from.
type ActionsFrom = int

// ActionsLimit defines model for actions_limit.
type ActionsLimit = int

// ActionsOffset defines model for actions_offset.
type ActionsOffset = int

// ActionsStatus defines model for actions_status.
type ActionsStatus = string

// ActionsTo defines model for actions_to.
type ActionsTo = int

// ActionsUserids defines model for actions_userids.
type ActionsUserids = string

// ActionsVia defines model for actions_via.
type ActionsVia = string

// GetActionsParams defines parameters for GetActions.
type GetActionsParams struct {
	// From Only include actions generated later than this timestamp. Format is UNIX time.
	From *ActionsFrom `form:"from,omitempty" json:"from,omitempty"`

	// To Only include actions generated prior to this timestamp. Format is UNIX time.
	To *ActionsTo `form:"to,omitempty" json:"to,omitempty"`

	// Limit Limits the number of returned results to the specified quantity.
	Limit *ActionsLimit `form:"limit,omitempty" json:"limit,omitempty"`

	// Offset Offset for listing.
	Offset *ActionsOffset `form:"offset,omitempty" json:"offset,omitempty"`

	// Checkids Comma-separated list of check identifiers. Limit results to actions generated from these checks. Default: all checks.
	Checkids *ActionsCheckids `form:"checkids,omitempty" json:"checkids,omitempty"`

	// Userids Comma-separated list of user identifiers. Limit results to actions sent to these users. Default: all users.
	Userids *ActionsUserids `form:"userids,omitempty" json:"userids,omitempty"`

	// Status Comma-separated list of statuses. Limit results to actions with these statuses. Default: all statuses.
	Status *ActionsStatus `form:"status,omitempty" json:"status,omitempty"`

	// Via Comma-separated list of via mediums. Limit results to actions with these mediums. Default: all mediums.
	Via *ActionsVia `form:"via,omitempty" json:"via,omitempty"`
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
	// GetActions request
	GetActions(ctx context.Context, params *GetActionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetActions(ctx context.Context, params *GetActionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetActionsRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetActionsRequest generates requests for GetActions
func NewGetActionsRequest(server string, params *GetActionsParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/actions")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

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

		if params.Checkids != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "checkids", runtime.ParamLocationQuery, *params.Checkids); err != nil {
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

		if params.Userids != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "userids", runtime.ParamLocationQuery, *params.Userids); err != nil {
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

		if params.Via != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "via", runtime.ParamLocationQuery, *params.Via); err != nil {
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
	// GetActionsWithResponse request
	GetActionsWithResponse(ctx context.Context, params *GetActionsParams, reqEditors ...RequestEditorFn) (*GetActionsResponse, error)
}

type GetActionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ActionsAlertsEntry
}

// Status returns HTTPResponse.Status
func (r GetActionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetActionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetActionsWithResponse request returning *GetActionsResponse
func (c *ClientWithResponses) GetActionsWithResponse(ctx context.Context, params *GetActionsParams, reqEditors ...RequestEditorFn) (*GetActionsResponse, error) {
	rsp, err := c.GetActions(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetActionsResponse(rsp)
}

// ParseGetActionsResponse parses an HTTP response from a GetActionsWithResponse call
func ParseGetActionsResponse(rsp *http.Response) (*GetActionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetActionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ActionsAlertsEntry
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
