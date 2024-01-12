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

// checks if the DetailedCheckDnsCheck type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedCheckDnsCheck{}

// DetailedCheckDnsCheck struct for DetailedCheckDnsCheck
type DetailedCheckDnsCheck struct {
	Type *DetailedDnsAttributesType `json:"type,omitempty"`
	// Filters used for probe selections
	ProbeFilters []string `json:"probe_filters,omitempty"`
	Sendnotificationwhendown *int32 `json:"sendnotificationwhendown,omitempty"`
	Notifyagainevery *int32 `json:"notifyagainevery,omitempty"`
	Notifywhenbackup *bool `json:"notifywhenbackup,omitempty"`
	ResponsetimeThreshold *bool `json:"responsetime_threshold,omitempty"`
	CustomMessage *string `json:"custom_message,omitempty"`
	Integrationids []int32 `json:"integrationids,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Timestamp of last error (if any). Format is UNIX timestamp
	Lasterrortime *int32 `json:"lasterrortime,omitempty"`
	// Timestamp of last test (if any). Format is UNIX timestamp
	Lasttesttime *int32 `json:"lasttesttime,omitempty"`
	// Response time (in milliseconds) of last test.
	Lastresponsetime *int32 `json:"lastresponsetime,omitempty"`
	// Timestamp of start of last check down (if any). Format is UNIX timestamp.
	Lastdownstart *int32 `json:"lastdownstart,omitempty"`
	// Timestamp of end of last check down (if any). Format is UNIX timestamp. During a downtime it will be lasttesttime.
	Lastdownend *int32 `json:"lastdownend,omitempty"`
	Status *string `json:"status,omitempty"`
	// How often should the check be tested? (minutes)
	Resolution *int32 `json:"resolution,omitempty"`
	// Target host
	Hostname *string `json:"hostname,omitempty"`
	// Creating time. Format is UNIX timestamp
	Created *int32 `json:"created,omitempty"`
	// List of tags for check
	Tags []Tag `json:"tags,omitempty"`
	// Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`
}

// NewDetailedCheckDnsCheck instantiates a new DetailedCheckDnsCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedCheckDnsCheck() *DetailedCheckDnsCheck {
	this := DetailedCheckDnsCheck{}
	return &this
}

// NewDetailedCheckDnsCheckWithDefaults instantiates a new DetailedCheckDnsCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedCheckDnsCheckWithDefaults() *DetailedCheckDnsCheck {
	this := DetailedCheckDnsCheck{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetType() DetailedDnsAttributesType {
	if o == nil || IsNil(o.Type) {
		var ret DetailedDnsAttributesType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetTypeOk() (*DetailedDnsAttributesType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DetailedDnsAttributesType and assigns it to the Type field.
func (o *DetailedCheckDnsCheck) SetType(v DetailedDnsAttributesType) {
	o.Type = &v
}

// GetProbeFilters returns the ProbeFilters field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetProbeFilters() []string {
	if o == nil || IsNil(o.ProbeFilters) {
		var ret []string
		return ret
	}
	return o.ProbeFilters
}

// GetProbeFiltersOk returns a tuple with the ProbeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetProbeFiltersOk() ([]string, bool) {
	if o == nil || IsNil(o.ProbeFilters) {
		return nil, false
	}
	return o.ProbeFilters, true
}

// HasProbeFilters returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasProbeFilters() bool {
	if o != nil && !IsNil(o.ProbeFilters) {
		return true
	}

	return false
}

// SetProbeFilters gets a reference to the given []string and assigns it to the ProbeFilters field.
func (o *DetailedCheckDnsCheck) SetProbeFilters(v []string) {
	o.ProbeFilters = v
}

// GetSendnotificationwhendown returns the Sendnotificationwhendown field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetSendnotificationwhendown() int32 {
	if o == nil || IsNil(o.Sendnotificationwhendown) {
		var ret int32
		return ret
	}
	return *o.Sendnotificationwhendown
}

// GetSendnotificationwhendownOk returns a tuple with the Sendnotificationwhendown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetSendnotificationwhendownOk() (*int32, bool) {
	if o == nil || IsNil(o.Sendnotificationwhendown) {
		return nil, false
	}
	return o.Sendnotificationwhendown, true
}

