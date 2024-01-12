# \SummaryProbesAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SummaryProbesCheckidGet**](SummaryProbesAPI.md#SummaryProbesCheckidGet) | **Get** /summary.probes/{checkid} | Get a list of probes that performed tests



## SummaryProbesCheckidGet

> SummaryProbesRespAttrs SummaryProbesCheckidGet(ctx, checkid).From(from).To(to).Execute()

Get a list of probes that performed tests



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
	checkid := int32(56) // int32 | 
	from := int32(56) // int32 | Start time of period. Format is UNIX timestamp
	to := int32(56) // int32 | End time of period. Format is UNIX timestamp. The defualt value is current time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummaryProbesAPI.SummaryProbesCheckidGet(context.Background(), checkid).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummaryProbesAPI.SummaryProbesCheckidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryProbesCheckidGet`: SummaryProbesRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `SummaryProbesAPI.SummaryProbesCheckidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSummaryProbesCheckidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | Start time of period. Format is UNIX timestamp | 
 **to** | **int32** | End time of period. Format is UNIX timestamp. The defualt value is current time. | 

### Return type

[**SummaryProbesRespAttrs**](SummaryProbesRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

