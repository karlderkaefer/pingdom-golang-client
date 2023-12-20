// Package maintenanceoccurrences provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package maintenanceoccurrences

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// MaintenanceOccurrencesDelete defines model for maintenance.occurrences_delete.
type MaintenanceOccurrencesDelete struct {
	// Occurrenceids Comma-separated list of identifiers of maintenance occurrences to delete.
	Occurrenceids []int `json:"occurrenceids"`
}

// MaintenanceOccurrencesDeleteRespAttrs defines model for maintenance.occurrences_delete_resp_attrs.
type MaintenanceOccurrencesDeleteRespAttrs struct {
	// Message Result description
	Message *string `json:"message,omitempty"`
}

// MaintenanceOccurrencesFrom defines model for maintenance.occurrences_from.
type MaintenanceOccurrencesFrom = int

// MaintenanceOccurrencesId defines model for maintenance.occurrences_id.
type MaintenanceOccurrencesId = int

// MaintenanceOccurrencesIdDeleteRespAttrs defines model for maintenance.occurrences_id_delete_resp_attrs.
type MaintenanceOccurrencesIdDeleteRespAttrs struct {
	// Message Result description
	Message *string `json:"message,omitempty"`
}

// MaintenanceOccurrencesIdPut defines model for maintenance.occurrences_id_put.
type MaintenanceOccurrencesIdPut struct {
	// From Beginning of the maintenance occurrence. Format UNIX time. (Only future allowed. Use 1 for the current timestamp.)
	From *int `json:"from,omitempty"`

	// To End of the maintenance occurrence. Format UNIX time. (Only future allowed. Use 1 for the current timestamp.)
	To *int `json:"to,omitempty"`
}

// MaintenanceOccurrencesIdPutRespAttrs defines model for maintenance.occurrences_id_put_resp_attrs.
type MaintenanceOccurrencesIdPutRespAttrs struct {
	// Message Modification result description
	Message *string `json:"message,omitempty"`
}

// MaintenanceOccurrencesIdRespAttrs defines model for maintenance.occurrences_id_resp_attrs.
type MaintenanceOccurrencesIdRespAttrs struct {
	Occurrence *struct {
		// From Beginning of the occurrence. Format UNIX timestamp.
		From *float32 `json:"from,omitempty"`

		// Id Identifier of an occurence.
		Id *float32 `json:"id,omitempty"`

		// Maintenanceid Identifier of the related maintenance window.
		Maintenanceid *float32 `json:"maintenanceid,omitempty"`

		// To The end of the occurrence. Format UNIX timestamp.
		To *float32 `json:"to,omitempty"`
	} `json:"occurrence,omitempty"`
}

// MaintenanceOccurrencesMaintenanceid defines model for maintenance.occurrences_maintenanceid.
type MaintenanceOccurrencesMaintenanceid = int

// MaintenanceOccurrencesRespAttrs defines model for maintenance.occurrences_resp_attrs.
type MaintenanceOccurrencesRespAttrs struct {
	// Occurrences A list of maintenance occurrences.
	Occurrences *[]struct {
		// From Beginning of the occurrence. Format UNIX timestamp.
		From *float32 `json:"from,omitempty"`

		// Id Identifier of an occurence.
		Id *float32 `json:"id,omitempty"`

		// Maintenanceid Identifier of the related maintenance window.
		Maintenanceid *float32 `json:"maintenanceid,omitempty"`

		// To The end of the occurrence. Format UNIX timestamp.
		To *float32 `json:"to,omitempty"`
	} `json:"occurrences,omitempty"`
}

// MaintenanceOccurrencesTo defines model for maintenance.occurrences_to.
type MaintenanceOccurrencesTo = int

// DeleteMaintenanceOccurrencesParams defines parameters for DeleteMaintenanceOccurrences.
type DeleteMaintenanceOccurrencesParams struct {
	Occurenceids MaintenanceOccurrencesDelete `form:"occurenceids" json:"occurenceids"`
}