// HasSendnotificationwhendown returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasSendnotificationwhendown() bool {
	if o != nil && !IsNil(o.Sendnotificationwhendown) {
		return true
	}

	return false
}

// SetSendnotificationwhendown gets a reference to the given int32 and assigns it to the Sendnotificationwhendown field.
func (o *DetailedCheckDnsCheck) SetSendnotificationwhendown(v int32) {
	o.Sendnotificationwhendown = &v
}

// GetNotifyagainevery returns the Notifyagainevery field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetNotifyagainevery() int32 {
	if o == nil || IsNil(o.Notifyagainevery) {
		var ret int32
		return ret
	}
	return *o.Notifyagainevery
}

// GetNotifyagaineveryOk returns a tuple with the Notifyagainevery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetNotifyagaineveryOk() (*int32, bool) {
	if o == nil || IsNil(o.Notifyagainevery) {
		return nil, false
	}
	return o.Notifyagainevery, true
}

// HasNotifyagainevery returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasNotifyagainevery() bool {
	if o != nil && !IsNil(o.Notifyagainevery) {
		return true
	}

	return false
}

// SetNotifyagainevery gets a reference to the given int32 and assigns it to the Notifyagainevery field.
func (o *DetailedCheckDnsCheck) SetNotifyagainevery(v int32) {
	o.Notifyagainevery = &v
}

// GetNotifywhenbackup returns the Notifywhenbackup field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetNotifywhenbackup() bool {
	if o == nil || IsNil(o.Notifywhenbackup) {
		var ret bool
		return ret
	}
	return *o.Notifywhenbackup
}

// GetNotifywhenbackupOk returns a tuple with the Notifywhenbackup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetNotifywhenbackupOk() (*bool, bool) {
	if o == nil || IsNil(o.Notifywhenbackup) {
		return nil, false
	}
	return o.Notifywhenbackup, true
}

// HasNotifywhenbackup returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasNotifywhenbackup() bool {
	if o != nil && !IsNil(o.Notifywhenbackup) {
		return true
	}

	return false
}

// SetNotifywhenbackup gets a reference to the given bool and assigns it to the Notifywhenbackup field.
func (o *DetailedCheckDnsCheck) SetNotifywhenbackup(v bool) {
	o.Notifywhenbackup = &v
}

// GetResponsetimeThreshold returns the ResponsetimeThreshold field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetResponsetimeThreshold() bool {
	if o == nil || IsNil(o.ResponsetimeThreshold) {
		var ret bool
		return ret
	}
	return *o.ResponsetimeThreshold
}

// GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetResponsetimeThresholdOk() (*bool, bool) {
	if o == nil || IsNil(o.ResponsetimeThreshold) {
		return nil, false
	}
	return o.ResponsetimeThreshold, true
}

// HasResponsetimeThreshold returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasResponsetimeThreshold() bool {
	if o != nil && !IsNil(o.ResponsetimeThreshold) {
		return true
	}

	return false
}

// SetResponsetimeThreshold gets a reference to the given bool and assigns it to the ResponsetimeThreshold field.
func (o *DetailedCheckDnsCheck) SetResponsetimeThreshold(v bool) {
	o.ResponsetimeThreshold = &v
}

// GetCustomMessage returns the CustomMessage field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetCustomMessage() string {
	if o == nil || IsNil(o.CustomMessage) {
		var ret string
		return ret
	}
	return *o.CustomMessage
}

// GetCustomMessageOk returns a tuple with the CustomMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetCustomMessageOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMessage) {
		return nil, false
	}
	return o.CustomMessage, true
}

// HasCustomMessage returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasCustomMessage() bool {
	if o != nil && !IsNil(o.CustomMessage) {
		return true
	}

	return false
}

// SetCustomMessage gets a reference to the given string and assigns it to the CustomMessage field.
func (o *DetailedCheckDnsCheck) SetCustomMessage(v string) {
	o.CustomMessage = &v
}

