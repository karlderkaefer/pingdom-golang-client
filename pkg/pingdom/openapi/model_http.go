/*
Pingdom Public API

# Welcome to the Pingdom API! The Pingdom API is a way for you to automate your interaction with the Pingdom system. With the API, you can create your own scripts or applications with most of the functionality you can find inside the Pingdom control panel.  The Pingdom API is RESTful and HTTP-based. Basically, this means that the communication is made through normal HTTP requests.  # Authentication Authentication is needed in order to use the Pingdom API, and for this a Pingdom API Token is required. You generate your Pingdom API Token inside My Pingdom. The Pingdom API Token has a property called “Access level” to define its permissions. All operations that create or modify something (e.g. checks) need the Read/Write permission. If you only need to read data using the API token, we recommend to set the access level to “Read access”.  The authentication method for using tokens is HTTP Bearer Authentication (encrypted over HTTPS). This means that you will provide your token every time you make a request. No sessions are used.  Request ``` GET /checks HTTP/1.1 Host: api.pingdom.com Authorization: Bearer ofOhK18Ca6w4S_XmInGv0QPkqly-rbRBBoHsp_2FEH5QnIbH0VZhRPO3tlvrjMIKQ36VapX ```  Response ``` HTTP/1.1 200 OK Content-Length: 13 Content-Type: application/json {\"checks\":[]} ```  ## Basic Auth For compatibility reasons, the Pingdom API allows to use HTTP Basic Authentication with tokens. In cases where this is necessary, input the API token as the username and leave the password field empty.  An example request of how that would look like with the following API token: ofOhK18Ca6w4S_XmInGv0QPkqly-rbRBBoHsp_2FEH5QnIbH0VZhRPO3tlvrjMIKQ36VapX  ``` GET /checks HTTP/1.1 Host: api.pingdom.com Authorization: Basic b2ZPaEsxOENhNnc0U19YbUluR3YwUVBrcWx5LXJiUkJCb0hzcF8yRkVINVFuSWJIMFZaaFJQTzN0bHZyak1JS1EzNlZhcFg6 ```  # Server Address The base server address is: https://api.pingdom.com  Please note that HTTPS is required. You will not be able to connect through unencrypted HTTP.  # Providing Parameters GET requests should provide their parameters as a query string, part of the URL.  POST, PUT and DELETE requests should provide their parameters as a JSON. This should be part of the request body. Remember to add the proper content type header to the request: `Content-Type: application/json`.  We still support providing parameters as a query string for POST, PUT and DELETE requests, but we recommend using JSON going forward. If you are using query strings, they should be part of the body, URL or a combination. The encoding of the query string should be standard URL-encoding, as provided by various programming libraries.  When using `requests` library for Python, use `json` parameter instead of `data`. Due to the inner mechanisms of requests.post() etc. some endpoints may return responses not conforming to the documentation when dealing with `data` body.  # HTTP/1.1 Status Code Definitions The HTTP status code returned by a successful API request is defined in the documentation for that method. Usually, this will be 200 OK.  If something goes wrong, other codes may be returned. The API uses standard HTTP/1.1 status codes defined by [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html).  # JSON Responses All responses are sent JSON-encoded. The specific responses (successful ones) are described in the documentation section for each method.  However, if something goes wrong, our standard JSON error message (together with an appropriate status code) follows this format: ``` {     \"error\": {         \"statuscode\": 403,         \"statusdesc\": \"Forbidden\",         \"errormessage\":\" Something went wrong! This string describes what happened.\"     } } ``` See http://en.wikipedia.org/wiki/Json for more information on JSON.  Please note that all attributes of a method response are not always present. A client application should never assume that a certain attribute is present in a response.  # Limits The Pingdom API has usage limits to avoid individual rampant applications degrading the overall user experience. There are two layers of limits, the first cover a shorter period of time and the second a longer period. This enables you to \"burst\" requests for shorter periods. There are two HTTP headers in every response describing your limits status.  The response headers are: * **Req-Limit-Short** * **Req-Limit-Long**  An example of the values of these headers: * **Req-Limit-Short: Remaining: 394 Time until reset: 3589** * **Req-Limit-Long: Remaining: 71994 Time until reset: 2591989**  In this case, we can see that the user has 394 requests left until the short limit is reached. In 3589 seconds, the short limit will be reset. In similar fashion, the long limit has 71994 requests left, and will be reset in 2591989 seconds.  If limits are exceeded, an HTTP 429 error code with the message \"Request limit exceeded, try again later\" is sent back.  # gzip Responses can be gzip-encoded on demand. This is nice if your bandwidth is limited, or if big results are causing performance issues.  To enable gzip, simply add the following header to your request:  Accept-Encoding: gzip  # Best Practices ## Use caching If you are building a web page using the Pingdom API, we recommend that you do all API request on the server side, and if possible cache them. If you get any substantial traffic, you do not want to call the API each time you get a page hit, since this may cause you to hit the request limit faster than expected. In general, whenever you can cache data, do so.  ## Send your user credentials in a preemptive manner Some HTTP clients omit the authentication header, and make a second request with the header when they get a 401 Unauthorized response. Please make sure you send the credentials directly, to avoid unnecessary requests.  ## Use common sense Should be simple enough. For example, don't check for the status of a check every other second. The highest check resolution is one minute. Checking more often than that won't give you much of an advantage.  ## The Internet is unreliable Networks in general are unreliable, and particularly one as large and complex as the Internet. Your application should not assume it will get an answer. There may be timeouts.  # PHP Code Example **\"This is too much to read. I just want to get started right now! Give me a simple example!\"**  Here is a short example of how you can use the API with PHP. You need the cURL extension for PHP.  The example prints the current status of all your checks. This sample obviously focuses on Pingdom API code and does not worry about any potential problems connecting to the API, but your code should.  Code: ```php <?php     // Init cURL     $curl = curl_init();     // Set target URL     curl_setopt($curl, CURLOPT_URL, \"https://api.pingdom.com/api/3.1/checks\");     // Set the desired HTTP method (GET is default, see the documentation for each request)     curl_setopt($curl, CURLOPT_CUSTOMREQUEST, \"GET\");     // Add header with Bearer Authorization     curl_setopt($curl, CURLOPT_HTTPHEADER, array(\"Authorization: Bearer 907c762e069589c2cd2a229cdae7b8778caa9f07\"));     // Ask cURL to return the result as a string     curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);     // Execute the request and decode the json result into an associative array     $response = json_decode(curl_exec($curl), true);     // Check for errors returned by the API     if (isset($response['error'])) {         print \"Error: \" . $response['error']['errormessage'] . \"\\n\";         exit;     }     // Fetch the list of checks from the response     $checksList = $response['checks'];     // Print the names and statuses of all checks in the list     foreach ($checksList as $check) {         print $check['name'] . \" is \" . $check['status'] . \"\\n\";     } ?> ```  Example output: ``` Ubuntu Packages is up Google is up Pingdom is up My server 1 is down My server 2 is up ```  If you are running PHP on Windows, you need to be sure that you have installed the CA certificates for HTTPS/SSL to work correctly. Please see the cURL manual for more information. As a quick (but unsafe) workaround, you can add the following cURL option to ignore certificate validation.  ` curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, 0); `  # TMS Steps Vocabulary  There are two types of transaction checks: <ul>     <li><b>script</b>: the legacy TMS check created with predefined commands in the Pingdom UI or via the public API</li>     <li><b>recording</b>: the TMS check created by recording performed actions in WPM recorder.         More information about how to use it can be found in the         <a class=\"default-link\" href=\"https://documentation.solarwinds.com/en/success_center/wpm/Content/WPM-Use-the-WPM3-Recorder.htm\">             WPM recorder documentation</a>     </li> </ul>  ## Script transaction checks    ### Commands   Actions to be performed for the script TMS check    Step Name                                 | \"fn\"                  | Required \"args\"     | Remarks   ----------------------------------------- | --------------------- | --------------      | -------   Go to URL                                 | go_to                 | url                 | There has to be **exactly one** go_to step   Click                                     | click                 | element             | label, name or CSS selector   Fill in field                             | fill                  | input, value        | input: label, name or CSS selector, value: text   Check checkbox                            | check                 | checkbox            | label, name or CSS selector   Uncheck checkbox                          | uncheck               | checkbox            | label, name or CSS selector   Sleep for                                 | sleep                 | seconds             | number (e.g. 0.1)   Select radio button                       | select_radio          | radio               | name of the radio button   Select dropdown                           | select                | select, option      | select: label, name or CSS selector, option: content, value or CSS selector   Basic auth login with                     | basic_auth            | username, password  | username and password as text   Submit form                               | submit                | form                | name or CSS selector   Wait for element                          | wait_for_element      | element             | label, name or CSS selector   Wait for element to contain               | wait_for_contains     | element, value      | element: label, name or CSS selector, value: text    ### Validations   Verify the state of the page    Step Name                                 | \"fn\"                  | Required \"args\"     | Remarks   ----------------------------------------- | --------------------- | --------------      | -------   URL should be                             | url                   | url                 | url to be verified   Element should exist                      | exists                | element             | label, name or CSS selector   Element shouldn't exist                   | not_exists            | element             | label, name or CSS selector   Element should contain                    | contains              | element, value      | element: label, name or CSS selector, value: text   Element shouldn't containt                | not_contains          | element, value      | element: label, name or CSS selector, value: text   Text field should contain                 | field_contains        | input, value        | input: label, name or CSS selector, value: text   Text field shouldn't contain              | field_not_contains    | input, value        | input: label, name or CSS selector, value: text   Checkbox should be checked                | is_checked            | checkbox            | label, name or CSS selector   Checkbox shouldn't be checked             | is_not_checked        | checkbox            | label, name or CSS selector   Radio button should be selected           | radio_selected        | radio               | name of the radio button   Dropdown with name should be selected     | dropdown_selected     | select, option      | select: label, name or CSS selector, option: content, value or CSS selector   Dropdown with name shouldn't be selected  | dropdown_not_selected | select, option      | select: label, name or CSS selector, option: content, value or CSS selector    ### Example steps    ```   \"steps\": [   {     \"fn\": \"go_to\",     \"args\": {       \"url\": \"pingdom.com\"     }   },   {     \"fn\": \"click\",     \"args\": {       \"element\": \"START FREE TRIAL\"     }   },   {     \"fn\": \"url\",     \"args\": {       \"url\": \"https://www.pingdom.com/sign-up/\"     }   }   ]   ```  ## Recording transaction checks  Each of check steps contains: <ul>   <li><b>fn</b>: function name of the step</li>   <li><b>args</b>: function arguments</li>   <li><b>guid</b>: automatically generated identifier</li>   <li><b>contains_navigate</b>: recorder sets it on true if the step would trigger a page navigation</li> </ul>    ### Commands   Actions to be performed for the recording TMS check    Step Name                 | \"fn\"                      | Required \"args\"                 | Remarks   ------------------------- | ------------------------- | ------------------------------- | -------   Go to URL                 | wpm_go_to                 | url                             | Goes to the given url location   Click                     | wpm_click                 | element, offsetX, offsetY       | **element**: label, name or CSS selector,</br> **offsetX/Y**: exact position of a click in the element   Click in a exact location | wpm_click_xy              | element, x, y, scrollX, scrollY | **element**: label, name or CSS selector,</br> **x/y**: coordinates for the click (in viewport),</br> **scrollX/Y**: scrollbar position   Fill                      | wpm_fill                  | input, value                    | **input**: target element,</br> **value**: text to fill the target   Move mouse to element     | wpm_move_mouse_to_element | element, offsetX, offsetY       | **element**: target element,</br> **offsetX/Y**: exact position of the mouse in the element   Select dropdown           | wpm_select_dropdown       | select, option                  | **select**: target element (drop-down),</br> **option**: text of the option to select   Wait                      | wpm_wait                  | seconds                         | **seconds:** numbers of seconds to wait   Close tab                 | wpm_close_tab             | -                               | Closes the latest tab on the tab stack    ### Validations   Verify the state of the page    Step Name              | \"fn\"                     | Required \"args\"                                | Remarks   ---------------------- | ------------------------ | ---------------------------------------------- | -------   Contains text          | wpm_contains_timeout     | element, value, waitTime, useRegularExpression | **element**: label, name or CSS selector,</br> **value**: text to search for,</br> **waitTime**: time to wait for the value to appear,</br> **useRegularExpression**: use the value as a RegEx   Does not contains text | wpm_not_contains_timeout | element, value, waitTime, useRegularExpression | **element**: label, name or CSS selector,</br> **value**: text to search for,</br> **waitTime**: time to wait for the value to appear,</br> **useRegularExpression**: use the value as a RegEx    ### Metadata   Recording checks contain metadata which is automatically generated by the WPM recorder. Modify with caution! 

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HTTP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HTTP{}

// HTTP struct for HTTP
type HTTP struct {
	// Target host
	Host string `json:"host"`
	Type string `json:"type"`
	// Probe identifier
	Probeid *int32 `json:"probeid,omitempty"`
	// Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are 'EU', 'NA', 'APAC' and 'LATAM'.
	ProbeFilters *int32 `json:"probe_filters,omitempty"`
	// Use ipv6 instead of ipv4
	Ipv6 *bool `json:"ipv6,omitempty"`
	// Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
	ResponsetimeThreshold *int32 `json:"responsetime_threshold,omitempty"`
	// (http specific) Target path on server
	Url *string `json:"url,omitempty"`
	// (http specific) Connection encryption
	Encryption *bool `json:"encryption,omitempty"`
	// (http specific) Target port
	Port *int32 `json:"port,omitempty"`
	// (http specific) Username and password for target HTTP authentication.
	Auth *string `json:"auth,omitempty"`
	// (http specific) Target site should contain this string
	Shouldcontain *string `json:"shouldcontain,omitempty"`
	// (http specific) Target site should NOT contain this string
	Shouldnotcontain *string `json:"shouldnotcontain,omitempty"`
	// (http specific) Data that should be posted to the web page, for example submission data for a sign-up or login form. The data needs to be formatted in the same way as a web browser would send it to the web server
	Postdata *string `json:"postdata,omitempty"`
	// (http specific) Custom HTTP header name. Replace {X} with a number unique for each header argument.
	RequestheaderX *string `json:"requestheader{X},omitempty"`
	// (http specific) Treat target site as down if an invalid/unverifiable certificate is found.
	VerifyCertificate *bool `json:"verify_certificate,omitempty"`
	// (http specific) Treat the target site as down if a certificate expires within the given number of days. This parameter will be ignored if `verify_certificate` is set to `false`.
	SslDownDaysBefore *int32 `json:"ssl_down_days_before,omitempty"`
}

type _HTTP HTTP

// NewHTTP instantiates a new HTTP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHTTP(host string, type_ string) *HTTP {
	this := HTTP{}
	this.Host = host
	this.Type = type_
	var responsetimeThreshold int32 = 30000
	this.ResponsetimeThreshold = &responsetimeThreshold
	var url string = "/"
	this.Url = &url
	var encryption bool = false
	this.Encryption = &encryption
	var port int32 = 80
	this.Port = &port
	var verifyCertificate bool = true
	this.VerifyCertificate = &verifyCertificate
	var sslDownDaysBefore int32 = 0
	this.SslDownDaysBefore = &sslDownDaysBefore
	return &this
}

// NewHTTPWithDefaults instantiates a new HTTP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHTTPWithDefaults() *HTTP {
	this := HTTP{}
	var responsetimeThreshold int32 = 30000
	this.ResponsetimeThreshold = &responsetimeThreshold
	var url string = "/"
	this.Url = &url
	var encryption bool = false
	this.Encryption = &encryption
	var port int32 = 80
	this.Port = &port
	var verifyCertificate bool = true
	this.VerifyCertificate = &verifyCertificate
	var sslDownDaysBefore int32 = 0
	this.SslDownDaysBefore = &sslDownDaysBefore
	return &this
}

// GetHost returns the Host field value
func (o *HTTP) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *HTTP) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *HTTP) SetHost(v string) {
	o.Host = v
}

// GetType returns the Type field value
func (o *HTTP) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HTTP) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HTTP) SetType(v string) {
	o.Type = v
}

// GetProbeid returns the Probeid field value if set, zero value otherwise.
func (o *HTTP) GetProbeid() int32 {
	if o == nil || IsNil(o.Probeid) {
		var ret int32
		return ret
	}
	return *o.Probeid
}

// GetProbeidOk returns a tuple with the Probeid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetProbeidOk() (*int32, bool) {
	if o == nil || IsNil(o.Probeid) {
		return nil, false
	}
	return o.Probeid, true
}

// HasProbeid returns a boolean if a field has been set.
func (o *HTTP) HasProbeid() bool {
	if o != nil && !IsNil(o.Probeid) {
		return true
	}

	return false
}

// SetProbeid gets a reference to the given int32 and assigns it to the Probeid field.
func (o *HTTP) SetProbeid(v int32) {
	o.Probeid = &v
}

// GetProbeFilters returns the ProbeFilters field value if set, zero value otherwise.
func (o *HTTP) GetProbeFilters() int32 {
	if o == nil || IsNil(o.ProbeFilters) {
		var ret int32
		return ret
	}
	return *o.ProbeFilters
}

// GetProbeFiltersOk returns a tuple with the ProbeFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetProbeFiltersOk() (*int32, bool) {
	if o == nil || IsNil(o.ProbeFilters) {
		return nil, false
	}
	return o.ProbeFilters, true
}

// HasProbeFilters returns a boolean if a field has been set.
func (o *HTTP) HasProbeFilters() bool {
	if o != nil && !IsNil(o.ProbeFilters) {
		return true
	}

	return false
}

// SetProbeFilters gets a reference to the given int32 and assigns it to the ProbeFilters field.
func (o *HTTP) SetProbeFilters(v int32) {
	o.ProbeFilters = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *HTTP) GetIpv6() bool {
	if o == nil || IsNil(o.Ipv6) {
		var ret bool
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetIpv6Ok() (*bool, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *HTTP) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given bool and assigns it to the Ipv6 field.
func (o *HTTP) SetIpv6(v bool) {
	o.Ipv6 = &v
}

// GetResponsetimeThreshold returns the ResponsetimeThreshold field value if set, zero value otherwise.
func (o *HTTP) GetResponsetimeThreshold() int32 {
	if o == nil || IsNil(o.ResponsetimeThreshold) {
		var ret int32
		return ret
	}
	return *o.ResponsetimeThreshold
}

// GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetResponsetimeThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ResponsetimeThreshold) {
		return nil, false
	}
	return o.ResponsetimeThreshold, true
}

// HasResponsetimeThreshold returns a boolean if a field has been set.
func (o *HTTP) HasResponsetimeThreshold() bool {
	if o != nil && !IsNil(o.ResponsetimeThreshold) {
		return true
	}

	return false
}

// SetResponsetimeThreshold gets a reference to the given int32 and assigns it to the ResponsetimeThreshold field.
func (o *HTTP) SetResponsetimeThreshold(v int32) {
	o.ResponsetimeThreshold = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *HTTP) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *HTTP) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *HTTP) SetUrl(v string) {
	o.Url = &v
}

// GetEncryption returns the Encryption field value if set, zero value otherwise.
func (o *HTTP) GetEncryption() bool {
	if o == nil || IsNil(o.Encryption) {
		var ret bool
		return ret
	}
	return *o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetEncryptionOk() (*bool, bool) {
	if o == nil || IsNil(o.Encryption) {
		return nil, false
	}
	return o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *HTTP) HasEncryption() bool {
	if o != nil && !IsNil(o.Encryption) {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given bool and assigns it to the Encryption field.
func (o *HTTP) SetEncryption(v bool) {
	o.Encryption = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *HTTP) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *HTTP) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *HTTP) SetPort(v int32) {
	o.Port = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *HTTP) GetAuth() string {
	if o == nil || IsNil(o.Auth) {
		var ret string
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetAuthOk() (*string, bool) {
	if o == nil || IsNil(o.Auth) {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *HTTP) HasAuth() bool {
	if o != nil && !IsNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given string and assigns it to the Auth field.
func (o *HTTP) SetAuth(v string) {
	o.Auth = &v
}

// GetShouldcontain returns the Shouldcontain field value if set, zero value otherwise.
func (o *HTTP) GetShouldcontain() string {
	if o == nil || IsNil(o.Shouldcontain) {
		var ret string
		return ret
	}
	return *o.Shouldcontain
}

// GetShouldcontainOk returns a tuple with the Shouldcontain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetShouldcontainOk() (*string, bool) {
	if o == nil || IsNil(o.Shouldcontain) {
		return nil, false
	}
	return o.Shouldcontain, true
}

// HasShouldcontain returns a boolean if a field has been set.
func (o *HTTP) HasShouldcontain() bool {
	if o != nil && !IsNil(o.Shouldcontain) {
		return true
	}

	return false
}

// SetShouldcontain gets a reference to the given string and assigns it to the Shouldcontain field.
func (o *HTTP) SetShouldcontain(v string) {
	o.Shouldcontain = &v
}

// GetShouldnotcontain returns the Shouldnotcontain field value if set, zero value otherwise.
func (o *HTTP) GetShouldnotcontain() string {
	if o == nil || IsNil(o.Shouldnotcontain) {
		var ret string
		return ret
	}
	return *o.Shouldnotcontain
}

// GetShouldnotcontainOk returns a tuple with the Shouldnotcontain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetShouldnotcontainOk() (*string, bool) {
	if o == nil || IsNil(o.Shouldnotcontain) {
		return nil, false
	}
	return o.Shouldnotcontain, true
}

// HasShouldnotcontain returns a boolean if a field has been set.
func (o *HTTP) HasShouldnotcontain() bool {
	if o != nil && !IsNil(o.Shouldnotcontain) {
		return true
	}

	return false
}

// SetShouldnotcontain gets a reference to the given string and assigns it to the Shouldnotcontain field.
func (o *HTTP) SetShouldnotcontain(v string) {
	o.Shouldnotcontain = &v
}

// GetPostdata returns the Postdata field value if set, zero value otherwise.
func (o *HTTP) GetPostdata() string {
	if o == nil || IsNil(o.Postdata) {
		var ret string
		return ret
	}
	return *o.Postdata
}

// GetPostdataOk returns a tuple with the Postdata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetPostdataOk() (*string, bool) {
	if o == nil || IsNil(o.Postdata) {
		return nil, false
	}
	return o.Postdata, true
}

// HasPostdata returns a boolean if a field has been set.
func (o *HTTP) HasPostdata() bool {
	if o != nil && !IsNil(o.Postdata) {
		return true
	}

	return false
}

// SetPostdata gets a reference to the given string and assigns it to the Postdata field.
func (o *HTTP) SetPostdata(v string) {
	o.Postdata = &v
}

// GetRequestheaderX returns the RequestheaderX field value if set, zero value otherwise.
func (o *HTTP) GetRequestheaderX() string {
	if o == nil || IsNil(o.RequestheaderX) {
		var ret string
		return ret
	}
	return *o.RequestheaderX
}

// GetRequestheaderXOk returns a tuple with the RequestheaderX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetRequestheaderXOk() (*string, bool) {
	if o == nil || IsNil(o.RequestheaderX) {
		return nil, false
	}
	return o.RequestheaderX, true
}

// HasRequestheaderX returns a boolean if a field has been set.
func (o *HTTP) HasRequestheaderX() bool {
	if o != nil && !IsNil(o.RequestheaderX) {
		return true
	}

	return false
}

// SetRequestheaderX gets a reference to the given string and assigns it to the RequestheaderX field.
func (o *HTTP) SetRequestheaderX(v string) {
	o.RequestheaderX = &v
}

// GetVerifyCertificate returns the VerifyCertificate field value if set, zero value otherwise.
func (o *HTTP) GetVerifyCertificate() bool {
	if o == nil || IsNil(o.VerifyCertificate) {
		var ret bool
		return ret
	}
	return *o.VerifyCertificate
}

// GetVerifyCertificateOk returns a tuple with the VerifyCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetVerifyCertificateOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyCertificate) {
		return nil, false
	}
	return o.VerifyCertificate, true
}

// HasVerifyCertificate returns a boolean if a field has been set.
func (o *HTTP) HasVerifyCertificate() bool {
	if o != nil && !IsNil(o.VerifyCertificate) {
		return true
	}

	return false
}

// SetVerifyCertificate gets a reference to the given bool and assigns it to the VerifyCertificate field.
func (o *HTTP) SetVerifyCertificate(v bool) {
	o.VerifyCertificate = &v
}

// GetSslDownDaysBefore returns the SslDownDaysBefore field value if set, zero value otherwise.
func (o *HTTP) GetSslDownDaysBefore() int32 {
	if o == nil || IsNil(o.SslDownDaysBefore) {
		var ret int32
		return ret
	}
	return *o.SslDownDaysBefore
}

// GetSslDownDaysBeforeOk returns a tuple with the SslDownDaysBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTP) GetSslDownDaysBeforeOk() (*int32, bool) {
	if o == nil || IsNil(o.SslDownDaysBefore) {
		return nil, false
	}
	return o.SslDownDaysBefore, true
}

// HasSslDownDaysBefore returns a boolean if a field has been set.
func (o *HTTP) HasSslDownDaysBefore() bool {
	if o != nil && !IsNil(o.SslDownDaysBefore) {
		return true
	}

	return false
}

// SetSslDownDaysBefore gets a reference to the given int32 and assigns it to the SslDownDaysBefore field.
func (o *HTTP) SetSslDownDaysBefore(v int32) {
	o.SslDownDaysBefore = &v
}

func (o HTTP) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HTTP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	toSerialize["type"] = o.Type
	if !IsNil(o.Probeid) {
		toSerialize["probeid"] = o.Probeid
	}
	if !IsNil(o.ProbeFilters) {
		toSerialize["probe_filters"] = o.ProbeFilters
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	if !IsNil(o.ResponsetimeThreshold) {
		toSerialize["responsetime_threshold"] = o.ResponsetimeThreshold
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Encryption) {
		toSerialize["encryption"] = o.Encryption
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	if !IsNil(o.Shouldcontain) {
		toSerialize["shouldcontain"] = o.Shouldcontain
	}
	if !IsNil(o.Shouldnotcontain) {
		toSerialize["shouldnotcontain"] = o.Shouldnotcontain
	}
	if !IsNil(o.Postdata) {
		toSerialize["postdata"] = o.Postdata
	}
	if !IsNil(o.RequestheaderX) {
		toSerialize["requestheader{X}"] = o.RequestheaderX
	}
	if !IsNil(o.VerifyCertificate) {
		toSerialize["verify_certificate"] = o.VerifyCertificate
	}
	if !IsNil(o.SslDownDaysBefore) {
		toSerialize["ssl_down_days_before"] = o.SslDownDaysBefore
	}
	return toSerialize, nil
}

func (o *HTTP) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"host",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHTTP := _HTTP{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHTTP)

	if err != nil {
		return err
	}

	*o = HTTP(varHTTP)

	return err
}

type NullableHTTP struct {
	value *HTTP
	isSet bool
}

func (v NullableHTTP) Get() *HTTP {
	return v.value
}

func (v *NullableHTTP) Set(val *HTTP) {
	v.value = val
	v.isSet = true
}

func (v NullableHTTP) IsSet() bool {
	return v.isSet
}

func (v *NullableHTTP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHTTP(val *HTTP) *NullableHTTP {
	return &NullableHTTP{value: val, isSet: true}
}

func (v NullableHTTP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHTTP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


