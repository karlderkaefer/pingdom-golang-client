// Package contacts provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package contacts

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

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for ContactTargetsOwner.
const (
	ContactTargetsOwnerFalse ContactTargetsOwner = false
	ContactTargetsOwnerTrue  ContactTargetsOwner = true
)

// Defines values for ContactTargetsPaused.
const (
	ContactTargetsPausedFalse ContactTargetsPaused = false
	ContactTargetsPausedTrue  ContactTargetsPaused = true
)

// Defines values for ContactTargetsType.
const (
	ContactTargetsTypeContact ContactTargetsType = "contact"
	ContactTargetsTypeUser    ContactTargetsType = "user"
)

// Defines values for CreateContactPaused.
const (
	CreateContactPausedFalse CreateContactPaused = false
	CreateContactPausedTrue  CreateContactPaused = true
)

// Defines values for UpdateContactPaused.
const (
	UpdateContactPausedFalse UpdateContactPaused = false
	UpdateContactPausedTrue  UpdateContactPaused = true
)

// AGCM defines model for AGCM.
type AGCM = []struct {
	// AgcmId Contact target's agcm id
	AgcmId *string `json:"agcm_id,omitempty"`

	// Severity Contact target's severity level
	Severity *string `json:"severity,omitempty"`
}

// APNS defines model for APNS.
type APNS = []struct {
	// ApnsDevice Contact target's apns
	ApnsDevice *string `json:"apns_device,omitempty"`

	// DeviceName Contact targets's device name
	DeviceName *string `json:"device_name,omitempty"`

	// Severity Contact target's severity level
	Severity *string `json:"severity,omitempty"`
}

// Contact defines model for Contact.
type Contact struct {
	Contact *ContactTargets `json:"contact,omitempty"`
}

// ContactTargets defines model for ContactTargets.
type ContactTargets struct {
	// Id Contact ID
	Id *int `json:"id,omitempty"`

	// Name Contact name
	Name                *string                             `json:"name,omitempty"`
	NotificationTargets *ContactTargets_NotificationTargets `json:"notification_targets,omitempty"`

	// Owner Indicates whether the contact is the owner of the organization
	Owner *ContactTargetsOwner `json:"owner,omitempty"`

	// Paused Describes whether alerts are paused for this contact
	Paused *ContactTargetsPaused `json:"paused,omitempty"`
	Teams  *[]struct {
		// Id Team ID
		Id *int `json:"id,omitempty"`

		// Name Name of the team
		Name *string `json:"name,omitempty"`
	} `json:"teams,omitempty"`

	// Type Type defines whether this is a user (login user) or a contact only
	Type *ContactTargetsType `json:"type,omitempty"`
}

// ContactTargets_NotificationTargets defines model for ContactTargets.NotificationTargets.
type ContactTargets_NotificationTargets struct {
	union json.RawMessage
}

// ContactTargetsOwner Indicates whether the contact is the owner of the organization
type ContactTargetsOwner bool

// ContactTargetsPaused Describes whether alerts are paused for this contact
type ContactTargetsPaused bool

// ContactTargetsType Type defines whether this is a user (login user) or a contact only
type ContactTargetsType string

// ContactsList defines model for ContactsList.
type ContactsList struct {
	// Contacts List of all contacts targets
	Contacts *[]ContactTargets `json:"contacts,omitempty"`
}

// CreateContact defines model for CreateContact.
type CreateContact struct {
	// Name Contact name
	Name                string                            `json:"name"`
	NotificationTargets CreateContact_NotificationTargets `json:"notification_targets"`

	// Paused Pause contact
	Paused *CreateContactPaused `json:"paused,omitempty"`
}

// CreateContact_NotificationTargets defines model for CreateContact.NotificationTargets.
type CreateContact_NotificationTargets struct {
	union json.RawMessage
}

// CreateContactPaused Pause contact
type CreateContactPaused bool

// Emails defines model for Emails.
type Emails = []struct {
	// Address Email address
	Address *string `json:"address,omitempty"`

	// Severity Contact target's severity level
	Severity *string `json:"severity,omitempty"`
}

// SMSes defines model for SMSes.
type SMSes = []struct {
	// CountryCode Country code
	CountryCode *string `json:"country_code,omitempty"`

	// Number Phone number
	Number *string `json:"number,omitempty"`

	// Provider Provider
	Provider *string `json:"provider,omitempty"`

	// Severity Contact target's severity level
	Severity *string `json:"severity,omitempty"`
}

// UpdateContact defines model for UpdateContact.
type UpdateContact struct {
	// Name Contact name
	Name                string                            `json:"name"`
	NotificationTargets UpdateContact_NotificationTargets `json:"notification_targets"`

	// Paused Pause contact
	Paused UpdateContactPaused `json:"paused"`
}

