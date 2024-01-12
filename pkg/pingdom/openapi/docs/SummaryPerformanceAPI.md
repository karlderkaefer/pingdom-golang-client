# \SummaryPerformanceAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SummaryPerformanceCheckidGet**](SummaryPerformanceAPI.md#SummaryPerformanceCheckidGet) | **Get** /summary.performance/{checkid} | For a given interval return a list of subintervals



## SummaryPerformanceCheckidGet

> SummaryPerformanceRespAttrs SummaryPerformanceCheckidGet(ctx, checkid).From(from).To(to).Resolution(resolution).Includeuptime(includeuptime).Probes(probes).Order(order).Execute()

For a given interval return a list of subintervals



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
	from := int32(56) // int32 | Start time of period. Format is UNIX timestamp. Default value is 10 intervals earlier than `to`. (optional)
	to := int32(56) // int32 | End time of period. Format is UNIX timestamp. Default value is the current time. (optional)
	resolution := openapiclient.summary.performance_resolution("hour") // SummaryPerformanceResolution | Interval size (optional) (default to "hour")
	includeuptime := true // bool | Include uptime information. (optional) (default to false)
	probes := "probes_example" // string | Filter to only use results from a list of probes. Format is a comma separated list of probe identifiers. Can not be used if includeuptime is set to true. Also note that this can cause intervals to be omitted, since there may be no results from the desired probes in them. By deafult results from all probes are returned. (optional)
	order := openapiclient.summary.performance_order("asc") // SummaryPerformanceOrder | Sorting order of sub intervals. Ascending or descending. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummaryPerformanceAPI.SummaryPerformanceCheckidGet(context.Background(), checkid).From(from).To(to).Resolution(resolution).Includeuptime(includeuptime).Probes(probes).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummaryPerformanceAPI.SummaryPerformanceCheckidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryPerformanceCheckidGet`: SummaryPerformanceRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `SummaryPerformanceAPI.SummaryPerformanceCheckidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSummaryPerformanceCheckidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | Start time of period. Format is UNIX timestamp. Default value is 10 intervals earlier than &#x60;to&#x60;. | 
 **to** | **int32** | End time of period. Format is UNIX timestamp. Default value is the current time. | 
 **resolution** | [**SummaryPerformanceResolution**](SummaryPerformanceResolution.md) | Interval size | [default to &quot;hour&quot;]
 **includeuptime** | **bool** | Include uptime information. | [default to false]
 **probes** | **string** | Filter to only use results from a list of probes. Format is a comma separated list of probe identifiers. Can not be used if includeuptime is set to true. Also note that this can cause intervals to be omitted, since there may be no results from the desired probes in them. By deafult results from all probes are returned. | 
 **order** | [**SummaryPerformanceOrder**](SummaryPerformanceOrder.md) | Sorting order of sub intervals. Ascending or descending. | [default to &quot;asc&quot;]

### Return type

[**SummaryPerformanceRespAttrs**](SummaryPerformanceRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