// GetIntegrationids returns the Integrationids field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetIntegrationids() []int32 {
	if o == nil || IsNil(o.Integrationids) {
		var ret []int32
		return ret
	}
	return o.Integrationids
}

// GetIntegrationidsOk returns a tuple with the Integrationids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetIntegrationidsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Integrationids) {
		return nil, false
	}
	return o.Integrationids, true
}

// HasIntegrationids returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasIntegrationids() bool {
	if o != nil && !IsNil(o.Integrationids) {
		return true
	}

	return false
}

// SetIntegrationids gets a reference to the given []int32 and assigns it to the Integrationids field.
func (o *DetailedCheckDnsCheck) SetIntegrationids(v []int32) {
	o.Integrationids = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DetailedCheckDnsCheck) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DetailedCheckDnsCheck) SetName(v string) {
	o.Name = &v
}

// GetLasterrortime returns the Lasterrortime field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetLasterrortime() int32 {
	if o == nil || IsNil(o.Lasterrortime) {
		var ret int32
		return ret
	}
	return *o.Lasterrortime
}

// GetLasterrortimeOk returns a tuple with the Lasterrortime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetLasterrortimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Lasterrortime) {
		return nil, false
	}
	return o.Lasterrortime, true
}

// HasLasterrortime returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasLasterrortime() bool {
	if o != nil && !IsNil(o.Lasterrortime) {
		return true
	}

	return false
}

// SetLasterrortime gets a reference to the given int32 and assigns it to the Lasterrortime field.
func (o *DetailedCheckDnsCheck) SetLasterrortime(v int32) {
	o.Lasterrortime = &v
}

// GetLasttesttime returns the Lasttesttime field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetLasttesttime() int32 {
	if o == nil || IsNil(o.Lasttesttime) {
		var ret int32
		return ret
	}
	return *o.Lasttesttime
}

// GetLasttesttimeOk returns a tuple with the Lasttesttime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetLasttesttimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Lasttesttime) {
		return nil, false
	}
	return o.Lasttesttime, true
}

// HasLasttesttime returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasLasttesttime() bool {
	if o != nil && !IsNil(o.Lasttesttime) {
		return true
	}

	return false
}

// SetLasttesttime gets a reference to the given int32 and assigns it to the Lasttesttime field.
func (o *DetailedCheckDnsCheck) SetLasttesttime(v int32) {
	o.Lasttesttime = &v
}

// GetLastresponsetime returns the Lastresponsetime field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetLastresponsetime() int32 {
	if o == nil || IsNil(o.Lastresponsetime) {
		var ret int32
		return ret
	}
	return *o.Lastresponsetime
}

// GetLastresponsetimeOk returns a tuple with the Lastresponsetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetLastresponsetimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Lastresponsetime) {
		return nil, false
	}
	return o.Lastresponsetime, true
}

// HasLastresponsetime returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasLastresponsetime() bool {
	if o != nil && !IsNil(o.Lastresponsetime) {
		return true
	}

	return false
}

// SetLastresponsetime gets a reference to the given int32 and assigns it to the Lastresponsetime field.
func (o *DetailedCheckDnsCheck) SetLastresponsetime(v int32) {
	o.Lastresponsetime = &v
}

// GetLastdownstart returns the Lastdownstart field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetLastdownstart() int32 {
	if o == nil || IsNil(o.Lastdownstart) {
		var ret int32
		return ret
	}
	return *o.Lastdownstart
}

// GetLastdownstartOk returns a tuple with the Lastdownstart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetLastdownstartOk() (*int32, bool) {
	if o == nil || IsNil(o.Lastdownstart) {
		return nil, false
	}
	return o.Lastdownstart, true
}

// HasLastdownstart returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasLastdownstart() bool {
	if o != nil && !IsNil(o.Lastdownstart) {
		return true
	}

	return false
}

// SetLastdownstart gets a reference to the given int32 and assigns it to the Lastdownstart field.
func (o *DetailedCheckDnsCheck) SetLastdownstart(v int32) {
	o.Lastdownstart = &v
}

