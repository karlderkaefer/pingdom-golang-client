# \TMSChecksAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCheck**](TMSChecksAPI.md#AddCheck) | **Post** /tms/check | Creates a new transaction check.
[**DeleteCheck**](TMSChecksAPI.md#DeleteCheck) | **Delete** /tms/check/{cid} | Deletes a transaction check.
[**GetAllChecks**](TMSChecksAPI.md#GetAllChecks) | **Get** /tms/check | Returns a list overview of all transaction checks.
[**GetCheck**](TMSChecksAPI.md#GetCheck) | **Get** /tms/check/{cid} | Returns a single transaction check.
[**GetCheckReportPerformance**](TMSChecksAPI.md#GetCheckReportPerformance) | **Get** /tms/check/{check_id}/report/performance | Returns a performance report for a single transaction check
[**GetCheckReportStatus**](TMSChecksAPI.md#GetCheckReportStatus) | **Get** /tms/check/{check_id}/report/status | Returns a status change report for a single transaction check
[**GetCheckReportStatusAll**](TMSChecksAPI.md#GetCheckReportStatusAll) | **Get** /tms/check/report/status | Returns a status change report for all transaction checks in the current organization
[**ModifyCheck**](TMSChecksAPI.md#ModifyCheck) | **Put** /tms/check/{cid} | Modify settings for transaction check.



## AddCheck

> CheckSimple AddCheck(ctx).CheckWithoutID(checkWithoutID).Execute()

Creates a new transaction check.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	checkWithoutID := *openapiclient.NewCheckWithoutID("My awesome check", []openapiclient.Step{*openapiclient.NewStep()}) // CheckWithoutID | Specifies the check to be added

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TMSChecksAPI.AddCheck(context.Background()).CheckWithoutID(checkWithoutID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TMSChecksAPI.AddCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCheck`: CheckSimple
	fmt.Fprintf(os.Stdout, "Response from `TMSChecksAPI.AddCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkWithoutID** | [**CheckWithoutID**](CheckWithoutID.md) | Specifies the check to be added | 

### Return type

[**CheckSimple**](CheckSimple.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCheck

> DeleteCheck200Response DeleteCheck(ctx, cid).Execute()

Deletes a transaction check.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cid := int64(789) // int64 | Specifies the id of the check to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TMSChecksAPI.DeleteCheck(context.Background(), cid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TMSChecksAPI.DeleteCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCheck`: DeleteCheck200Response
	fmt.Fprintf(os.Stdout, "Response from `TMSChecksAPI.DeleteCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int64** | Specifies the id of the check to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteCheck200Response**](DeleteCheck200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllChecks

> ChecksAll GetAllChecks(ctx).ExtendedTags(extendedTags).Tags(tags).Type_(type_).Limit(limit).Offset(offset).Execute()

Returns a list overview of all transaction checks.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	extendedTags := true // bool | Specifies whether to return an extended tags representation in the response (with type and count). (optional)
	tags := []string{"Inner_example"} // []string | Tag list separated by commas. As an example \"nginx,apache\" would filter out all responses except those tagged nginx or apache (optional)
	type_ := "type__example" // string | Filter to return only checks of a given type (a TMS `script` or a WPM `recording`). If not provided, all checks are returned. (optional)
	limit := "limit_example" // string | Limit of returned checks (optional) (default to "1000")
	offset := "offset_example" // string | Offset of returned checks (optional) (default to "0")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TMSChecksAPI.GetAllChecks(context.Background()).ExtendedTags(extendedTags).Tags(tags).Type_(type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TMSChecksAPI.GetAllChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllChecks`: ChecksAll
	fmt.Fprintf(os.Stdout, "Response from `TMSChecksAPI.GetAllChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extendedTags** | **bool** | Specifies whether to return an extended tags representation in the response (with type and count). | 
 **tags** | **[]string** | Tag list separated by commas. As an example \&quot;nginx,apache\&quot; would filter out all responses except those tagged nginx or apache | 
 **type_** | **string** | Filter to return only checks of a given type (a TMS &#x60;script&#x60; or a WPM &#x60;recording&#x60;). If not provided, all checks are returned. | 
 **limit** | **string** | Limit of returned checks | [default to &quot;1000&quot;]
 **offset** | **string** | Offset of returned checks | [default to &quot;0&quot;]

### Return type

[**ChecksAll**](ChecksAll.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheck

> CheckWithoutIDGET GetCheck(ctx, cid).ExtendedTags(extendedTags).Execute()

Returns a single transaction check.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cid := int64(789) // int64 | Specifies the id of the check to be fetched
	extendedTags := true // bool | Specifies whether to return an extended tags representation in the response (with type and count). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TMSChecksAPI.GetCheck(context.Background(), cid).ExtendedTags(extendedTags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TMSChecksAPI.GetCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheck`: CheckWithoutIDGET
	fmt.Fprintf(os.Stdout, "Response from `TMSChecksAPI.GetCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int64** | Specifies the id of the check to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extendedTags** | **bool** | Specifies whether to return an extended tags representation in the response (with type and count). | 

### Return type

[**CheckWithoutIDGET**](CheckWithoutIDGET.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckReportPerformance

> ReportPerformance GetCheckReportPerformance(ctx, checkId).From(from).To(to).Order(order).Resolution(resolution).IncludeUptime(includeUptime).Execute()

Returns a performance report for a single transaction check



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	checkId := int64(789) // int64 | Specifies the id of the check for which the performance report will be fetched
	from := time.Now() // time.Time | Start time of period. The format is `RFC 3339` (properly URL-encoded!). The default value is 10 times the resolution (10 hours, 10 day, 10 weeks) earlier than `to`. The value is extended to the nearest hour, day or week, depending on the 'resolution' parameter and configured time zone of the account. (optional)
	to := time.Now() // time.Time | End time of period. The format is `RFC 3339` (properly URL-encoded!). The default value is the current time. The value is extended to the nearest hour, day or week, depending on the 'resolution' parameter and configured time zone of the account. (optional)
	order := "["desc"]" // string | Sorting order of outages. Ascending or descending (optional) (default to "asc")
	resolution := "day" // string | Size of an interval for which the results are calculated. For the `hour` resolution, the max allowed period is one week and 1 day. For the `day` resolution, the max allowed period is 6 months and 1 day. (optional) (default to "hour")
	includeUptime := true // bool | Include uptime information. Adds field downtime, uptime, and unmonitored to the interval array objects. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TMSChecksAPI.GetCheckReportPerformance(context.Background(), checkId).From(from).To(to).Order(order).Resolution(resolution).IncludeUptime(includeUptime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TMSChecksAPI.GetCheckReportPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheckReportPerformance`: ReportPerformance
	fmt.Fprintf(os.Stdout, "Response from `TMSChecksAPI.GetCheckReportPerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **int64** | Specifies the id of the check for which the performance report will be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckReportPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start time of period. The format is &#x60;RFC 3339&#x60; (properly URL-encoded!). The default value is 10 times the resolution (10 hours, 10 day, 10 weeks) earlier than &#x60;to&#x60;. The value is extended to the nearest hour, day or week, depending on the &#39;resolution&#39; parameter and configured time zone of the account. | 
 **to** | **time.Time** | End time of period. The format is &#x60;RFC 3339&#x60; (properly URL-encoded!). The default value is the current time. The value is extended to the nearest hour, day or week, depending on the &#39;resolution&#39; parameter and configured time zone of the account. | 
 **order** | **string** | Sorting order of outages. Ascending or descending | [default to &quot;asc&quot;]
 **resolution** | **string** | Size of an interval for which the results are calculated. For the &#x60;hour&#x60; resolution, the max allowed period is one week and 1 day. For the &#x60;day&#x60; resolution, the max allowed period is 6 months and 1 day. | [default to &quot;hour&quot;]
 **includeUptime** | **bool** | Include uptime information. Adds field downtime, uptime, and unmonitored to the interval array objects. | [default to false]

### Return type

[**ReportPerformance**](ReportPerformance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckReportStatus

> ReportStatusSingle GetCheckReportStatus(ctx, checkId).From(from).To(to).Order(order).Execute()

Returns a status change report for a single transaction check



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	checkId := int64(789) // int64 | Specifies the id of the check for which the status change report will be fetched
	from := time.Now() // time.Time | Start time of period. The format is `RFC 3339` (properly URL-encoded!). The default value is one week earlier than `to` (optional)
	to := time.Now() // time.Time | End time of period. The format is `RFC 3339` (properly URL-encoded!). The default value is the current time (optional)
	order := "["desc"]" // string | Sorting order of outages. Ascending or descending (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TMSChecksAPI.GetCheckReportStatus(context.Background(), checkId).From(from).To(to).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TMSChecksAPI.GetCheckReportStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheckReportStatus`: ReportStatusSingle
	fmt.Fprintf(os.Stdout, "Response from `TMSChecksAPI.GetCheckReportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **int64** | Specifies the id of the check for which the status change report will be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckReportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Start time of period. The format is &#x60;RFC 3339&#x60; (properly URL-encoded!). The default value is one week earlier than &#x60;to&#x60; | 
 **to** | **time.Time** | End time of period. The format is &#x60;RFC 3339&#x60; (properly URL-encoded!). The default value is the current time | 
 **order** | **string** | Sorting order of outages. Ascending or descending | [default to &quot;asc&quot;]

### Return type

[**ReportStatusSingle**](ReportStatusSingle.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckReportStatusAll

> ReportStatusAll GetCheckReportStatusAll(ctx).From(from).To(to).Order(order).Limit(limit).Offset(offset).OmitEmpty(omitEmpty).Execute()

Returns a status change report for all transaction checks in the current organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	from := time.Now() // time.Time | Start time of period. The format is `RFC 3339` (properly URL-encoded!). The default value is one week earlier than `to` (optional)
	to := time.Now() // time.Time | End time of period. The format is `RFC 3339` (properly URL-encoded!). The default value is the current time (optional)
	order := "["desc"]" // string | Sorting order of outages. Ascending or descending (optional) (default to "asc")
	limit := "limit_example" // string | Limit of returned checks (optional) (default to "100")
	offset := "offset_example" // string | Offset of returned checks (optional) (default to "0")
	omitEmpty := true // bool | Omits checks without any results within specified time (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TMSChecksAPI.GetCheckReportStatusAll(context.Background()).From(from).To(to).Order(order).Limit(limit).Offset(offset).OmitEmpty(omitEmpty).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TMSChecksAPI.GetCheckReportStatusAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCheckReportStatusAll`: ReportStatusAll
	fmt.Fprintf(os.Stdout, "Response from `TMSChecksAPI.GetCheckReportStatusAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckReportStatusAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | Start time of period. The format is &#x60;RFC 3339&#x60; (properly URL-encoded!). The default value is one week earlier than &#x60;to&#x60; | 
 **to** | **time.Time** | End time of period. The format is &#x60;RFC 3339&#x60; (properly URL-encoded!). The default value is the current time | 
 **order** | **string** | Sorting order of outages. Ascending or descending | [default to &quot;asc&quot;]
 **limit** | **string** | Limit of returned checks | [default to &quot;100&quot;]
 **offset** | **string** | Offset of returned checks | [default to &quot;0&quot;]
 **omitEmpty** | **bool** | Omits checks without any results within specified time | [default to false]

### Return type

[**ReportStatusAll**](ReportStatusAll.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCheck

> CheckWithoutIDGET ModifyCheck(ctx, cid).CheckWithoutIDPUT(checkWithoutIDPUT).Execute()

Modify settings for transaction check.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cid := int64(789) // int64 | Specifies the id of the check to be modified
	checkWithoutIDPUT := *openapiclient.NewCheckWithoutIDPUT() // CheckWithoutIDPUT | Specifies the data to be modified for the check

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TMSChecksAPI.ModifyCheck(context.Background(), cid).CheckWithoutIDPUT(checkWithoutIDPUT).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TMSChecksAPI.ModifyCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyCheck`: CheckWithoutIDGET
	fmt.Fprintf(os.Stdout, "Response from `TMSChecksAPI.ModifyCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int64** | Specifies the id of the check to be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkWithoutIDPUT** | [**CheckWithoutIDPUT**](CheckWithoutIDPUT.md) | Specifies the data to be modified for the check | 

### Return type

[**CheckWithoutIDGET**](CheckWithoutIDGET.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

