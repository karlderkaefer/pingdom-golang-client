# Go API client for openapi

# Welcome to the Pingdom API!
The Pingdom API is a way for you to automate your interaction with the Pingdom system. With the API, you can create your own scripts or applications with most of the functionality you can find inside the Pingdom control panel.

The Pingdom API is RESTful and HTTP-based. Basically, this means that the communication is made through normal HTTP requests.

# Authentication
Authentication is needed in order to use the Pingdom API, and for this a Pingdom API Token is required. You generate your Pingdom API Token inside My Pingdom. The Pingdom API Token has a property called “Access level” to define its permissions. All operations that create or modify something (e.g. checks) need the Read/Write permission. If you only need to read data using the API token, we recommend to set the access level to “Read access”.

The authentication method for using tokens is HTTP Bearer Authentication (encrypted over HTTPS). This means that you will provide your token every time you make a request. No sessions are used.

Request
```
GET /checks HTTP/1.1
Host: api.pingdom.com
Authorization: Bearer ofOhK18Ca6w4S_XmInGv0QPkqly-rbRBBoHsp_2FEH5QnIbH0VZhRPO3tlvrjMIKQ36VapX
```

Response
```
HTTP/1.1 200 OK
Content-Length: 13
Content-Type: application/json
{\"checks\":[]}
```

## Basic Auth
For compatibility reasons, the Pingdom API allows to use HTTP Basic Authentication with tokens. In cases where this is necessary, input the API token as the username and leave the password field empty.

An example request of how that would look like with the following API token: ofOhK18Ca6w4S_XmInGv0QPkqly-rbRBBoHsp_2FEH5QnIbH0VZhRPO3tlvrjMIKQ36VapX

```
GET /checks HTTP/1.1
Host: api.pingdom.com
Authorization: Basic b2ZPaEsxOENhNnc0U19YbUluR3YwUVBrcWx5LXJiUkJCb0hzcF8yRkVINVFuSWJIMFZaaFJQTzN0bHZyak1JS1EzNlZhcFg6
```

# Server Address
The base server address is: https://api.pingdom.com

Please note that HTTPS is required. You will not be able to connect through unencrypted HTTP.

# Providing Parameters
GET requests should provide their parameters as a query string, part of the URL.

POST, PUT and DELETE requests should provide their parameters as a JSON. This should be part of the request body. Remember to add the proper content type header to the request: `Content-Type: application/json`.

We still support providing parameters as a query string for POST, PUT and DELETE requests, but we recommend using JSON going forward. If you are using query strings, they should be part of the body, URL or a combination. The encoding of the query string should be standard URL-encoding, as provided by various programming libraries.

When using `requests` library for Python, use `json` parameter instead of `data`. Due to the inner mechanisms of requests.post() etc. some endpoints may return responses not conforming to the documentation when dealing with `data` body.

# HTTP/1.1 Status Code Definitions
The HTTP status code returned by a successful API request is defined in the documentation for that method. Usually, this will be 200 OK.

If something goes wrong, other codes may be returned. The API uses standard HTTP/1.1 status codes defined by [RFC 2616](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html).

# JSON Responses
All responses are sent JSON-encoded. The specific responses (successful ones) are described in the documentation section for each method.

However, if something goes wrong, our standard JSON error message (together with an appropriate status code) follows this format:
```
{
    \"error\": {
        \"statuscode\": 403,
        \"statusdesc\": \"Forbidden\",
        \"errormessage\":\" Something went wrong! This string describes what happened.\"
    }
}
```
See http://en.wikipedia.org/wiki/Json for more information on JSON.

Please note that all attributes of a method response are not always present. A client application should never assume that a certain attribute is present in a response.

# Limits
The Pingdom API has usage limits to avoid individual rampant applications degrading the overall user experience. There are two layers of limits, the first cover a shorter period of time and the second a longer period. This enables you to \"burst\" requests for shorter periods. There are two HTTP headers in every response describing your limits status.

The response headers are:
* **Req-Limit-Short**
* **Req-Limit-Long**

An example of the values of these headers:
* **Req-Limit-Short: Remaining: 394 Time until reset: 3589**
* **Req-Limit-Long: Remaining: 71994 Time until reset: 2591989**

In this case, we can see that the user has 394 requests left until the short limit is reached. In 3589 seconds, the short limit will be reset. In similar fashion, the long limit has 71994 requests left, and will be reset in 2591989 seconds.

If limits are exceeded, an HTTP 429 error code with the message \"Request limit exceeded, try again later\" is sent back.

# gzip
Responses can be gzip-encoded on demand. This is nice if your bandwidth is limited, or if big results are causing performance issues.

To enable gzip, simply add the following header to your request:

Accept-Encoding: gzip