// UpdateContact_NotificationTargets defines model for UpdateContact.NotificationTargets.
type UpdateContact_NotificationTargets struct {
	union json.RawMessage
}

// UpdateContactPaused Pause contact
type UpdateContactPaused bool

// PostAlertingContactsJSONRequestBody defines body for PostAlertingContacts for application/json ContentType.
type PostAlertingContactsJSONRequestBody = CreateContact

// PutAlertingContactsContactidJSONRequestBody defines body for PutAlertingContactsContactid for application/json ContentType.
type PutAlertingContactsContactidJSONRequestBody = UpdateContact

// AsSMSes returns the union data inside the ContactTargets_NotificationTargets as a SMSes
func (t ContactTargets_NotificationTargets) AsSMSes() (SMSes, error) {
	var body SMSes
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSMSes overwrites any union data inside the ContactTargets_NotificationTargets as the provided SMSes
func (t *ContactTargets_NotificationTargets) FromSMSes(v SMSes) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSMSes performs a merge with any union data inside the ContactTargets_NotificationTargets, using the provided SMSes
func (t *ContactTargets_NotificationTargets) MergeSMSes(v SMSes) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEmails returns the union data inside the ContactTargets_NotificationTargets as a Emails
func (t ContactTargets_NotificationTargets) AsEmails() (Emails, error) {
	var body Emails
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEmails overwrites any union data inside the ContactTargets_NotificationTargets as the provided Emails
func (t *ContactTargets_NotificationTargets) FromEmails(v Emails) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEmails performs a merge with any union data inside the ContactTargets_NotificationTargets, using the provided Emails
func (t *ContactTargets_NotificationTargets) MergeEmails(v Emails) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAPNS returns the union data inside the ContactTargets_NotificationTargets as a APNS
func (t ContactTargets_NotificationTargets) AsAPNS() (APNS, error) {
	var body APNS
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAPNS overwrites any union data inside the ContactTargets_NotificationTargets as the provided APNS
func (t *ContactTargets_NotificationTargets) FromAPNS(v APNS) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAPNS performs a merge with any union data inside the ContactTargets_NotificationTargets, using the provided APNS
func (t *ContactTargets_NotificationTargets) MergeAPNS(v APNS) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAGCM returns the union data inside the ContactTargets_NotificationTargets as a AGCM
func (t ContactTargets_NotificationTargets) AsAGCM() (AGCM, error) {
	var body AGCM
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAGCM overwrites any union data inside the ContactTargets_NotificationTargets as the provided AGCM
func (t *ContactTargets_NotificationTargets) FromAGCM(v AGCM) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAGCM performs a merge with any union data inside the ContactTargets_NotificationTargets, using the provided AGCM
func (t *ContactTargets_NotificationTargets) MergeAGCM(v AGCM) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ContactTargets_NotificationTargets) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ContactTargets_NotificationTargets) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSMSes returns the union data inside the CreateContact_NotificationTargets as a SMSes
func (t CreateContact_NotificationTargets) AsSMSes() (SMSes, error) {
	var body SMSes
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSMSes overwrites any union data inside the CreateContact_NotificationTargets as the provided SMSes
func (t *CreateContact_NotificationTargets) FromSMSes(v SMSes) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSMSes performs a merge with any union data inside the CreateContact_NotificationTargets, using the provided SMSes
func (t *CreateContact_NotificationTargets) MergeSMSes(v SMSes) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEmails returns the union data inside the CreateContact_NotificationTargets as a Emails
func (t CreateContact_NotificationTargets) AsEmails() (Emails, error) {
	var body Emails
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEmails overwrites any union data inside the CreateContact_NotificationTargets as the provided Emails
func (t *CreateContact_NotificationTargets) FromEmails(v Emails) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEmails performs a merge with any union data inside the CreateContact_NotificationTargets, using the provided Emails
func (t *CreateContact_NotificationTargets) MergeEmails(v Emails) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t CreateContact_NotificationTargets) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CreateContact_NotificationTargets) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSMSes returns the union data inside the UpdateContact_NotificationTargets as a SMSes
func (t UpdateContact_NotificationTargets) AsSMSes() (SMSes, error) {
	var body SMSes
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSMSes overwrites any union data inside the UpdateContact_NotificationTargets as the provided SMSes
func (t *UpdateContact_NotificationTargets) FromSMSes(v SMSes) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSMSes performs a merge with any union data inside the UpdateContact_NotificationTargets, using the provided SMSes
func (t *UpdateContact_NotificationTargets) MergeSMSes(v SMSes) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEmails returns the union data inside the UpdateContact_NotificationTargets as a Emails
func (t UpdateContact_NotificationTargets) AsEmails() (Emails, error) {
	var body Emails
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEmails overwrites any union data inside the UpdateContact_NotificationTargets as the provided Emails
func (t *UpdateContact_NotificationTargets) FromEmails(v Emails) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEmails performs a merge with any union data inside the UpdateContact_NotificationTargets, using the provided Emails
func (t *UpdateContact_NotificationTargets) MergeEmails(v Emails) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t UpdateContact_NotificationTargets) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *UpdateContact_NotificationTargets) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
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
	// GetAlertingContacts request
	GetAlertingContacts(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostAlertingContactsWithBody request with any body
	PostAlertingContactsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostAlertingContacts(ctx context.Context, body PostAlertingContactsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteAlertingContactsContactid request
	DeleteAlertingContactsContactid(ctx context.Context, contactid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetAlertingContactsContactid request
	GetAlertingContactsContactid(ctx context.Context, contactid int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutAlertingContactsContactidWithBody request with any body
	PutAlertingContactsContactidWithBody(ctx context.Context, contactid int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutAlertingContactsContactid(ctx context.Context, contactid int, body PutAlertingContactsContactidJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetAlertingContacts(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAlertingContactsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostAlertingContactsWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostAlertingContactsRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostAlertingContacts(ctx context.Context, body PostAlertingContactsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostAlertingContactsRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteAlertingContactsContactid(ctx context.Context, contactid int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteAlertingContactsContactidRequest(c.Server, contactid)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetAlertingContactsContactid(ctx context.Context, contactid int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetAlertingContactsContactidRequest(c.Server, contactid)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutAlertingContactsContactidWithBody(ctx context.Context, contactid int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutAlertingContactsContactidRequestWithBody(c.Server, contactid, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutAlertingContactsContactid(ctx context.Context, contactid int, body PutAlertingContactsContactidJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutAlertingContactsContactidRequest(c.Server, contactid, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetAlertingContactsRequest generates requests for GetAlertingContacts
func NewGetAlertingContactsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/alerting/contacts")
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

// NewPostAlertingContactsRequest calls the generic PostAlertingContacts builder with application/json body
func NewPostAlertingContactsRequest(server string, body PostAlertingContactsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostAlertingContactsRequestWithBody(server, "application/json", bodyReader)
}

// NewPostAlertingContactsRequestWithBody generates requests for PostAlertingContacts with any type of body
func NewPostAlertingContactsRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/alerting/contacts")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteAlertingContactsContactidRequest generates requests for DeleteAlertingContactsContactid
func NewDeleteAlertingContactsContactidRequest(server string, contactid int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "contactid", runtime.ParamLocationPath, contactid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/alerting/contacts/%s", pathParam0)
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

// NewGetAlertingContactsContactidRequest generates requests for GetAlertingContactsContactid
func NewGetAlertingContactsContactidRequest(server string, contactid int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "contactid", runtime.ParamLocationPath, contactid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/alerting/contacts/%s", pathParam0)
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

// NewPutAlertingContactsContactidRequest calls the generic PutAlertingContactsContactid builder with application/json body
func NewPutAlertingContactsContactidRequest(server string, contactid int, body PutAlertingContactsContactidJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutAlertingContactsContactidRequestWithBody(server, contactid, "application/json", bodyReader)
}

// NewPutAlertingContactsContactidRequestWithBody generates requests for PutAlertingContactsContactid with any type of body
func NewPutAlertingContactsContactidRequestWithBody(server string, contactid int, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "contactid", runtime.ParamLocationPath, contactid)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/alerting/contacts/%s", pathParam0)
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
	// GetAlertingContactsWithResponse request
	GetAlertingContactsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAlertingContactsResponse, error)

	// PostAlertingContactsWithBodyWithResponse request with any body
	PostAlertingContactsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostAlertingContactsResponse, error)

	PostAlertingContactsWithResponse(ctx context.Context, body PostAlertingContactsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostAlertingContactsResponse, error)

	// DeleteAlertingContactsContactidWithResponse request
	DeleteAlertingContactsContactidWithResponse(ctx context.Context, contactid int, reqEditors ...RequestEditorFn) (*DeleteAlertingContactsContactidResponse, error)

	// GetAlertingContactsContactidWithResponse request
	GetAlertingContactsContactidWithResponse(ctx context.Context, contactid int, reqEditors ...RequestEditorFn) (*GetAlertingContactsContactidResponse, error)

	// PutAlertingContactsContactidWithBodyWithResponse request with any body
	PutAlertingContactsContactidWithBodyWithResponse(ctx context.Context, contactid int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutAlertingContactsContactidResponse, error)

	PutAlertingContactsContactidWithResponse(ctx context.Context, contactid int, body PutAlertingContactsContactidJSONRequestBody, reqEditors ...RequestEditorFn) (*PutAlertingContactsContactidResponse, error)
}

type GetAlertingContactsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ContactsList
}

// Status returns HTTPResponse.Status
func (r GetAlertingContactsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAlertingContactsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostAlertingContactsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Contact *struct {
			// Id ID of the created contact
			Id *string `json:"id,omitempty"`
		} `json:"contact,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r PostAlertingContactsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostAlertingContactsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteAlertingContactsContactidResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		Message *string `json:"message,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r DeleteAlertingContactsContactidResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteAlertingContactsContactidResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetAlertingContactsContactidResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Contact
}

// Status returns HTTPResponse.Status
func (r GetAlertingContactsContactidResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetAlertingContactsContactidResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutAlertingContactsContactidResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Contact
}

// Status returns HTTPResponse.Status
func (r PutAlertingContactsContactidResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutAlertingContactsContactidResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetAlertingContactsWithResponse request returning *GetAlertingContactsResponse
func (c *ClientWithResponses) GetAlertingContactsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetAlertingContactsResponse, error) {
	rsp, err := c.GetAlertingContacts(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAlertingContactsResponse(rsp)
}

// PostAlertingContactsWithBodyWithResponse request with arbitrary body returning *PostAlertingContactsResponse
func (c *ClientWithResponses) PostAlertingContactsWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostAlertingContactsResponse, error) {
	rsp, err := c.PostAlertingContactsWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostAlertingContactsResponse(rsp)
}

func (c *ClientWithResponses) PostAlertingContactsWithResponse(ctx context.Context, body PostAlertingContactsJSONRequestBody, reqEditors ...RequestEditorFn) (*PostAlertingContactsResponse, error) {
	rsp, err := c.PostAlertingContacts(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostAlertingContactsResponse(rsp)
}

// DeleteAlertingContactsContactidWithResponse request returning *DeleteAlertingContactsContactidResponse
func (c *ClientWithResponses) DeleteAlertingContactsContactidWithResponse(ctx context.Context, contactid int, reqEditors ...RequestEditorFn) (*DeleteAlertingContactsContactidResponse, error) {
	rsp, err := c.DeleteAlertingContactsContactid(ctx, contactid, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteAlertingContactsContactidResponse(rsp)
}

// GetAlertingContactsContactidWithResponse request returning *GetAlertingContactsContactidResponse
func (c *ClientWithResponses) GetAlertingContactsContactidWithResponse(ctx context.Context, contactid int, reqEditors ...RequestEditorFn) (*GetAlertingContactsContactidResponse, error) {
	rsp, err := c.GetAlertingContactsContactid(ctx, contactid, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetAlertingContactsContactidResponse(rsp)
}

// PutAlertingContactsContactidWithBodyWithResponse request with arbitrary body returning *PutAlertingContactsContactidResponse
func (c *ClientWithResponses) PutAlertingContactsContactidWithBodyWithResponse(ctx context.Context, contactid int, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutAlertingContactsContactidResponse, error) {
	rsp, err := c.PutAlertingContactsContactidWithBody(ctx, contactid, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutAlertingContactsContactidResponse(rsp)
}

func (c *ClientWithResponses) PutAlertingContactsContactidWithResponse(ctx context.Context, contactid int, body PutAlertingContactsContactidJSONRequestBody, reqEditors ...RequestEditorFn) (*PutAlertingContactsContactidResponse, error) {
	rsp, err := c.PutAlertingContactsContactid(ctx, contactid, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutAlertingContactsContactidResponse(rsp)
}

// ParseGetAlertingContactsResponse parses an HTTP response from a GetAlertingContactsWithResponse call
func ParseGetAlertingContactsResponse(rsp *http.Response) (*GetAlertingContactsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAlertingContactsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContactsList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostAlertingContactsResponse parses an HTTP response from a PostAlertingContactsWithResponse call
func ParsePostAlertingContactsResponse(rsp *http.Response) (*PostAlertingContactsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostAlertingContactsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Contact *struct {
				// Id ID of the created contact
				Id *string `json:"id,omitempty"`
			} `json:"contact,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseDeleteAlertingContactsContactidResponse parses an HTTP response from a DeleteAlertingContactsContactidWithResponse call
func ParseDeleteAlertingContactsContactidResponse(rsp *http.Response) (*DeleteAlertingContactsContactidResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteAlertingContactsContactidResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			Message *string `json:"message,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetAlertingContactsContactidResponse parses an HTTP response from a GetAlertingContactsContactidWithResponse call
func ParseGetAlertingContactsContactidResponse(rsp *http.Response) (*GetAlertingContactsContactidResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetAlertingContactsContactidResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Contact
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePutAlertingContactsContactidResponse parses an HTTP response from a PutAlertingContactsContactidWithResponse call
func ParsePutAlertingContactsContactidResponse(rsp *http.Response) (*PutAlertingContactsContactidResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutAlertingContactsContactidResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Contact
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}
