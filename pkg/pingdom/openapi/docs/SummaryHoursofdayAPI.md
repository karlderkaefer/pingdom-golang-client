# \SummaryHoursofdayAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SummaryHoursofdayCheckidGet**](SummaryHoursofdayAPI.md#SummaryHoursofdayCheckidGet) | **Get** /summary.hoursofday/{checkid} | Returns the average response time for each hour.



## SummaryHoursofdayCheckidGet

> SummaryHoursofdayRespAttrs SummaryHoursofdayCheckidGet(ctx, checkid).From(from).To(to).Probes(probes).Uselocaltime(uselocaltime).Execute()

Returns the average response time for each hour.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/openapi"
)

func main() {
	checkid := int32(56) // int32 | 
	from := int32(56) // int32 | Start time of period. Format is UNIX timestamp. Default value is one week eariler than `to`. (optional)
	to := int32(56) // int32 | End time of period. Format is UNIX timestamp. Default value is current time. (optional)
	probes := "probes_example" // string | Filter to only use results from a list of probes. Format is a comma separated list of probe identifiers. By default all probes results are returned. (optional)
	uselocaltime := true // bool | If true, use the user's local time zone for results (from and to parameters should still be specified in UTC). If false, use UTC for results. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummaryHoursofdayAPI.SummaryHoursofdayCheckidGet(context.Background(), checkid).From(from).To(to).Probes(probes).Uselocaltime(uselocaltime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummaryHoursofdayAPI.SummaryHoursofdayCheckidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryHoursofdayCheckidGet`: SummaryHoursofdayRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `SummaryHoursofdayAPI.SummaryHoursofdayCheckidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSummaryHoursofdayCheckidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | Start time of period. Format is UNIX timestamp. Default value is one week eariler than &#x60;to&#x60;. | 
 **to** | **int32** | End time of period. Format is UNIX timestamp. Default value is current time. | 
 **probes** | **string** | Filter to only use results from a list of probes. Format is a comma separated list of probe identifiers. By default all probes results are returned. | 
 **uselocaltime** | **bool** | If true, use the user&#39;s local time zone for results (from and to parameters should still be specified in UTC). If false, use UTC for results. | [default to false]

### Return type

[**SummaryHoursofdayRespAttrs**](SummaryHoursofdayRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

