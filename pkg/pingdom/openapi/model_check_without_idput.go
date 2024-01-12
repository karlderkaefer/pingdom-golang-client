/*
Pingdom Public API

# Welcome to the Pingdom API! The Pingdom API is a way for you to automate your interaction with the Pingdom system. With the API, you can create your own scripts or applications with most of the functionality you can find inside the Pingdom control panel.  The Pingdom API is RESTful and HTTP-based. Basically, this means that the communication is made through normal HTTP requests.  # Authentication Authentication is needed in order to use the Pingdom API, and for this a Pingdom API Token is required. You generate your Pingdom API Token inside My Pingdom. The Pingdom API Token has a property called “Access level” to define its permissions. All operations that create or modify something (e.g. checks) need the Read/Write permission. If you only need to read data using the API token, we recommend to set the access level to “Read access”.  The authentication method for using tokens is HTTP Bearer Authentication (encrypted over HTTPS). This means that you will provide your token every time you make a request. No sessions are used.  Request ``` GET /checks HTTP/1.1 Host: api.pingdom.com Authorization: Bearer ofOhK18Ca6w4S_XmInGv0QPkqly-rbRBBoHsp_2FEH5QnIbH0VZhRPO3tlvrjMIKQ36VapX ```  Response ``` HTTP/1.1 200 OK Content-Length: 13 Content-Type: application/json {\"checks\":[]} ```  ## Basic Auth For compatibility reasons, the Pingdom API allows to use HTTP Basic Authentication with tokens. In cases where this is necessary, input the API token as the username and leave the password field empty.  An example request of how that would look like with the following API token: ofOhK18Ca6w4S_XmInGv0QPkqly-rbRBBoHsp_2FEH5QnIbH0VZhRPO3tlvrjMIKQ36VapX  ``` GET /checks HTTP/1.1 Host: api.pingdom.com Authorization: Basic b2ZPaEsxOENhNnc0U19YbUluR3YwUVBrcWx5LXJiUkJCb0hzcF8yRkVINVFuSWJIMFZaaFJQTzN0bHZyak1JS1EzNlZhcFg6 ```  # Server Address The base server address is: https://api.pingdom.com  Please note that HTTPS is required. You will not be able to connect through unencrypted HTTP.  # Providing Parameters GET requests should provide their parameters as a query string, part of the URL.  POST, PUT and DELETE requests should provide their parameters as a JSON. This should be part of the request body. Remember to add the proper content type header to the request: `Content-Type: application/json`.  We still support providing parameters as a query string for POST, PUT and DELETE requests, but we recommend using JSON going forward. If you are using query strings, they should be part of the body, URL or a combination. The encoding of the query string should be standard URL-encoding, as provided by various programming libraries.  When using `requests` library for Python, use `json` parameter instead of `data`. Due to the inner mechanisms of requests.post() etc. some endpoints may return responses not conforming to the documentation when dealing with `data` body.  # HTTP/1.1 Status Code Definitions The HTTP status code returned by a successful API request is defined in the documentation for that method. Usually, this will be 200 OK.  If something goes wrong, other codes may be returned. The API uses standard HTTP/1.1 status codes defined by [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html).  # JSON Responses All responses are sent JSON-encoded. The specific responses (successful ones) are described in the documentation section for each method.  However, if something goes wrong, our standard JSON error message (together with an appropriate status code) follows this format: ``` {     \"error\": {         \"statuscode\": 403,         \"statusdesc\": \"Forbidden\",         \"errormessage\":\" Something went wrong! This string describes what happened.\"     } } ``` See http://en.wikipedia.org/wiki/Json for more information on JSON.  Please note that all attributes of a method response are not always present. A client application should never assume that a certain attribute is present in a response.  # Limits The Pingdom API has usage limits to avoid individual rampant applications degrading the overall user experience. There are two layers of limits, the first cover a shorter period of time and the second a longer period. This enables you to \"burst\" requests for shorter periods. There are two HTTP headers in every response describing your limits status.  The response headers are: * **Req-Limit-Short** * **Req-Limit-Long**  An example of the values of these headers: * **Req-Limit-Short: Remaining: 394 Time until reset: 3589** * **Req-Limit-Long: Remaining: 71994 Time until reset: 2591989**  In this case, we can see that the user has 394 requests left until the short limit is reached. In 3589 seconds, the short limit will be reset. In similar fashion, the long limit has 71994 requests left, and will be reset in 2591989 seconds.  If limits are exceeded, an HTTP 429 error code with the message \"Request limit exceeded, try again later\" is sent back.  # gzip Responses can be gzip-encoded on demand. This is nice if your bandwidth is limited, or if big results are causing performance issues.  To enable gzip, simply add the following header to your request:  Accept-Encoding: gzip  # Best Practices ## Use caching If you are building a web page using the Pingdom API, we recommend that you do all API request on the server side, and if possible cache them. If you get any substantial traffic, you do not want to call the API each time you get a page hit, since this may cause you to hit the request limit faster than expected. In general, whenever you can cache data, do so.  ## Send your user credentials in a preemptive manner Some HTTP clients omit the authentication header, and make a second request with the header when they get a 401 Unauthorized response. Please make sure you send the credentials directly, to avoid unnecessary requests.  ## Use common sense Should be simple enough. For example, don't check for the status of a check every other second. The highest check resolution is one minute. Checking more often than that won't give you much of an advantage.  ## The Internet is unreliable Networks in general are unreliable, and particularly one as large and complex as the Internet. Your application should not assume it will get an answer. There may be timeouts.  # PHP Code Example **\"This is too much to read. I just want to get started right now! Give me a simple example!\"**  Here is a short example of how you can use the API with PHP. You need the cURL extension for PHP.  The example prints the current status of all your checks. This sample obviously focuses on Pingdom API code and does not worry about any potential problems connecting to the API, but your code should.  Code: ```php <?php     // Init cURL     $curl = curl_init();     // Set target URL     curl_setopt($curl, CURLOPT_URL, \"https://api.pingdom.com/api/3.1/checks\");     // Set the desired HTTP method (GET is default, see the documentation for each request)     curl_setopt($curl, CURLOPT_CUSTOMREQUEST, \"GET\");     // Add header with Bearer Authorization     curl_setopt($curl, CURLOPT_HTTPHEADER, array(\"Authorization: Bearer 907c762e069589c2cd2a229cdae7b8778caa9f07\"));     // Ask cURL to return the result as a string     curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);     // Execute the request and decode the json result into an associative array     $response = json_decode(curl_exec($curl), true);     // Check for errors returned by the API     if (isset($response['error'])) {         print \"Error: \" . $response['error']['errormessage'] . \"\\n\";         exit;     }     // Fetch the list of checks from the response     $checksList = $response['checks'];     // Print the names and statuses of all checks in the list     foreach ($checksList as $check) {         print $check['name'] . \" is \" . $check['status'] . \"\\n\";     } ?> ```  Example output: ``` Ubuntu Packages is up Google is up Pingdom is up My server 1 is down My server 2 is up ```  If you are running PHP on Windows, you need to be sure that you have installed the CA certificates for HTTPS/SSL to work correctly. Please see the cURL manual for more information. As a quick (but unsafe) workaround, you can add the following cURL option to ignore certificate validation.  ` curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, 0); `  # TMS Steps Vocabulary  There are two types of transaction checks: <ul>     <li><b>script</b>: the legacy TMS check created with predefined commands in the Pingdom UI or via the public API</li>     <li><b>recording</b>: the TMS check created by recording performed actions in WPM recorder.         More information about how to use it can be found in the         <a class=\"default-link\" href=\"https://documentation.solarwinds.com/en/success_center/wpm/Content/WPM-Use-the-WPM3-Recorder.htm\">             WPM recorder documentation</a>     </li> </ul>  ## Script transaction checks    ### Commands   Actions to be performed for the script TMS check    Step Name                                 | \"fn\"                  | Required \"args\"     | Remarks   ----------------------------------------- | --------------------- | --------------      | -------   Go to URL                                 | go_to                 | url                 | There has to be **exactly one** go_to step   Click                                     | click                 | element             | label, name or CSS selector   Fill in field                             | fill                  | input, value        | input: label, name or CSS selector, value: text   Check checkbox                            | check                 | checkbox            | label, name or CSS selector   Uncheck checkbox                          | uncheck               | checkbox            | label, name or CSS selector   Sleep for                                 | sleep                 | seconds             | number (e.g. 0.1)   Select radio button                       | select_radio          | radio               | name of the radio button   Select dropdown                           | select                | select, option      | select: label, name or CSS selector, option: content, value or CSS selector   Basic auth login with                     | basic_auth            | username, password  | username and password as text   Submit form                               | submit                | form                | name or CSS selector   Wait for element                          | wait_for_element      | element             | label, name or CSS selector   Wait for element to contain               | wait_for_contains     | element, value      | element: label, name or CSS selector, value: text    ### Validations   Verify the state of the page    Step Name                                 | \"fn\"                  | Required \"args\"     | Remarks   ----------------------------------------- | --------------------- | --------------      | -------   URL should be                             | url                   | url                 | url to be verified   Element should exist                      | exists                | element             | label, name or CSS selector   Element shouldn't exist                   | not_exists            | element             | label, name or CSS selector   Element should contain                    | contains              | element, value      | element: label, name or CSS selector, value: text   Element shouldn't containt                | not_contains          | element, value      | element: label, name or CSS selector, value: text   Text field should contain                 | field_contains        | input, value        | input: label, name or CSS selector, value: text   Text field shouldn't contain              | field_not_contains    | input, value        | input: label, name or CSS selector, value: text   Checkbox should be checked                | is_checked            | checkbox            | label, name or CSS selector   Checkbox shouldn't be checked             | is_not_checked        | checkbox            | label, name or CSS selector   Radio button should be selected           | radio_selected        | radio               | name of the radio button   Dropdown with name should be selected     | dropdown_selected     | select, option      | select: label, name or CSS selector, option: content, value or CSS selector   Dropdown with name shouldn't be selected  | dropdown_not_selected | select, option      | select: label, name or CSS selector, option: content, value or CSS selector    ### Example steps    ```   \"steps\": [   {     \"fn\": \"go_to\",     \"args\": {       \"url\": \"pingdom.com\"     }   },   {     \"fn\": \"click\",     \"args\": {       \"element\": \"START FREE TRIAL\"     }   },   {     \"fn\": \"url\",     \"args\": {       \"url\": \"https://www.pingdom.com/sign-up/\"     }   }   ]   ```  ## Recording transaction checks  Each of check steps contains: <ul>   <li><b>fn</b>: function name of the step</li>   <li><b>args</b>: function arguments</li>   <li><b>guid</b>: automatically generated identifier</li>   <li><b>contains_navigate</b>: recorder sets it on true if the step would trigger a page navigation</li> </ul>    ### Commands   Actions to be performed for the recording TMS check    Step Name                 | \"fn\"                      | Required \"args\"                 | Remarks   ------------------------- | ------------------------- | ------------------------------- | -------   Go to URL                 | wpm_go_to                 | url                             | Goes to the given url location   Click                     | wpm_click                 | element, offsetX, offsetY       | **element**: label, name or CSS selector,</br> **offsetX/Y**: exact position of a click in the element   Click in a exact location | wpm_click_xy              | element, x, y, scrollX, scrollY | **element**: label, name or CSS selector,</br> **x/y**: coordinates for the click (in viewport),</br> **scrollX/Y**: scrollbar position   Fill                      | wpm_fill                  | input, value                    | **input**: target element,</br> **value**: text to fill the target   Move mouse to element     | wpm_move_mouse_to_element | element, offsetX, offsetY       | **element**: target element,</br> **offsetX/Y**: exact position of the mouse in the element   Select dropdown           | wpm_select_dropdown       | select, option                  | **select**: target element (drop-down),</br> **option**: text of the option to select   Wait                      | wpm_wait                  | seconds                         | **seconds:** numbers of seconds to wait   Close tab                 | wpm_close_tab             | -                               | Closes the latest tab on the tab stack    ### Validations   Verify the state of the page    Step Name              | \"fn\"                     | Required \"args\"                                | Remarks   ---------------------- | ------------------------ | ---------------------------------------------- | -------   Contains text          | wpm_contains_timeout     | element, value, waitTime, useRegularExpression | **element**: label, name or CSS selector,</br> **value**: text to search for,</br> **waitTime**: time to wait for the value to appear,</br> **useRegularExpression**: use the value as a RegEx   Does not contains text | wpm_not_contains_timeout | element, value, waitTime, useRegularExpression | **element**: label, name or CSS selector,</br> **value**: text to search for,</br> **waitTime**: time to wait for the value to appear,</br> **useRegularExpression**: use the value as a RegEx    ### Metadata   Recording checks contain metadata which is automatically generated by the WPM recorder. Modify with caution! 

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CheckWithoutIDPUT type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckWithoutIDPUT{}

// CheckWithoutIDPUT CheckWithoutIDPUT is a struct describing a TMS check data common for updating a check
type CheckWithoutIDPUT struct {
	// Check status: active or inactive
	Active *bool `json:"active,omitempty"`
	// Contacts to alert
	ContactIds []int64 `json:"contact_ids,omitempty"`
	// Custom message that is part of the email and webhook alerts
	CustomMessage *string `json:"custom_message,omitempty"`
	// TMS test intervals in minutes. Allowed intervals: 5,10,20,60,720,1440. The interval you're allowed to set may vary depending on your current plan.
	Interval *int64 `json:"interval,omitempty"`
	// Name of the check
	Name *string `json:"name,omitempty"`
	// Name of the region where the check is executed. Supported regions: us-east, us-west, eu, au
	Region *string `json:"region,omitempty"`
	// Send notification when down X times
	SendNotificationWhenDown *int64 `json:"send_notification_when_down,omitempty"`
	// Check importance- how important are the alerts when the check fails. Allowed values: low, high
	SeverityLevel *string `json:"severity_level,omitempty"`
	// steps to be executed as part of the check
	Steps []Step `json:"steps,omitempty"`
	// Teams to alert
	TeamIds []int64 `json:"team_ids,omitempty"`
	Metadata *Metadata `json:"metadata,omitempty"`
	// List of tags for a check. The tag name may contain the characters 'A-Z', 'a-z', '0-9', '_' and '-'. The maximum length of a tag is 64 characters.
	Tags []string `json:"tags,omitempty"`
	// Integration identifiers as a list of integers.
	IntegrationIds []int32 `json:"integration_ids,omitempty"`
}

// NewCheckWithoutIDPUT instantiates a new CheckWithoutIDPUT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckWithoutIDPUT() *CheckWithoutIDPUT {
	this := CheckWithoutIDPUT{}
	return &this
}

// NewCheckWithoutIDPUTWithDefaults instantiates a new CheckWithoutIDPUT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckWithoutIDPUTWithDefaults() *CheckWithoutIDPUT {
	this := CheckWithoutIDPUT{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CheckWithoutIDPUT) SetActive(v bool) {
	o.Active = &v
}

// GetContactIds returns the ContactIds field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetContactIds() []int64 {
	if o == nil || IsNil(o.ContactIds) {
		var ret []int64
		return ret
	}
	return o.ContactIds
}

// GetContactIdsOk returns a tuple with the ContactIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetContactIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ContactIds) {
		return nil, false
	}
	return o.ContactIds, true
}

// HasContactIds returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasContactIds() bool {
	if o != nil && !IsNil(o.ContactIds) {
		return true
	}

	return false
}

// SetContactIds gets a reference to the given []int64 and assigns it to the ContactIds field.
func (o *CheckWithoutIDPUT) SetContactIds(v []int64) {
	o.ContactIds = v
}

// GetCustomMessage returns the CustomMessage field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetCustomMessage() string {
	if o == nil || IsNil(o.CustomMessage) {
		var ret string
		return ret
	}
	return *o.CustomMessage
}

// GetCustomMessageOk returns a tuple with the CustomMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetCustomMessageOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMessage) {
		return nil, false
	}
	return o.CustomMessage, true
}

// HasCustomMessage returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasCustomMessage() bool {
	if o != nil && !IsNil(o.CustomMessage) {
		return true
	}

	return false
}

// SetCustomMessage gets a reference to the given string and assigns it to the CustomMessage field.
func (o *CheckWithoutIDPUT) SetCustomMessage(v string) {
	o.CustomMessage = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetInterval() int64 {
	if o == nil || IsNil(o.Interval) {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetIntervalOk() (*int64, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *CheckWithoutIDPUT) SetInterval(v int64) {
	o.Interval = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CheckWithoutIDPUT) SetName(v string) {
	o.Name = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CheckWithoutIDPUT) SetRegion(v string) {
	o.Region = &v
}

// GetSendNotificationWhenDown returns the SendNotificationWhenDown field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetSendNotificationWhenDown() int64 {
	if o == nil || IsNil(o.SendNotificationWhenDown) {
		var ret int64
		return ret
	}
	return *o.SendNotificationWhenDown
}

// GetSendNotificationWhenDownOk returns a tuple with the SendNotificationWhenDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetSendNotificationWhenDownOk() (*int64, bool) {
	if o == nil || IsNil(o.SendNotificationWhenDown) {
		return nil, false
	}
	return o.SendNotificationWhenDown, true
}

// HasSendNotificationWhenDown returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasSendNotificationWhenDown() bool {
	if o != nil && !IsNil(o.SendNotificationWhenDown) {
		return true
	}

	return false
}

// SetSendNotificationWhenDown gets a reference to the given int64 and assigns it to the SendNotificationWhenDown field.
func (o *CheckWithoutIDPUT) SetSendNotificationWhenDown(v int64) {
	o.SendNotificationWhenDown = &v
}

// GetSeverityLevel returns the SeverityLevel field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetSeverityLevel() string {
	if o == nil || IsNil(o.SeverityLevel) {
		var ret string
		return ret
	}
	return *o.SeverityLevel
}

// GetSeverityLevelOk returns a tuple with the SeverityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetSeverityLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SeverityLevel) {
		return nil, false
	}
	return o.SeverityLevel, true
}

// HasSeverityLevel returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasSeverityLevel() bool {
	if o != nil && !IsNil(o.SeverityLevel) {
		return true
	}

	return false
}

// SetSeverityLevel gets a reference to the given string and assigns it to the SeverityLevel field.
func (o *CheckWithoutIDPUT) SetSeverityLevel(v string) {
	o.SeverityLevel = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetSteps() []Step {
	if o == nil || IsNil(o.Steps) {
		var ret []Step
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetStepsOk() ([]Step, bool) {
	if o == nil || IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasSteps() bool {
	if o != nil && !IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given []Step and assigns it to the Steps field.
func (o *CheckWithoutIDPUT) SetSteps(v []Step) {
	o.Steps = v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetTeamIds() []int64 {
	if o == nil || IsNil(o.TeamIds) {
		var ret []int64
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetTeamIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []int64 and assigns it to the TeamIds field.
func (o *CheckWithoutIDPUT) SetTeamIds(v []int64) {
	o.TeamIds = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata) {
		var ret Metadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetMetadataOk() (*Metadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Metadata and assigns it to the Metadata field.
func (o *CheckWithoutIDPUT) SetMetadata(v Metadata) {
	o.Metadata = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CheckWithoutIDPUT) SetTags(v []string) {
	o.Tags = v
}

// GetIntegrationIds returns the IntegrationIds field value if set, zero value otherwise.
func (o *CheckWithoutIDPUT) GetIntegrationIds() []int32 {
	if o == nil || IsNil(o.IntegrationIds) {
		var ret []int32
		return ret
	}
	return o.IntegrationIds
}

// GetIntegrationIdsOk returns a tuple with the IntegrationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckWithoutIDPUT) GetIntegrationIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.IntegrationIds) {
		return nil, false
	}
	return o.IntegrationIds, true
}

// HasIntegrationIds returns a boolean if a field has been set.
func (o *CheckWithoutIDPUT) HasIntegrationIds() bool {
	if o != nil && !IsNil(o.IntegrationIds) {
		return true
	}

	return false
}

// SetIntegrationIds gets a reference to the given []int32 and assigns it to the IntegrationIds field.
func (o *CheckWithoutIDPUT) SetIntegrationIds(v []int32) {
	o.IntegrationIds = v
}

func (o CheckWithoutIDPUT) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckWithoutIDPUT) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.ContactIds) {
		toSerialize["contact_ids"] = o.ContactIds
	}
	if !IsNil(o.CustomMessage) {
		toSerialize["custom_message"] = o.CustomMessage
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.SendNotificationWhenDown) {
		toSerialize["send_notification_when_down"] = o.SendNotificationWhenDown
	}
	if !IsNil(o.SeverityLevel) {
		toSerialize["severity_level"] = o.SeverityLevel
	}
	if !IsNil(o.Steps) {
		toSerialize["steps"] = o.Steps
	}
	if !IsNil(o.TeamIds) {
		toSerialize["team_ids"] = o.TeamIds
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.IntegrationIds) {
		toSerialize["integration_ids"] = o.IntegrationIds
	}
	return toSerialize, nil
}

type NullableCheckWithoutIDPUT struct {
	value *CheckWithoutIDPUT
	isSet bool
}

func (v NullableCheckWithoutIDPUT) Get() *CheckWithoutIDPUT {
	return v.value
}

func (v *NullableCheckWithoutIDPUT) Set(val *CheckWithoutIDPUT) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckWithoutIDPUT) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckWithoutIDPUT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckWithoutIDPUT(val *CheckWithoutIDPUT) *NullableCheckWithoutIDPUT {
	return &NullableCheckWithoutIDPUT{value: val, isSet: true}
}

func (v NullableCheckWithoutIDPUT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckWithoutIDPUT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


