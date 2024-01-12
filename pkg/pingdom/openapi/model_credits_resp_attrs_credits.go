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

// checks if the CreditsRespAttrsCredits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditsRespAttrsCredits{}

// CreditsRespAttrsCredits struct for CreditsRespAttrsCredits
type CreditsRespAttrsCredits struct {
	// Total number of check slots on this account
	Checklimit *int32 `json:"checklimit,omitempty"`
	// Free check slots available for new checks
	Availablechecks *int32 `json:"availablechecks,omitempty"`
	// Total number of default check slots
	Useddefault *int32 `json:"useddefault,omitempty"`
	// Total number of transaction check slots
	Usedtransaction *int32 `json:"usedtransaction,omitempty"`
	// SMS credits remaining on this account
	Availablesms *int32 `json:"availablesms,omitempty"`
	// SMS provider tests remaining on this account
	Availablesmstests *int32 `json:"availablesmstests,omitempty"`
	// Auto fill SMS
	Autofillsms *bool `json:"autofillsms,omitempty"`
	// The amount of sms to purchase when \"autofillsms_when_left\" is triggered
	AutofillsmsAmount *int32 `json:"autofillsms_amount,omitempty"`
	// The amount of sms left that is going to trigger purchase of sms
	AutofillsmsWhenLeft *int32 `json:"autofillsms_when_left,omitempty"`
	// The amount of overage SMSes that may be used, or null if SMS overage is not enabled.
	MaxSmsOverage *int32 `json:"max_sms_overage,omitempty"`
	// Open RUM site slots available
	Availablerumsites *int32 `json:"availablerumsites,omitempty"`
	// Number of used RUM sites
	Usedrumsites *int32 `json:"usedrumsites,omitempty"`
	// Number of maximum rum filters
	Maxrumfilters *int32 `json:"maxrumfilters,omitempty"`
	// Number of maximum pageviews per month
	Maxrumpageviews *int32 `json:"maxrumpageviews,omitempty"`
}

// NewCreditsRespAttrsCredits instantiates a new CreditsRespAttrsCredits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditsRespAttrsCredits() *CreditsRespAttrsCredits {
	this := CreditsRespAttrsCredits{}
	return &this
}

// NewCreditsRespAttrsCreditsWithDefaults instantiates a new CreditsRespAttrsCredits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditsRespAttrsCreditsWithDefaults() *CreditsRespAttrsCredits {
	this := CreditsRespAttrsCredits{}
	return &this
}

// GetChecklimit returns the Checklimit field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetChecklimit() int32 {
	if o == nil || IsNil(o.Checklimit) {
		var ret int32
		return ret
	}
	return *o.Checklimit
}

// GetChecklimitOk returns a tuple with the Checklimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetChecklimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Checklimit) {
		return nil, false
	}
	return o.Checklimit, true
}

// HasChecklimit returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasChecklimit() bool {
	if o != nil && !IsNil(o.Checklimit) {
		return true
	}

	return false
}

// SetChecklimit gets a reference to the given int32 and assigns it to the Checklimit field.
func (o *CreditsRespAttrsCredits) SetChecklimit(v int32) {
	o.Checklimit = &v
}

// GetAvailablechecks returns the Availablechecks field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetAvailablechecks() int32 {
	if o == nil || IsNil(o.Availablechecks) {
		var ret int32
		return ret
	}
	return *o.Availablechecks
}

// GetAvailablechecksOk returns a tuple with the Availablechecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetAvailablechecksOk() (*int32, bool) {
	if o == nil || IsNil(o.Availablechecks) {
		return nil, false
	}
	return o.Availablechecks, true
}

// HasAvailablechecks returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasAvailablechecks() bool {
	if o != nil && !IsNil(o.Availablechecks) {
		return true
	}

	return false
}

// SetAvailablechecks gets a reference to the given int32 and assigns it to the Availablechecks field.
func (o *CreditsRespAttrsCredits) SetAvailablechecks(v int32) {
	o.Availablechecks = &v
}

// GetUseddefault returns the Useddefault field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetUseddefault() int32 {
	if o == nil || IsNil(o.Useddefault) {
		var ret int32
		return ret
	}
	return *o.Useddefault
}

// GetUseddefaultOk returns a tuple with the Useddefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetUseddefaultOk() (*int32, bool) {
	if o == nil || IsNil(o.Useddefault) {
		return nil, false
	}
	return o.Useddefault, true
}

