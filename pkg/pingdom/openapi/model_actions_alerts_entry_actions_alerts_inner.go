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

// checks if the ActionsAlertsEntryActionsAlertsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionsAlertsEntryActionsAlertsInner{}

// ActionsAlertsEntryActionsAlertsInner struct for ActionsAlertsEntryActionsAlertsInner
type ActionsAlertsEntryActionsAlertsInner struct {
	// Name of alerted user
	Username *string `json:"username,omitempty"`
	// Identifier of alerted user
	Userid *string `json:"userid,omitempty"`
	// Identifier of alerted user
	Checkid *string `json:"checkid,omitempty"`
	// Time of alert generation. Format UNIX time
	Time *string `json:"time,omitempty"`
	// Alert medium. apns - iphone, agcm - android
	Via *string `json:"via,omitempty"`
	// Alert status
	Status *string `json:"status,omitempty"`
	// Short description of message
	Messageshort *string `json:"messageshort,omitempty"`
	// Full message body
	Messagefull *string `json:"messagefull,omitempty"`
	// Target address, phone number etc
	Sentto *string `json:"sentto,omitempty"`
	// True if your account was charged for this message
	Charged *string `json:"charged,omitempty"`
}

// NewActionsAlertsEntryActionsAlertsInner instantiates a new ActionsAlertsEntryActionsAlertsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionsAlertsEntryActionsAlertsInner() *ActionsAlertsEntryActionsAlertsInner {
	this := ActionsAlertsEntryActionsAlertsInner{}
	return &this
}

// NewActionsAlertsEntryActionsAlertsInnerWithDefaults instantiates a new ActionsAlertsEntryActionsAlertsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsAlertsEntryActionsAlertsInnerWithDefaults() *ActionsAlertsEntryActionsAlertsInner {
	this := ActionsAlertsEntryActionsAlertsInner{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetUsername(v string) {
	o.Username = &v
}

// GetUserid returns the Userid field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetUserid() string {
	if o == nil || IsNil(o.Userid) {
		var ret string
		return ret
	}
	return *o.Userid
}

// GetUseridOk returns a tuple with the Userid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetUseridOk() (*string, bool) {
	if o == nil || IsNil(o.Userid) {
		return nil, false
	}
	return o.Userid, true
}

// HasUserid returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasUserid() bool {
	if o != nil && !IsNil(o.Userid) {
		return true
	}

	return false
}

// SetUserid gets a reference to the given string and assigns it to the Userid field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetUserid(v string) {
	o.Userid = &v
}

// GetCheckid returns the Checkid field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetCheckid() string {
	if o == nil || IsNil(o.Checkid) {
		var ret string
		return ret
	}
	return *o.Checkid
}

// GetCheckidOk returns a tuple with the Checkid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetCheckidOk() (*string, bool) {
	if o == nil || IsNil(o.Checkid) {
		return nil, false
	}
	return o.Checkid, true
}

// HasCheckid returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasCheckid() bool {
	if o != nil && !IsNil(o.Checkid) {
		return true
	}

	return false
}

// SetCheckid gets a reference to the given string and assigns it to the Checkid field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetCheckid(v string) {
	o.Checkid = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetTime(v string) {
	o.Time = &v
}

// GetVia returns the Via field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetVia() string {
	if o == nil || IsNil(o.Via) {
		var ret string
		return ret
	}
	return *o.Via
}

// GetViaOk returns a tuple with the Via field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetViaOk() (*string, bool) {
	if o == nil || IsNil(o.Via) {
		return nil, false
	}
	return o.Via, true
}

// HasVia returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasVia() bool {
	if o != nil && !IsNil(o.Via) {
		return true
	}

	return false
}

// SetVia gets a reference to the given string and assigns it to the Via field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetVia(v string) {
	o.Via = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetStatus(v string) {
	o.Status = &v
}

// GetMessageshort returns the Messageshort field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetMessageshort() string {
	if o == nil || IsNil(o.Messageshort) {
		var ret string
		return ret
	}
	return *o.Messageshort
}

// GetMessageshortOk returns a tuple with the Messageshort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetMessageshortOk() (*string, bool) {
	if o == nil || IsNil(o.Messageshort) {
		return nil, false
	}
	return o.Messageshort, true
}