// GetMaintenanceOccurrencesParams defines parameters for GetMaintenanceOccurrences.
type GetMaintenanceOccurrencesParams struct {
	// Maintenanceid Maintenance window identifier. (List only occurrences of a specific maintenance window.)
	Maintenanceid *MaintenanceOccurrencesMaintenanceid `form:"maintenanceid,omitempty" json:"maintenanceid,omitempty"`

	// From Effective from (unix timestamp). (List occurrences which are effective from the specified unix timestamp. If not specified, current timestamp is used.)
	From *MaintenanceOccurrencesFrom `form:"from,omitempty" json:"from,omitempty"`

	// To Effective to (unix timestamp). (List occurrences which are effective to the specified unix timestamp.)
	To *MaintenanceOccurrencesTo `form:"to,omitempty" json:"to,omitempty"`
}

// PutMaintenanceOccurrencesIdJSONRequestBody defines body for PutMaintenanceOccurrencesId for application/json ContentType.
type PutMaintenanceOccurrencesIdJSONRequestBody = MaintenanceOccurrencesIdPut

// PutMaintenanceOccurrencesIdFormdataRequestBody defines body for PutMaintenanceOccurrencesId for application/x-www-form-urlencoded ContentType.
type PutMaintenanceOccurrencesIdFormdataRequestBody = MaintenanceOccurrencesIdPut

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
	// DeleteMaintenanceOccurrences request
	DeleteMaintenanceOccurrences(ctx context.Context, params *DeleteMaintenanceOccurrencesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetMaintenanceOccurrences request
	GetMaintenanceOccurrences(ctx context.Context, params *GetMaintenanceOccurrencesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteMaintenanceOccurrencesId request
	DeleteMaintenanceOccurrencesId(ctx context.Context, id MaintenanceOccurrencesId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetMaintenanceOccurrencesId request
	GetMaintenanceOccurrencesId(ctx context.Context, id MaintenanceOccurrencesId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutMaintenanceOccurrencesIdWithBody request with any body
	PutMaintenanceOccurrencesIdWithBody(ctx context.Context, id MaintenanceOccurrencesId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutMaintenanceOccurrencesId(ctx context.Context, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutMaintenanceOccurrencesIdWithFormdataBody(ctx context.Context, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) DeleteMaintenanceOccurrences(ctx context.Context, params *DeleteMaintenanceOccurrencesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteMaintenanceOccurrencesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetMaintenanceOccurrences(ctx context.Context, params *GetMaintenanceOccurrencesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMaintenanceOccurrencesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteMaintenanceOccurrencesId(ctx context.Context, id MaintenanceOccurrencesId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteMaintenanceOccurrencesIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetMaintenanceOccurrencesId(ctx context.Context, id MaintenanceOccurrencesId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetMaintenanceOccurrencesIdRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutMaintenanceOccurrencesIdWithBody(ctx context.Context, id MaintenanceOccurrencesId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutMaintenanceOccurrencesIdRequestWithBody(c.Server, id, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutMaintenanceOccurrencesId(ctx context.Context, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutMaintenanceOccurrencesIdRequest(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutMaintenanceOccurrencesIdWithFormdataBody(ctx context.Context, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutMaintenanceOccurrencesIdRequestWithFormdataBody(c.Server, id, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDeleteMaintenanceOccurrencesRequest generates requests for DeleteMaintenanceOccurrences
func NewDeleteMaintenanceOccurrencesRequest(server string, params *DeleteMaintenanceOccurrencesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/maintenance.occurrences")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "occurenceids", runtime.ParamLocationQuery, params.Occurenceids); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetMaintenanceOccurrencesRequest generates requests for GetMaintenanceOccurrences
func NewGetMaintenanceOccurrencesRequest(server string, params *GetMaintenanceOccurrencesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/maintenance.occurrences")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Maintenanceid != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "maintenanceid", runtime.ParamLocationQuery, *params.Maintenanceid); err != nil {
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

// NewDeleteMaintenanceOccurrencesIdRequest generates requests for DeleteMaintenanceOccurrencesId
func NewDeleteMaintenanceOccurrencesIdRequest(server string, id MaintenanceOccurrencesId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/maintenance.occurrences/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetMaintenanceOccurrencesIdRequest generates requests for GetMaintenanceOccurrencesId
func NewGetMaintenanceOccurrencesIdRequest(server string, id MaintenanceOccurrencesId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/maintenance.occurrences/%s", pathParam0)
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

// NewPutMaintenanceOccurrencesIdRequest calls the generic PutMaintenanceOccurrencesId builder with application/json body
func NewPutMaintenanceOccurrencesIdRequest(server string, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutMaintenanceOccurrencesIdRequestWithBody(server, id, "application/json", bodyReader)
}

// NewPutMaintenanceOccurrencesIdRequestWithFormdataBody calls the generic PutMaintenanceOccurrencesId builder with application/x-www-form-urlencoded body
func NewPutMaintenanceOccurrencesIdRequestWithFormdataBody(server string, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewPutMaintenanceOccurrencesIdRequestWithBody(server, id, "application/x-www-form-urlencoded", bodyReader)
}

// NewPutMaintenanceOccurrencesIdRequestWithBody generates requests for PutMaintenanceOccurrencesId with any type of body
func NewPutMaintenanceOccurrencesIdRequestWithBody(server string, id MaintenanceOccurrencesId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "id", runtime.ParamLocationPath, id)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/maintenance.occurrences/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

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
	// DeleteMaintenanceOccurrencesWithResponse request
	DeleteMaintenanceOccurrencesWithResponse(ctx context.Context, params *DeleteMaintenanceOccurrencesParams, reqEditors ...RequestEditorFn) (*DeleteMaintenanceOccurrencesResponse, error)

	// GetMaintenanceOccurrencesWithResponse request
	GetMaintenanceOccurrencesWithResponse(ctx context.Context, params *GetMaintenanceOccurrencesParams, reqEditors ...RequestEditorFn) (*GetMaintenanceOccurrencesResponse, error)

	// DeleteMaintenanceOccurrencesIdWithResponse request
	DeleteMaintenanceOccurrencesIdWithResponse(ctx context.Context, id MaintenanceOccurrencesId, reqEditors ...RequestEditorFn) (*DeleteMaintenanceOccurrencesIdResponse, error)

	// GetMaintenanceOccurrencesIdWithResponse request
	GetMaintenanceOccurrencesIdWithResponse(ctx context.Context, id MaintenanceOccurrencesId, reqEditors ...RequestEditorFn) (*GetMaintenanceOccurrencesIdResponse, error)

	// PutMaintenanceOccurrencesIdWithBodyWithResponse request with any body
	PutMaintenanceOccurrencesIdWithBodyWithResponse(ctx context.Context, id MaintenanceOccurrencesId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutMaintenanceOccurrencesIdResponse, error)

	PutMaintenanceOccurrencesIdWithResponse(ctx context.Context, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutMaintenanceOccurrencesIdResponse, error)

	PutMaintenanceOccurrencesIdWithFormdataBodyWithResponse(ctx context.Context, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdFormdataRequestBody, reqEditors ...RequestEditorFn) (*PutMaintenanceOccurrencesIdResponse, error)
}

type DeleteMaintenanceOccurrencesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MaintenanceOccurrencesDeleteRespAttrs
}

// Status returns HTTPResponse.Status
func (r DeleteMaintenanceOccurrencesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteMaintenanceOccurrencesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMaintenanceOccurrencesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MaintenanceOccurrencesRespAttrs
}

// Status returns HTTPResponse.Status
func (r GetMaintenanceOccurrencesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetMaintenanceOccurrencesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteMaintenanceOccurrencesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MaintenanceOccurrencesIdDeleteRespAttrs
}

// Status returns HTTPResponse.Status
func (r DeleteMaintenanceOccurrencesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteMaintenanceOccurrencesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetMaintenanceOccurrencesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MaintenanceOccurrencesIdRespAttrs
}

// Status returns HTTPResponse.Status
func (r GetMaintenanceOccurrencesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetMaintenanceOccurrencesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutMaintenanceOccurrencesIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *MaintenanceOccurrencesIdPutRespAttrs
}

// Status returns HTTPResponse.Status
func (r PutMaintenanceOccurrencesIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutMaintenanceOccurrencesIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DeleteMaintenanceOccurrencesWithResponse request returning *DeleteMaintenanceOccurrencesResponse
func (c *ClientWithResponses) DeleteMaintenanceOccurrencesWithResponse(ctx context.Context, params *DeleteMaintenanceOccurrencesParams, reqEditors ...RequestEditorFn) (*DeleteMaintenanceOccurrencesResponse, error) {
	rsp, err := c.DeleteMaintenanceOccurrences(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteMaintenanceOccurrencesResponse(rsp)
}

// GetMaintenanceOccurrencesWithResponse request returning *GetMaintenanceOccurrencesResponse
func (c *ClientWithResponses) GetMaintenanceOccurrencesWithResponse(ctx context.Context, params *GetMaintenanceOccurrencesParams, reqEditors ...RequestEditorFn) (*GetMaintenanceOccurrencesResponse, error) {
	rsp, err := c.GetMaintenanceOccurrences(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMaintenanceOccurrencesResponse(rsp)
}

// DeleteMaintenanceOccurrencesIdWithResponse request returning *DeleteMaintenanceOccurrencesIdResponse
func (c *ClientWithResponses) DeleteMaintenanceOccurrencesIdWithResponse(ctx context.Context, id MaintenanceOccurrencesId, reqEditors ...RequestEditorFn) (*DeleteMaintenanceOccurrencesIdResponse, error) {
	rsp, err := c.DeleteMaintenanceOccurrencesId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteMaintenanceOccurrencesIdResponse(rsp)
}

// GetMaintenanceOccurrencesIdWithResponse request returning *GetMaintenanceOccurrencesIdResponse
func (c *ClientWithResponses) GetMaintenanceOccurrencesIdWithResponse(ctx context.Context, id MaintenanceOccurrencesId, reqEditors ...RequestEditorFn) (*GetMaintenanceOccurrencesIdResponse, error) {
	rsp, err := c.GetMaintenanceOccurrencesId(ctx, id, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetMaintenanceOccurrencesIdResponse(rsp)
}

// PutMaintenanceOccurrencesIdWithBodyWithResponse request with arbitrary body returning *PutMaintenanceOccurrencesIdResponse
func (c *ClientWithResponses) PutMaintenanceOccurrencesIdWithBodyWithResponse(ctx context.Context, id MaintenanceOccurrencesId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutMaintenanceOccurrencesIdResponse, error) {
	rsp, err := c.PutMaintenanceOccurrencesIdWithBody(ctx, id, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutMaintenanceOccurrencesIdResponse(rsp)
}

func (c *ClientWithResponses) PutMaintenanceOccurrencesIdWithResponse(ctx context.Context, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutMaintenanceOccurrencesIdResponse, error) {
	rsp, err := c.PutMaintenanceOccurrencesId(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutMaintenanceOccurrencesIdResponse(rsp)
}

func (c *ClientWithResponses) PutMaintenanceOccurrencesIdWithFormdataBodyWithResponse(ctx context.Context, id MaintenanceOccurrencesId, body PutMaintenanceOccurrencesIdFormdataRequestBody, reqEditors ...RequestEditorFn) (*PutMaintenanceOccurrencesIdResponse, error) {
	rsp, err := c.PutMaintenanceOccurrencesIdWithFormdataBody(ctx, id, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutMaintenanceOccurrencesIdResponse(rsp)
}

// ParseDeleteMaintenanceOccurrencesResponse parses an HTTP response from a DeleteMaintenanceOccurrencesWithResponse call
func ParseDeleteMaintenanceOccurrencesResponse(rsp *http.Response) (*DeleteMaintenanceOccurrencesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteMaintenanceOccurrencesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MaintenanceOccurrencesDeleteRespAttrs
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetMaintenanceOccurrencesResponse parses an HTTP response from a GetMaintenanceOccurrencesWithResponse call
func ParseGetMaintenanceOccurrencesResponse(rsp *http.Response) (*GetMaintenanceOccurrencesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetMaintenanceOccurrencesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MaintenanceOccurrencesRespAttrs
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteMaintenanceOccurrencesIdResponse parses an HTTP response from a DeleteMaintenanceOccurrencesIdWithResponse call
func ParseDeleteMaintenanceOccurrencesIdResponse(rsp *http.Response) (*DeleteMaintenanceOccurrencesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteMaintenanceOccurrencesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MaintenanceOccurrencesIdDeleteRespAttrs
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetMaintenanceOccurrencesIdResponse parses an HTTP response from a GetMaintenanceOccurrencesIdWithResponse call
func ParseGetMaintenanceOccurrencesIdResponse(rsp *http.Response) (*GetMaintenanceOccurrencesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetMaintenanceOccurrencesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MaintenanceOccurrencesIdRespAttrs
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePutMaintenanceOccurrencesIdResponse parses an HTTP response from a PutMaintenanceOccurrencesIdWithResponse call
func ParsePutMaintenanceOccurrencesIdResponse(rsp *http.Response) (*PutMaintenanceOccurrencesIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutMaintenanceOccurrencesIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest MaintenanceOccurrencesIdPutRespAttrs
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}