// HasUseddefault returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasUseddefault() bool {
	if o != nil && !IsNil(o.Useddefault) {
		return true
	}

	return false
}

// SetUseddefault gets a reference to the given int32 and assigns it to the Useddefault field.
func (o *CreditsRespAttrsCredits) SetUseddefault(v int32) {
	o.Useddefault = &v
}

// GetUsedtransaction returns the Usedtransaction field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetUsedtransaction() int32 {
	if o == nil || IsNil(o.Usedtransaction) {
		var ret int32
		return ret
	}
	return *o.Usedtransaction
}

// GetUsedtransactionOk returns a tuple with the Usedtransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetUsedtransactionOk() (*int32, bool) {
	if o == nil || IsNil(o.Usedtransaction) {
		return nil, false
	}
	return o.Usedtransaction, true
}

// HasUsedtransaction returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasUsedtransaction() bool {
	if o != nil && !IsNil(o.Usedtransaction) {
		return true
	}

	return false
}

// SetUsedtransaction gets a reference to the given int32 and assigns it to the Usedtransaction field.
func (o *CreditsRespAttrsCredits) SetUsedtransaction(v int32) {
	o.Usedtransaction = &v
}

// GetAvailablesms returns the Availablesms field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetAvailablesms() int32 {
	if o == nil || IsNil(o.Availablesms) {
		var ret int32
		return ret
	}
	return *o.Availablesms
}

// GetAvailablesmsOk returns a tuple with the Availablesms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetAvailablesmsOk() (*int32, bool) {
	if o == nil || IsNil(o.Availablesms) {
		return nil, false
	}
	return o.Availablesms, true
}

// HasAvailablesms returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasAvailablesms() bool {
	if o != nil && !IsNil(o.Availablesms) {
		return true
	}

	return false
}

// SetAvailablesms gets a reference to the given int32 and assigns it to the Availablesms field.
func (o *CreditsRespAttrsCredits) SetAvailablesms(v int32) {
	o.Availablesms = &v
}

// GetAvailablesmstests returns the Availablesmstests field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetAvailablesmstests() int32 {
	if o == nil || IsNil(o.Availablesmstests) {
		var ret int32
		return ret
	}
	return *o.Availablesmstests
}

// GetAvailablesmstestsOk returns a tuple with the Availablesmstests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetAvailablesmstestsOk() (*int32, bool) {
	if o == nil || IsNil(o.Availablesmstests) {
		return nil, false
	}
	return o.Availablesmstests, true
}

// HasAvailablesmstests returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasAvailablesmstests() bool {
	if o != nil && !IsNil(o.Availablesmstests) {
		return true
	}

	return false
}

// SetAvailablesmstests gets a reference to the given int32 and assigns it to the Availablesmstests field.
func (o *CreditsRespAttrsCredits) SetAvailablesmstests(v int32) {
	o.Availablesmstests = &v
}

// GetAutofillsms returns the Autofillsms field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetAutofillsms() bool {
	if o == nil || IsNil(o.Autofillsms) {
		var ret bool
		return ret
	}
	return *o.Autofillsms
}

// GetAutofillsmsOk returns a tuple with the Autofillsms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetAutofillsmsOk() (*bool, bool) {
	if o == nil || IsNil(o.Autofillsms) {
		return nil, false
	}
	return o.Autofillsms, true
}

// HasAutofillsms returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasAutofillsms() bool {
	if o != nil && !IsNil(o.Autofillsms) {
		return true
	}

	return false
}

// SetAutofillsms gets a reference to the given bool and assigns it to the Autofillsms field.
func (o *CreditsRespAttrsCredits) SetAutofillsms(v bool) {
	o.Autofillsms = &v
}

// GetAutofillsmsAmount returns the AutofillsmsAmount field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetAutofillsmsAmount() int32 {
	if o == nil || IsNil(o.AutofillsmsAmount) {
		var ret int32
		return ret
	}
	return *o.AutofillsmsAmount
}

// GetAutofillsmsAmountOk returns a tuple with the AutofillsmsAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetAutofillsmsAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.AutofillsmsAmount) {
		return nil, false
	}
	return o.AutofillsmsAmount, true
}