# Best Practices
## Use caching
If you are building a web page using the Pingdom API, we recommend that you do all API request on the server side, and if possible cache them. If you get any substantial traffic, you do not want to call the API each time you get a page hit, since this may cause you to hit the request limit faster than expected. In general, whenever you can cache data, do so.

## Send your user credentials in a preemptive manner
Some HTTP clients omit the authentication header, and make a second request with the header when they get a 401 Unauthorized response. Please make sure you send the credentials directly, to avoid unnecessary requests.

## Use common sense
Should be simple enough. For example, don't check for the status of a check every other second. The highest check resolution is one minute. Checking more often than that won't give you much of an advantage.

## The Internet is unreliable
Networks in general are unreliable, and particularly one as large and complex as the Internet. Your application should not assume it will get an answer. There may be timeouts.

# PHP Code Example
**\"This is too much to read. I just want to get started right now! Give me a simple example!\"**

Here is a short example of how you can use the API with PHP. You need the cURL extension for PHP.

The example prints the current status of all your checks. This sample obviously focuses on Pingdom API code and does not worry about any potential problems connecting to the API, but your code should.

Code:
```php
<?php
    // Init cURL
    $curl = curl_init();
    // Set target URL
    curl_setopt($curl, CURLOPT_URL, \"https://api.pingdom.com/api/3.1/checks\");
    // Set the desired HTTP method (GET is default, see the documentation for each request)
    curl_setopt($curl, CURLOPT_CUSTOMREQUEST, \"GET\");
    // Add header with Bearer Authorization
    curl_setopt($curl, CURLOPT_HTTPHEADER, array(\"Authorization: Bearer 907c762e069589c2cd2a229cdae7b8778caa9f07\"));
    // Ask cURL to return the result as a string
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
    // Execute the request and decode the json result into an associative array
    $response = json_decode(curl_exec($curl), true);
    // Check for errors returned by the API
    if (isset($response['error'])) {
        print \"Error: \" . $response['error']['errormessage'] . \"\\n\";
        exit;
    }
    // Fetch the list of checks from the response
    $checksList = $response['checks'];
    // Print the names and statuses of all checks in the list
    foreach ($checksList as $check) {
        print $check['name'] . \" is \" . $check['status'] . \"\\n\";
    }
?>
```

Example output:
```
Ubuntu Packages is up
Google is up
Pingdom is up
My server 1 is down
My server 2 is up
```

If you are running PHP on Windows, you need to be sure that you have installed the CA certificates for HTTPS/SSL to work correctly. Please see the cURL manual for more information. As a quick (but unsafe) workaround, you can add the following cURL option to ignore certificate validation.

`
curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, 0);
`

# TMS Steps Vocabulary

There are two types of transaction checks:
<ul>
    <li><b>script</b>: the legacy TMS check created with predefined commands in the Pingdom UI or via the public API</li>
    <li><b>recording</b>: the TMS check created by recording performed actions in WPM recorder.
        More information about how to use it can be found in the
        <a class=\"default-link\" href=\"https://documentation.solarwinds.com/en/success_center/wpm/Content/WPM-Use-the-WPM3-Recorder.htm\">
            WPM recorder documentation</a>
    </li>
</ul>

## Script transaction checks

  ### Commands
  Actions to be performed for the script TMS check

  Step Name                                 | \"fn\"                  | Required \"args\"     | Remarks
  ----------------------------------------- | --------------------- | --------------      | -------
  Go to URL                                 | go_to                 | url                 | There has to be **exactly one** go_to step
  Click                                     | click                 | element             | label, name or CSS selector
  Fill in field                             | fill                  | input, value        | input: label, name or CSS selector, value: text
  Check checkbox                            | check                 | checkbox            | label, name or CSS selector
  Uncheck checkbox                          | uncheck               | checkbox            | label, name or CSS selector
  Sleep for                                 | sleep                 | seconds             | number (e.g. 0.1)
  Select radio button                       | select_radio          | radio               | name of the radio button
  Select dropdown                           | select                | select, option      | select: label, name or CSS selector, option: content, value or CSS selector
  Basic auth login with                     | basic_auth            | username, password  | username and password as text
  Submit form                               | submit                | form                | name or CSS selector
  Wait for element                          | wait_for_element      | element             | label, name or CSS selector
  Wait for element to contain               | wait_for_contains     | element, value      | element: label, name or CSS selector, value: text

  ### Validations
  Verify the state of the page

  Step Name                                 | \"fn\"                  | Required \"args\"     | Remarks
  ----------------------------------------- | --------------------- | --------------      | -------
  URL should be                             | url                   | url                 | url to be verified
  Element should exist                      | exists                | element             | label, name or CSS selector
  Element shouldn't exist                   | not_exists            | element             | label, name or CSS selector
  Element should contain                    | contains              | element, value      | element: label, name or CSS selector, value: text
  Element shouldn't containt                | not_contains          | element, value      | element: label, name or CSS selector, value: text
  Text field should contain                 | field_contains        | input, value        | input: label, name or CSS selector, value: text
  Text field shouldn't contain              | field_not_contains    | input, value        | input: label, name or CSS selector, value: text
  Checkbox should be checked                | is_checked            | checkbox            | label, name or CSS selector
  Checkbox shouldn't be checked             | is_not_checked        | checkbox            | label, name or CSS selector
  Radio button should be selected           | radio_selected        | radio               | name of the radio button
  Dropdown with name should be selected     | dropdown_selected     | select, option      | select: label, name or CSS selector, option: content, value or CSS selector
  Dropdown with name shouldn't be selected  | dropdown_not_selected | select, option      | select: label, name or CSS selector, option: content, value or CSS selector

  ### Example steps

  ```
  \"steps\": [
  {
    \"fn\": \"go_to\",
    \"args\": {
      \"url\": \"pingdom.com\"
    }
  },
  {
    \"fn\": \"click\",
    \"args\": {
      \"element\": \"START FREE TRIAL\"
    }
  },
  {
    \"fn\": \"url\",
    \"args\": {
      \"url\": \"https://www.pingdom.com/sign-up/\"
    }
  }
  ]
  ```