// HasMessageshort returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasMessageshort() bool {
	if o != nil && !IsNil(o.Messageshort) {
		return true
	}

	return false
}

// SetMessageshort gets a reference to the given string and assigns it to the Messageshort field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetMessageshort(v string) {
	o.Messageshort = &v
}

// GetMessagefull returns the Messagefull field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetMessagefull() string {
	if o == nil || IsNil(o.Messagefull) {
		var ret string
		return ret
	}
	return *o.Messagefull
}

// GetMessagefullOk returns a tuple with the Messagefull field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetMessagefullOk() (*string, bool) {
	if o == nil || IsNil(o.Messagefull) {
		return nil, false
	}
	return o.Messagefull, true
}

// HasMessagefull returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasMessagefull() bool {
	if o != nil && !IsNil(o.Messagefull) {
		return true
	}

	return false
}

// SetMessagefull gets a reference to the given string and assigns it to the Messagefull field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetMessagefull(v string) {
	o.Messagefull = &v
}

// GetSentto returns the Sentto field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetSentto() string {
	if o == nil || IsNil(o.Sentto) {
		var ret string
		return ret
	}
	return *o.Sentto
}

// GetSenttoOk returns a tuple with the Sentto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetSenttoOk() (*string, bool) {
	if o == nil || IsNil(o.Sentto) {
		return nil, false
	}
	return o.Sentto, true
}

// HasSentto returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasSentto() bool {
	if o != nil && !IsNil(o.Sentto) {
		return true
	}

	return false
}

// SetSentto gets a reference to the given string and assigns it to the Sentto field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetSentto(v string) {
	o.Sentto = &v
}

// GetCharged returns the Charged field value if set, zero value otherwise.
func (o *ActionsAlertsEntryActionsAlertsInner) GetCharged() string {
	if o == nil || IsNil(o.Charged) {
		var ret string
		return ret
	}
	return *o.Charged
}

// GetChargedOk returns a tuple with the Charged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) GetChargedOk() (*string, bool) {
	if o == nil || IsNil(o.Charged) {
		return nil, false
	}
	return o.Charged, true
}

// HasCharged returns a boolean if a field has been set.
func (o *ActionsAlertsEntryActionsAlertsInner) HasCharged() bool {
	if o != nil && !IsNil(o.Charged) {
		return true
	}

	return false
}

// SetCharged gets a reference to the given string and assigns it to the Charged field.
func (o *ActionsAlertsEntryActionsAlertsInner) SetCharged(v string) {
	o.Charged = &v
}

func (o ActionsAlertsEntryActionsAlertsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionsAlertsEntryActionsAlertsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Userid) {
		toSerialize["userid"] = o.Userid
	}
	if !IsNil(o.Checkid) {
		toSerialize["checkid"] = o.Checkid
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Via) {
		toSerialize["via"] = o.Via
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Messageshort) {
		toSerialize["messageshort"] = o.Messageshort
	}
	if !IsNil(o.Messagefull) {
		toSerialize["messagefull"] = o.Messagefull
	}
	if !IsNil(o.Sentto) {
		toSerialize["sentto"] = o.Sentto
	}
	if !IsNil(o.Charged) {
		toSerialize["charged"] = o.Charged
	}
	return toSerialize, nil
}

type NullableActionsAlertsEntryActionsAlertsInner struct {
	value *ActionsAlertsEntryActionsAlertsInner
	isSet bool
}

func (v NullableActionsAlertsEntryActionsAlertsInner) Get() *ActionsAlertsEntryActionsAlertsInner {
	return v.value
}

func (v *NullableActionsAlertsEntryActionsAlertsInner) Set(val *ActionsAlertsEntryActionsAlertsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableActionsAlertsEntryActionsAlertsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableActionsAlertsEntryActionsAlertsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionsAlertsEntryActionsAlertsInner(val *ActionsAlertsEntryActionsAlertsInner) *NullableActionsAlertsEntryActionsAlertsInner {
	return &NullableActionsAlertsEntryActionsAlertsInner{value: val, isSet: true}
}

func (v NullableActionsAlertsEntryActionsAlertsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionsAlertsEntryActionsAlertsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