// GetLastdownend returns the Lastdownend field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetLastdownend() int32 {
	if o == nil || IsNil(o.Lastdownend) {
		var ret int32
		return ret
	}
	return *o.Lastdownend
}

// GetLastdownendOk returns a tuple with the Lastdownend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetLastdownendOk() (*int32, bool) {
	if o == nil || IsNil(o.Lastdownend) {
		return nil, false
	}
	return o.Lastdownend, true
}

// HasLastdownend returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasLastdownend() bool {
	if o != nil && !IsNil(o.Lastdownend) {
		return true
	}

	return false
}

// SetLastdownend gets a reference to the given int32 and assigns it to the Lastdownend field.
func (o *DetailedCheckDnsCheck) SetLastdownend(v int32) {
	o.Lastdownend = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DetailedCheckDnsCheck) SetStatus(v string) {
	o.Status = &v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetResolution() int32 {
	if o == nil || IsNil(o.Resolution) {
		var ret int32
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetResolutionOk() (*int32, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given int32 and assigns it to the Resolution field.
func (o *DetailedCheckDnsCheck) SetResolution(v int32) {
	o.Resolution = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *DetailedCheckDnsCheck) SetHostname(v string) {
	o.Hostname = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetCreated() int32 {
	if o == nil || IsNil(o.Created) {
		var ret int32
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetCreatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int32 and assigns it to the Created field.
func (o *DetailedCheckDnsCheck) SetCreated(v int32) {
	o.Created = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *DetailedCheckDnsCheck) SetTags(v []Tag) {
	o.Tags = v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *DetailedCheckDnsCheck) GetIpv6() bool {
	if o == nil || IsNil(o.Ipv6) {
		var ret bool
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedCheckDnsCheck) GetIpv6Ok() (*bool, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *DetailedCheckDnsCheck) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given bool and assigns it to the Ipv6 field.
func (o *DetailedCheckDnsCheck) SetIpv6(v bool) {
	o.Ipv6 = &v
}

func (o DetailedCheckDnsCheck) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailedCheckDnsCheck) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ProbeFilters) {
		toSerialize["probe_filters"] = o.ProbeFilters
	}
	if !IsNil(o.Sendnotificationwhendown) {
		toSerialize["sendnotificationwhendown"] = o.Sendnotificationwhendown
	}
	if !IsNil(o.Notifyagainevery) {
		toSerialize["notifyagainevery"] = o.Notifyagainevery
	}
	if !IsNil(o.Notifywhenbackup) {
		toSerialize["notifywhenbackup"] = o.Notifywhenbackup
	}
	if !IsNil(o.ResponsetimeThreshold) {
		toSerialize["responsetime_threshold"] = o.ResponsetimeThreshold
	}
	if !IsNil(o.CustomMessage) {
		toSerialize["custom_message"] = o.CustomMessage
	}
	if !IsNil(o.Integrationids) {
		toSerialize["integrationids"] = o.Integrationids
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Lasterrortime) {
		toSerialize["lasterrortime"] = o.Lasterrortime
	}
	if !IsNil(o.Lasttesttime) {
		toSerialize["lasttesttime"] = o.Lasttesttime
	}
	if !IsNil(o.Lastresponsetime) {
		toSerialize["lastresponsetime"] = o.Lastresponsetime
	}
	if !IsNil(o.Lastdownstart) {
		toSerialize["lastdownstart"] = o.Lastdownstart
	}
	if !IsNil(o.Lastdownend) {
		toSerialize["lastdownend"] = o.Lastdownend
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return toSerialize, nil
}

type NullableDetailedCheckDnsCheck struct {
	value *DetailedCheckDnsCheck
	isSet bool
}

func (v NullableDetailedCheckDnsCheck) Get() *DetailedCheckDnsCheck {
	return v.value
}

func (v *NullableDetailedCheckDnsCheck) Set(val *DetailedCheckDnsCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedCheckDnsCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedCheckDnsCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedCheckDnsCheck(val *DetailedCheckDnsCheck) *NullableDetailedCheckDnsCheck {
	return &NullableDetailedCheckDnsCheck{value: val, isSet: true}
}

func (v NullableDetailedCheckDnsCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedCheckDnsCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