## Recording transaction checks

Each of check steps contains:
<ul>
  <li><b>fn</b>: function name of the step</li>
  <li><b>args</b>: function arguments</li>
  <li><b>guid</b>: automatically generated identifier</li>
  <li><b>contains_navigate</b>: recorder sets it on true if the step would trigger a page navigation</li>
</ul>

  ### Commands
  Actions to be performed for the recording TMS check

  Step Name                 | \"fn\"                      | Required \"args\"                 | Remarks
  ------------------------- | ------------------------- | ------------------------------- | -------
  Go to URL                 | wpm_go_to                 | url                             | Goes to the given url location
  Click                     | wpm_click                 | element, offsetX, offsetY       | **element**: label, name or CSS selector,</br> **offsetX/Y**: exact position of a click in the element
  Click in a exact location | wpm_click_xy              | element, x, y, scrollX, scrollY | **element**: label, name or CSS selector,</br> **x/y**: coordinates for the click (in viewport),</br> **scrollX/Y**: scrollbar position
  Fill                      | wpm_fill                  | input, value                    | **input**: target element,</br> **value**: text to fill the target
  Move mouse to element     | wpm_move_mouse_to_element | element, offsetX, offsetY       | **element**: target element,</br> **offsetX/Y**: exact position of the mouse in the element
  Select dropdown           | wpm_select_dropdown       | select, option                  | **select**: target element (drop-down),</br> **option**: text of the option to select
  Wait                      | wpm_wait                  | seconds                         | **seconds:** numbers of seconds to wait
  Close tab                 | wpm_close_tab             | -                               | Closes the latest tab on the tab stack

  ### Validations
  Verify the state of the page

  Step Name              | \"fn\"                     | Required \"args\"                                | Remarks
  ---------------------- | ------------------------ | ---------------------------------------------- | -------
  Contains text          | wpm_contains_timeout     | element, value, waitTime, useRegularExpression | **element**: label, name or CSS selector,</br> **value**: text to search for,</br> **waitTime**: time to wait for the value to appear,</br> **useRegularExpression**: use the value as a RegEx
  Does not contains text | wpm_not_contains_timeout | element, value, waitTime, useRegularExpression | **element**: label, name or CSS selector,</br> **value**: text to search for,</br> **waitTime**: time to wait for the value to appear,</br> **useRegularExpression**: use the value as a RegEx

  ### Metadata
  Recording checks contain metadata which is automatically generated by the WPM recorder. Modify with caution!


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.pingdom.com/api/3.1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActionsAPI* | [**ActionsGet**](docs/ActionsAPI.md#actionsget) | **Get** /actions | Returns a list of actions alerts.
*AnalysisAPI* | [**AnalysisCheckidAnalysisidGet**](docs/AnalysisAPI.md#analysischeckidanalysisidget) | **Get** /analysis/{checkid}/{analysisid} | Returns the raw result for a specified analysis.
*AnalysisAPI* | [**AnalysisCheckidGet**](docs/AnalysisAPI.md#analysischeckidget) | **Get** /analysis/{checkid} | Returns a list of the latest root cause analysis
*ChecksAPI* | [**ChecksCheckidDelete**](docs/ChecksAPI.md#checkscheckiddelete) | **Delete** /checks/{checkid} | Deletes a check.
*ChecksAPI* | [**ChecksCheckidGet**](docs/ChecksAPI.md#checkscheckidget) | **Get** /checks/{checkid} | Returns a detailed description of a check.
*ChecksAPI* | [**ChecksCheckidPut**](docs/ChecksAPI.md#checkscheckidput) | **Put** /checks/{checkid} | Modify settings for a check.
*ChecksAPI* | [**ChecksDelete**](docs/ChecksAPI.md#checksdelete) | **Delete** /checks | Deletes a list of checks.
*ChecksAPI* | [**ChecksGet**](docs/ChecksAPI.md#checksget) | **Get** /checks | 
*ChecksAPI* | [**ChecksPost**](docs/ChecksAPI.md#checkspost) | **Post** /checks | Creates a new check.
*ChecksAPI* | [**ChecksPut**](docs/ChecksAPI.md#checksput) | **Put** /checks | Pause or change resolution for multiple checks.
*ContactsAPI* | [**AlertingContactsContactidDelete**](docs/ContactsAPI.md#alertingcontactscontactiddelete) | **Delete** /alerting/contacts/{contactid} | Deletes a contact with its contact methods
*ContactsAPI* | [**AlertingContactsContactidGet**](docs/ContactsAPI.md#alertingcontactscontactidget) | **Get** /alerting/contacts/{contactid} | Returns a contact with its contact methods
*ContactsAPI* | [**AlertingContactsContactidPut**](docs/ContactsAPI.md#alertingcontactscontactidput) | **Put** /alerting/contacts/{contactid} | Update a contact
*ContactsAPI* | [**AlertingContactsGet**](docs/ContactsAPI.md#alertingcontactsget) | **Get** /alerting/contacts | Returns a list of all contacts
*ContactsAPI* | [**AlertingContactsPost**](docs/ContactsAPI.md#alertingcontactspost) | **Post** /alerting/contacts | Creates a new contact
*CreditsAPI* | [**CreditsGet**](docs/CreditsAPI.md#creditsget) | **Get** /credits | Returns information about remaining credits
*MaintenanceAPI* | [**MaintenanceDelete**](docs/MaintenanceAPI.md#maintenancedelete) | **Delete** /maintenance | Delete multiple maintenance windows.
*MaintenanceAPI* | [**MaintenanceGet**](docs/MaintenanceAPI.md#maintenanceget) | **Get** /maintenance | 
*MaintenanceAPI* | [**MaintenanceIdDelete**](docs/MaintenanceAPI.md#maintenanceiddelete) | **Delete** /maintenance/{id} | Delete the maintenance window.
*MaintenanceAPI* | [**MaintenanceIdGet**](docs/MaintenanceAPI.md#maintenanceidget) | **Get** /maintenance/{id} | 
*MaintenanceAPI* | [**MaintenanceIdPut**](docs/MaintenanceAPI.md#maintenanceidput) | **Put** /maintenance/{id} | 
*MaintenanceAPI* | [**MaintenancePost**](docs/MaintenanceAPI.md#maintenancepost) | **Post** /maintenance | 
*MaintenanceOccurrencesAPI* | [**MaintenanceOccurrencesDelete**](docs/MaintenanceOccurrencesAPI.md#maintenanceoccurrencesdelete) | **Delete** /maintenance.occurrences | Deletes multiple maintenance occurrences
*MaintenanceOccurrencesAPI* | [**MaintenanceOccurrencesGet**](docs/MaintenanceOccurrencesAPI.md#maintenanceoccurrencesget) | **Get** /maintenance.occurrences | Returns a list of maintenance occurrences.
*MaintenanceOccurrencesAPI* | [**MaintenanceOccurrencesIdDelete**](docs/MaintenanceOccurrencesAPI.md#maintenanceoccurrencesiddelete) | **Delete** /maintenance.occurrences/{id} | Deletes the maintenance occurrence
*MaintenanceOccurrencesAPI* | [**MaintenanceOccurrencesIdGet**](docs/MaintenanceOccurrencesAPI.md#maintenanceoccurrencesidget) | **Get** /maintenance.occurrences/{id} | Gets a maintenance occurrence details
*MaintenanceOccurrencesAPI* | [**MaintenanceOccurrencesIdPut**](docs/MaintenanceOccurrencesAPI.md#maintenanceoccurrencesidput) | **Put** /maintenance.occurrences/{id} | Modifies a maintenance occurrence
*ProbesAPI* | [**ProbesGet**](docs/ProbesAPI.md#probesget) | **Get** /probes | Returns a list of Pingdom probe servers
*ReferenceAPI* | [**ReferenceGet**](docs/ReferenceAPI.md#referenceget) | **Get** /reference | Get regions, timezone and date/time/number references
*ResultsAPI* | [**ResultsCheckidGet**](docs/ResultsAPI.md#resultscheckidget) | **Get** /results/{checkid} | Return a list of raw test results
*SingleAPI* | [**SingleGet**](docs/SingleAPI.md#singleget) | **Get** /single | Performs a single check.
*SummaryAverageAPI* | [**SummaryAverageCheckidGet**](docs/SummaryAverageAPI.md#summaryaveragecheckidget) | **Get** /summary.average/{checkid} | Get the average time/uptime value for a specified
*SummaryHoursofdayAPI* | [**SummaryHoursofdayCheckidGet**](docs/SummaryHoursofdayAPI.md#summaryhoursofdaycheckidget) | **Get** /summary.hoursofday/{checkid} | Returns the average response time for each hour.
*SummaryOutageAPI* | [**SummaryOutageCheckidGet**](docs/SummaryOutageAPI.md#summaryoutagecheckidget) | **Get** /summary.outage/{checkid} | Get a list of status changes for a specified check
*SummaryPerformanceAPI* | [**SummaryPerformanceCheckidGet**](docs/SummaryPerformanceAPI.md#summaryperformancecheckidget) | **Get** /summary.performance/{checkid} | For a given interval return a list of subintervals
*SummaryProbesAPI* | [**SummaryProbesCheckidGet**](docs/SummaryProbesAPI.md#summaryprobescheckidget) | **Get** /summary.probes/{checkid} | Get a list of probes that performed tests
*TMSChecksAPI* | [**AddCheck**](docs/TMSChecksAPI.md#addcheck) | **Post** /tms/check | Creates a new transaction check.
*TMSChecksAPI* | [**DeleteCheck**](docs/TMSChecksAPI.md#deletecheck) | **Delete** /tms/check/{cid} | Deletes a transaction check.
*TMSChecksAPI* | [**GetAllChecks**](docs/TMSChecksAPI.md#getallchecks) | **Get** /tms/check | Returns a list overview of all transaction checks.
*TMSChecksAPI* | [**GetCheck**](docs/TMSChecksAPI.md#getcheck) | **Get** /tms/check/{cid} | Returns a single transaction check.
*TMSChecksAPI* | [**GetCheckReportPerformance**](docs/TMSChecksAPI.md#getcheckreportperformance) | **Get** /tms/check/{check_id}/report/performance | Returns a performance report for a single transaction check
*TMSChecksAPI* | [**GetCheckReportStatus**](docs/TMSChecksAPI.md#getcheckreportstatus) | **Get** /tms/check/{check_id}/report/status | Returns a status change report for a single transaction check
*TMSChecksAPI* | [**GetCheckReportStatusAll**](docs/TMSChecksAPI.md#getcheckreportstatusall) | **Get** /tms/check/report/status | Returns a status change report for all transaction checks in the current organization
*TMSChecksAPI* | [**ModifyCheck**](docs/TMSChecksAPI.md#modifycheck) | **Put** /tms/check/{cid} | Modify settings for transaction check.
*TeamsAPI* | [**AlertingTeamsGet**](docs/TeamsAPI.md#alertingteamsget) | **Get** /alerting/teams | 
*TeamsAPI* | [**AlertingTeamsPost**](docs/TeamsAPI.md#alertingteamspost) | **Post** /alerting/teams | Creates a new team
*TeamsAPI* | [**AlertingTeamsTeamidDelete**](docs/TeamsAPI.md#alertingteamsteamiddelete) | **Delete** /alerting/teams/{teamid} | 
*TeamsAPI* | [**AlertingTeamsTeamidGet**](docs/TeamsAPI.md#alertingteamsteamidget) | **Get** /alerting/teams/{teamid} | 
*TeamsAPI* | [**AlertingTeamsTeamidPut**](docs/TeamsAPI.md#alertingteamsteamidput) | **Put** /alerting/teams/{teamid} | 
*TracerouteAPI* | [**TracerouteGet**](docs/TracerouteAPI.md#tracerouteget) | **Get** /traceroute | Perform a traceroute


## Documentation For Models

 - [AGCMInner](docs/AGCMInner.md)
 - [APNSInner](docs/APNSInner.md)
 - [ActionsAlertsEntry](docs/ActionsAlertsEntry.md)
 - [ActionsAlertsEntryActions](docs/ActionsAlertsEntryActions.md)
 - [ActionsAlertsEntryActionsAlertsInner](docs/ActionsAlertsEntryActionsAlertsInner.md)
 - [AlertingContactsContactidDelete200Response](docs/AlertingContactsContactidDelete200Response.md)
 - [AlertingContactsPost200Response](docs/AlertingContactsPost200Response.md)
 - [AlertingContactsPost200ResponseContact](docs/AlertingContactsPost200ResponseContact.md)
 - [AlertingTeamID](docs/AlertingTeamID.md)
 - [AlertingTeams](docs/AlertingTeams.md)
 - [AlertingTeamsPost200Response](docs/AlertingTeamsPost200Response.md)
 - [AlertingTeamsPost200ResponseTeam](docs/AlertingTeamsPost200ResponseTeam.md)
 - [AlertingTeamsTeamidDelete200Response](docs/AlertingTeamsTeamidDelete200Response.md)
 - [AnalysisRespAttrs](docs/AnalysisRespAttrs.md)
 - [AnalysisRespAttrsAnalysisInner](docs/AnalysisRespAttrsAnalysisInner.md)
 - [Check](docs/Check.md)
 - [CheckGeneral](docs/CheckGeneral.md)
 - [CheckSimple](docs/CheckSimple.md)
 - [CheckStatus](docs/CheckStatus.md)
 - [CheckWithStringType](docs/CheckWithStringType.md)
 - [CheckWithoutID](docs/CheckWithoutID.md)
 - [CheckWithoutIDGET](docs/CheckWithoutIDGET.md)
 - [CheckWithoutIDPUT](docs/CheckWithoutIDPUT.md)
 - [Checks](docs/Checks.md)
 - [ChecksAll](docs/ChecksAll.md)
 - [ChecksCheckidDelete200Response](docs/ChecksCheckidDelete200Response.md)
 - [ChecksCheckidPut200Response](docs/ChecksCheckidPut200Response.md)
 - [ChecksDelete200Response](docs/ChecksDelete200Response.md)
 - [ChecksPost200Response](docs/ChecksPost200Response.md)
 - [ChecksPost200ResponseCheck](docs/ChecksPost200ResponseCheck.md)
 - [ChecksPut200Response](docs/ChecksPut200Response.md)
 - [ChecksPutRequest](docs/ChecksPutRequest.md)
 - [Contact](docs/Contact.md)
 - [ContactTargets](docs/ContactTargets.md)
 - [ContactTargetsNotificationTargets](docs/ContactTargetsNotificationTargets.md)
 - [ContactTargetsTeamsInner](docs/ContactTargetsTeamsInner.md)
 - [ContactsList](docs/ContactsList.md)
 - [Country](docs/Country.md)
 - [Counts](docs/Counts.md)
 - [CreateCheck](docs/CreateCheck.md)
 - [CreateContact](docs/CreateContact.md)
 - [CreateContactNotificationTargets](docs/CreateContactNotificationTargets.md)
 - [CreateTeam](docs/CreateTeam.md)
 - [CreditsRespAttrs](docs/CreditsRespAttrs.md)
 - [CreditsRespAttrsCredits](docs/CreditsRespAttrsCredits.md)
 - [DNS](docs/DNS.md)
 - [DateTimeFormat](docs/DateTimeFormat.md)
 - [Days](docs/Days.md)
 - [DeleteCheck200Response](docs/DeleteCheck200Response.md)
 - [DetailedCheck](docs/DetailedCheck.md)
 - [DetailedCheckAttributes](docs/DetailedCheckAttributes.md)
 - [DetailedCheckDns](docs/DetailedCheckDns.md)
 - [DetailedCheckDnsCheck](docs/DetailedCheckDnsCheck.md)
 - [DetailedCheckHttp](docs/DetailedCheckHttp.md)
 - [DetailedCheckHttpCheck](docs/DetailedCheckHttpCheck.md)
 - [DetailedCheckHttpCustom](docs/DetailedCheckHttpCustom.md)
 - [DetailedCheckHttpCustomCheck](docs/DetailedCheckHttpCustomCheck.md)
 - [DetailedCheckImap](docs/DetailedCheckImap.md)
 - [DetailedCheckImapCheck](docs/DetailedCheckImapCheck.md)
 - [DetailedCheckPop3](docs/DetailedCheckPop3.md)
 - [DetailedCheckPop3Check](docs/DetailedCheckPop3Check.md)
 - [DetailedCheckSmtp](docs/DetailedCheckSmtp.md)
 - [DetailedCheckSmtpCheck](docs/DetailedCheckSmtpCheck.md)
 - [DetailedCheckTcp](docs/DetailedCheckTcp.md)
 - [DetailedCheckTcpCheck](docs/DetailedCheckTcpCheck.md)
 - [DetailedCheckUdp](docs/DetailedCheckUdp.md)
 - [DetailedCheckUdpCheck](docs/DetailedCheckUdpCheck.md)
 - [DetailedDnsAttributes](docs/DetailedDnsAttributes.md)
 - [DetailedDnsAttributesType](docs/DetailedDnsAttributesType.md)
 - [DetailedHttpAttributes](docs/DetailedHttpAttributes.md)
 - [DetailedHttpAttributesType](docs/DetailedHttpAttributesType.md)
 - [DetailedHttpCustomAttributes](docs/DetailedHttpCustomAttributes.md)
 - [DetailedHttpCustomAttributesType](docs/DetailedHttpCustomAttributesType.md)
 - [DetailedImapAttributes](docs/DetailedImapAttributes.md)
 - [DetailedImapAttributesType](docs/DetailedImapAttributesType.md)
 - [DetailedPop3Attributes](docs/DetailedPop3Attributes.md)
 - [DetailedPop3AttributesType](docs/DetailedPop3AttributesType.md)
 - [DetailedSmtpAttributes](docs/DetailedSmtpAttributes.md)
 - [DetailedSmtpAttributesType](docs/DetailedSmtpAttributesType.md)
 - [DetailedTcpAttributes](docs/DetailedTcpAttributes.md)
 - [DetailedTcpAttributesType](docs/DetailedTcpAttributesType.md)
 - [DetailedUdpAttributes](docs/DetailedUdpAttributes.md)
 - [DetailedUdpAttributesType](docs/DetailedUdpAttributesType.md)
 - [DnsAttributes](docs/DnsAttributes.md)
 - [EmailsInner](docs/EmailsInner.md)
 - [HTTP](docs/HTTP.md)
 - [HTTPCustom](docs/HTTPCustom.md)
 - [Hours](docs/Hours.md)
 - [HttpAttributesBase](docs/HttpAttributesBase.md)
 - [HttpAttributesGet](docs/HttpAttributesGet.md)
 - [HttpAttributesSet](docs/HttpAttributesSet.md)
 - [HttpAuthentications](docs/HttpAuthentications.md)
 - [HttpAuthenticationsCredentials](docs/HttpAuthenticationsCredentials.md)
 - [HttpCertificateAttributes](docs/HttpCertificateAttributes.md)
 - [HttpCustomAttributes](docs/HttpCustomAttributes.md)
 - [IMAP](docs/IMAP.md)
 - [ImapAttributes](docs/ImapAttributes.md)
 - [MaintenanceDelete](docs/MaintenanceDelete.md)
 - [MaintenanceDeleteRespAttrs](docs/MaintenanceDeleteRespAttrs.md)
 - [MaintenanceIdDeleteRespAttrs](docs/MaintenanceIdDeleteRespAttrs.md)
 - [MaintenanceIdPut](docs/MaintenanceIdPut.md)
 - [MaintenanceIdPutRespAttrs](docs/MaintenanceIdPutRespAttrs.md)
 - [MaintenanceIdRespAttrs](docs/MaintenanceIdRespAttrs.md)
 - [MaintenanceIdRespAttrsMaintenance](docs/MaintenanceIdRespAttrsMaintenance.md)
 - [MaintenanceIdRespAttrsMaintenanceChecks](docs/MaintenanceIdRespAttrsMaintenanceChecks.md)
 - [MaintenanceOccurrencesDelete](docs/MaintenanceOccurrencesDelete.md)
 - [MaintenanceOccurrencesDeleteRespAttrs](docs/MaintenanceOccurrencesDeleteRespAttrs.md)
 - [MaintenanceOccurrencesIdDeleteRespAttrs](docs/MaintenanceOccurrencesIdDeleteRespAttrs.md)
 - [MaintenanceOccurrencesIdPut](docs/MaintenanceOccurrencesIdPut.md)
 - [MaintenanceOccurrencesIdPutRespAttrs](docs/MaintenanceOccurrencesIdPutRespAttrs.md)
 - [MaintenanceOccurrencesIdRespAttrs](docs/MaintenanceOccurrencesIdRespAttrs.md)
 - [MaintenanceOccurrencesIdRespAttrsOccurrence](docs/MaintenanceOccurrencesIdRespAttrsOccurrence.md)
 - [MaintenanceOccurrencesRespAttrs](docs/MaintenanceOccurrencesRespAttrs.md)
 - [MaintenanceOccurrencesRespAttrsOccurrencesInner](docs/MaintenanceOccurrencesRespAttrsOccurrencesInner.md)
 - [MaintenanceOrder](docs/MaintenanceOrder.md)
 - [MaintenanceOrderby](docs/MaintenanceOrderby.md)
 - [MaintenancePost](docs/MaintenancePost.md)
 - [MaintenancePostRespAttrs](docs/MaintenancePostRespAttrs.md)
 - [MaintenancePostRespAttrsMaintenance](docs/MaintenancePostRespAttrsMaintenance.md)
 - [MaintenanceRespAttrs](docs/MaintenanceRespAttrs.md)
 - [MaintenanceRespAttrsMaintenanceInner](docs/MaintenanceRespAttrsMaintenanceInner.md)
 - [MaintenanceRespAttrsMaintenanceInnerChecks](docs/MaintenanceRespAttrsMaintenanceInnerChecks.md)
 - [Members](docs/Members.md)
 - [Metadata](docs/Metadata.md)
 - [MetadataGET](docs/MetadataGET.md)
 - [MetadataGETAuthentications](docs/MetadataGETAuthentications.md)
 - [ModifyCheckSettings](docs/ModifyCheckSettings.md)
 - [NumberFormat](docs/NumberFormat.md)
 - [POP3](docs/POP3.md)
 - [PhoneCode](docs/PhoneCode.md)
 - [Pop3Attributes](docs/Pop3Attributes.md)
 - [Probe](docs/Probe.md)
 - [Probes](docs/Probes.md)
 - [References](docs/References.md)
 - [Region](docs/Region.md)
 - [ReportPerformance](docs/ReportPerformance.md)
 - [ReportPerformanceReport](docs/ReportPerformanceReport.md)
 - [ReportPerformanceReportIntervalsInner](docs/ReportPerformanceReportIntervalsInner.md)
 - [ReportPerformanceReportIntervalsInnerStepsInner](docs/ReportPerformanceReportIntervalsInnerStepsInner.md)
 - [ReportStatusAll](docs/ReportStatusAll.md)
 - [ReportStatusSingle](docs/ReportStatusSingle.md)
 - [ResultsRespAttrs](docs/ResultsRespAttrs.md)
 - [ResultsRespAttrsResultsInner](docs/ResultsRespAttrsResultsInner.md)
 - [SMSesInner](docs/SMSesInner.md)
 - [SMTP](docs/SMTP.md)
 - [SingleGetQueryParametersParameter](docs/SingleGetQueryParametersParameter.md)
 - [SingleResp](docs/SingleResp.md)
 - [SingleRespResult](docs/SingleRespResult.md)
 - [SmtpAttributesBase](docs/SmtpAttributesBase.md)
 - [SmtpAttributesGet](docs/SmtpAttributesGet.md)
 - [SmtpAttributesSet](docs/SmtpAttributesSet.md)
 - [State](docs/State.md)
 - [Step](docs/Step.md)
 - [StepArgs](docs/StepArgs.md)
 - [SummaryHoursofdayRespAttrs](docs/SummaryHoursofdayRespAttrs.md)
 - [SummaryHoursofdayRespAttrsHoursofdayInner](docs/SummaryHoursofdayRespAttrsHoursofdayInner.md)
 - [SummaryOutageOrder](docs/SummaryOutageOrder.md)
 - [SummaryOutageRespAttrs](docs/SummaryOutageRespAttrs.md)
 - [SummaryOutageRespAttrsSummary](docs/SummaryOutageRespAttrsSummary.md)
 - [SummaryOutageRespAttrsSummaryStatesInner](docs/SummaryOutageRespAttrsSummaryStatesInner.md)
 - [SummaryPerformanceOrder](docs/SummaryPerformanceOrder.md)
 - [SummaryPerformanceResolution](docs/SummaryPerformanceResolution.md)
 - [SummaryPerformanceRespAttrs](docs/SummaryPerformanceRespAttrs.md)
 - [SummaryPerformanceRespAttrsSummary](docs/SummaryPerformanceRespAttrsSummary.md)
 - [SummaryPerformanceResultsInner](docs/SummaryPerformanceResultsInner.md)
 - [SummaryProbesRespAttrs](docs/SummaryProbesRespAttrs.md)
 - [SummaryRespAttrs](docs/SummaryRespAttrs.md)
 - [SummaryRespAttrsSummary](docs/SummaryRespAttrsSummary.md)
 - [SummaryRespAttrsSummaryResponsetime](docs/SummaryRespAttrsSummaryResponsetime.md)
 - [SummaryRespAttrsSummaryResponsetimeAvgresponse](docs/SummaryRespAttrsSummaryResponsetimeAvgresponse.md)
 - [SummaryRespAttrsSummaryResponsetimeAvgresponseOneOfInner](docs/SummaryRespAttrsSummaryResponsetimeAvgresponseOneOfInner.md)
 - [SummaryRespAttrsSummaryStatus](docs/SummaryRespAttrsSummaryStatus.md)
 - [TCP](docs/TCP.md)
 - [Tag](docs/Tag.md)
 - [TcpAttributes](docs/TcpAttributes.md)
 - [TeamID](docs/TeamID.md)
 - [Teams](docs/Teams.md)
 - [Timezone](docs/Timezone.md)
 - [Traceroute](docs/Traceroute.md)
 - [TracerouteData](docs/TracerouteData.md)
 - [UDP](docs/UDP.md)
 - [UdpAttributes](docs/UdpAttributes.md)
 - [UpdateContact](docs/UpdateContact.md)
 - [UpdateTeam](docs/UpdateTeam.md)
 - [Weeks](docs/Weeks.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