// HasAutofillsmsAmount returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasAutofillsmsAmount() bool {
	if o != nil && !IsNil(o.AutofillsmsAmount) {
		return true
	}

	return false
}

// SetAutofillsmsAmount gets a reference to the given int32 and assigns it to the AutofillsmsAmount field.
func (o *CreditsRespAttrsCredits) SetAutofillsmsAmount(v int32) {
	o.AutofillsmsAmount = &v
}

// GetAutofillsmsWhenLeft returns the AutofillsmsWhenLeft field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetAutofillsmsWhenLeft() int32 {
	if o == nil || IsNil(o.AutofillsmsWhenLeft) {
		var ret int32
		return ret
	}
	return *o.AutofillsmsWhenLeft
}

// GetAutofillsmsWhenLeftOk returns a tuple with the AutofillsmsWhenLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetAutofillsmsWhenLeftOk() (*int32, bool) {
	if o == nil || IsNil(o.AutofillsmsWhenLeft) {
		return nil, false
	}
	return o.AutofillsmsWhenLeft, true
}

// HasAutofillsmsWhenLeft returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasAutofillsmsWhenLeft() bool {
	if o != nil && !IsNil(o.AutofillsmsWhenLeft) {
		return true
	}

	return false
}

// SetAutofillsmsWhenLeft gets a reference to the given int32 and assigns it to the AutofillsmsWhenLeft field.
func (o *CreditsRespAttrsCredits) SetAutofillsmsWhenLeft(v int32) {
	o.AutofillsmsWhenLeft = &v
}

// GetMaxSmsOverage returns the MaxSmsOverage field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetMaxSmsOverage() int32 {
	if o == nil || IsNil(o.MaxSmsOverage) {
		var ret int32
		return ret
	}
	return *o.MaxSmsOverage
}

// GetMaxSmsOverageOk returns a tuple with the MaxSmsOverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetMaxSmsOverageOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSmsOverage) {
		return nil, false
	}
	return o.MaxSmsOverage, true
}

// HasMaxSmsOverage returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasMaxSmsOverage() bool {
	if o != nil && !IsNil(o.MaxSmsOverage) {
		return true
	}

	return false
}

// SetMaxSmsOverage gets a reference to the given int32 and assigns it to the MaxSmsOverage field.
func (o *CreditsRespAttrsCredits) SetMaxSmsOverage(v int32) {
	o.MaxSmsOverage = &v
}

// GetAvailablerumsites returns the Availablerumsites field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetAvailablerumsites() int32 {
	if o == nil || IsNil(o.Availablerumsites) {
		var ret int32
		return ret
	}
	return *o.Availablerumsites
}

// GetAvailablerumsitesOk returns a tuple with the Availablerumsites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetAvailablerumsitesOk() (*int32, bool) {
	if o == nil || IsNil(o.Availablerumsites) {
		return nil, false
	}
	return o.Availablerumsites, true
}

// HasAvailablerumsites returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasAvailablerumsites() bool {
	if o != nil && !IsNil(o.Availablerumsites) {
		return true
	}

	return false
}

// SetAvailablerumsites gets a reference to the given int32 and assigns it to the Availablerumsites field.
func (o *CreditsRespAttrsCredits) SetAvailablerumsites(v int32) {
	o.Availablerumsites = &v
}

// GetUsedrumsites returns the Usedrumsites field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetUsedrumsites() int32 {
	if o == nil || IsNil(o.Usedrumsites) {
		var ret int32
		return ret
	}
	return *o.Usedrumsites
}

// GetUsedrumsitesOk returns a tuple with the Usedrumsites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetUsedrumsitesOk() (*int32, bool) {
	if o == nil || IsNil(o.Usedrumsites) {
		return nil, false
	}
	return o.Usedrumsites, true
}

// HasUsedrumsites returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasUsedrumsites() bool {
	if o != nil && !IsNil(o.Usedrumsites) {
		return true
	}

	return false
}

// SetUsedrumsites gets a reference to the given int32 and assigns it to the Usedrumsites field.
func (o *CreditsRespAttrsCredits) SetUsedrumsites(v int32) {
	o.Usedrumsites = &v
}

// GetMaxrumfilters returns the Maxrumfilters field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetMaxrumfilters() int32 {
	if o == nil || IsNil(o.Maxrumfilters) {
		var ret int32
		return ret
	}
	return *o.Maxrumfilters
}

