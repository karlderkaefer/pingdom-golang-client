# \SummaryOutageAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SummaryOutageCheckidGet**](SummaryOutageAPI.md#SummaryOutageCheckidGet) | **Get** /summary.outage/{checkid} | Get a list of status changes for a specified check



## SummaryOutageCheckidGet

> SummaryOutageRespAttrs SummaryOutageCheckidGet(ctx, checkid).From(from).To(to).Order(order).Execute()

Get a list of status changes for a specified check



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
	from := int32(56) // int32 | Start time of period. Format is UNIX timestamp. Default value is one week earlier than `to`. (optional)
	to := int32(56) // int32 | End time of period. Format is UNIX timestamp. Default value is the current time. (optional)
	order := openapiclient.summary.outage_order("asc") // SummaryOutageOrder | Sorting order of outages. Ascending or descending. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummaryOutageAPI.SummaryOutageCheckidGet(context.Background(), checkid).From(from).To(to).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummaryOutageAPI.SummaryOutageCheckidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryOutageCheckidGet`: SummaryOutageRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `SummaryOutageAPI.SummaryOutageCheckidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSummaryOutageCheckidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | Start time of period. Format is UNIX timestamp. Default value is one week earlier than &#x60;to&#x60;. | 
 **to** | **int32** | End time of period. Format is UNIX timestamp. Default value is the current time. | 
 **order** | [**SummaryOutageOrder**](SummaryOutageOrder.md) | Sorting order of outages. Ascending or descending. | [default to &quot;asc&quot;]

### Return type

[**SummaryOutageRespAttrs**](SummaryOutageRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