// GetMaxrumfiltersOk returns a tuple with the Maxrumfilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetMaxrumfiltersOk() (*int32, bool) {
	if o == nil || IsNil(o.Maxrumfilters) {
		return nil, false
	}
	return o.Maxrumfilters, true
}

// HasMaxrumfilters returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasMaxrumfilters() bool {
	if o != nil && !IsNil(o.Maxrumfilters) {
		return true
	}

	return false
}

// SetMaxrumfilters gets a reference to the given int32 and assigns it to the Maxrumfilters field.
func (o *CreditsRespAttrsCredits) SetMaxrumfilters(v int32) {
	o.Maxrumfilters = &v
}

// GetMaxrumpageviews returns the Maxrumpageviews field value if set, zero value otherwise.
func (o *CreditsRespAttrsCredits) GetMaxrumpageviews() int32 {
	if o == nil || IsNil(o.Maxrumpageviews) {
		var ret int32
		return ret
	}
	return *o.Maxrumpageviews
}

// GetMaxrumpageviewsOk returns a tuple with the Maxrumpageviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditsRespAttrsCredits) GetMaxrumpageviewsOk() (*int32, bool) {
	if o == nil || IsNil(o.Maxrumpageviews) {
		return nil, false
	}
	return o.Maxrumpageviews, true
}

// HasMaxrumpageviews returns a boolean if a field has been set.
func (o *CreditsRespAttrsCredits) HasMaxrumpageviews() bool {
	if o != nil && !IsNil(o.Maxrumpageviews) {
		return true
	}

	return false
}

// SetMaxrumpageviews gets a reference to the given int32 and assigns it to the Maxrumpageviews field.
func (o *CreditsRespAttrsCredits) SetMaxrumpageviews(v int32) {
	o.Maxrumpageviews = &v
}

func (o CreditsRespAttrsCredits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditsRespAttrsCredits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Checklimit) {
		toSerialize["checklimit"] = o.Checklimit
	}
	if !IsNil(o.Availablechecks) {
		toSerialize["availablechecks"] = o.Availablechecks
	}
	if !IsNil(o.Useddefault) {
		toSerialize["useddefault"] = o.Useddefault
	}
	if !IsNil(o.Usedtransaction) {
		toSerialize["usedtransaction"] = o.Usedtransaction
	}
	if !IsNil(o.Availablesms) {
		toSerialize["availablesms"] = o.Availablesms
	}
	if !IsNil(o.Availablesmstests) {
		toSerialize["availablesmstests"] = o.Availablesmstests
	}
	if !IsNil(o.Autofillsms) {
		toSerialize["autofillsms"] = o.Autofillsms
	}
	if !IsNil(o.AutofillsmsAmount) {
		toSerialize["autofillsms_amount"] = o.AutofillsmsAmount
	}
	if !IsNil(o.AutofillsmsWhenLeft) {
		toSerialize["autofillsms_when_left"] = o.AutofillsmsWhenLeft
	}
	if !IsNil(o.MaxSmsOverage) {
		toSerialize["max_sms_overage"] = o.MaxSmsOverage
	}
	if !IsNil(o.Availablerumsites) {
		toSerialize["availablerumsites"] = o.Availablerumsites
	}
	if !IsNil(o.Usedrumsites) {
		toSerialize["usedrumsites"] = o.Usedrumsites
	}
	if !IsNil(o.Maxrumfilters) {
		toSerialize["maxrumfilters"] = o.Maxrumfilters
	}
	if !IsNil(o.Maxrumpageviews) {
		toSerialize["maxrumpageviews"] = o.Maxrumpageviews
	}
	return toSerialize, nil
}

type NullableCreditsRespAttrsCredits struct {
	value *CreditsRespAttrsCredits
	isSet bool
}

func (v NullableCreditsRespAttrsCredits) Get() *CreditsRespAttrsCredits {
	return v.value
}

func (v *NullableCreditsRespAttrsCredits) Set(val *CreditsRespAttrsCredits) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditsRespAttrsCredits) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditsRespAttrsCredits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditsRespAttrsCredits(val *CreditsRespAttrsCredits) *NullableCreditsRespAttrsCredits {
	return &NullableCreditsRespAttrsCredits{value: val, isSet: true}
}

func (v NullableCreditsRespAttrsCredits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditsRespAttrsCredits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


