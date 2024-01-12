# \ResultsAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResultsCheckidGet**](ResultsAPI.md#ResultsCheckidGet) | **Get** /results/{checkid} | Return a list of raw test results



## ResultsCheckidGet

> ResultsRespAttrs ResultsCheckidGet(ctx, checkid).To(to).From(from).Probes(probes).Status(status).Limit(limit).Offset(offset).Includeanalysis(includeanalysis).Maxresponse(maxresponse).Minresponse(minresponse).Execute()

Return a list of raw test results



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
	to := int32(56) // int32 | End of period. Format is UNIX timestamp. Default value is current timestamp. (optional)
	from := int32(56) // int32 | Start of period. Format is UNIX timestamp. Default value is 1 day prior to `to`. (optional)
	probes := "probes_example" // string | Filter to only show results from a list of probes. Format is a comma separated list of probe identifiers (optional)
	status := "status_example" // string | Filter to only show results with specified statuses. Format is a comma separated list of (`down`, `up`, `unconfirmed`, `unknown`) (optional)
	limit := int32(56) // int32 | Number of results to show (Will be set to 1000 if the provided value is greater than 1000) (optional) (default to 1000)
	offset := int32(56) // int32 | Number of results to skip (Max value is `43200`) (optional) (default to 0)
	includeanalysis := true // bool | Attach available root cause analysis identifiers to corresponding results (optional) (default to false)
	maxresponse := int32(56) // int32 | Maximum response time (ms). If set, specified interval must not be larger than 31 days. (optional)
	minresponse := int32(56) // int32 | Minimum response time (ms). If set, specified interval must not be larger than 31 days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultsAPI.ResultsCheckidGet(context.Background(), checkid).To(to).From(from).Probes(probes).Status(status).Limit(limit).Offset(offset).Includeanalysis(includeanalysis).Maxresponse(maxresponse).Minresponse(minresponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultsAPI.ResultsCheckidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResultsCheckidGet`: ResultsRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `ResultsAPI.ResultsCheckidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResultsCheckidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **to** | **int32** | End of period. Format is UNIX timestamp. Default value is current timestamp. | 
 **from** | **int32** | Start of period. Format is UNIX timestamp. Default value is 1 day prior to &#x60;to&#x60;. | 
 **probes** | **string** | Filter to only show results from a list of probes. Format is a comma separated list of probe identifiers | 
 **status** | **string** | Filter to only show results with specified statuses. Format is a comma separated list of (&#x60;down&#x60;, &#x60;up&#x60;, &#x60;unconfirmed&#x60;, &#x60;unknown&#x60;) | 
 **limit** | **int32** | Number of results to show (Will be set to 1000 if the provided value is greater than 1000) | [default to 1000]
 **offset** | **int32** | Number of results to skip (Max value is &#x60;43200&#x60;) | [default to 0]
 **includeanalysis** | **bool** | Attach available root cause analysis identifiers to corresponding results | [default to false]
 **maxresponse** | **int32** | Maximum response time (ms). If set, specified interval must not be larger than 31 days. | 
 **minresponse** | **int32** | Minimum response time (ms). If set, specified interval must not be larger than 31 days. | 

### Return type

[**ResultsRespAttrs**](ResultsRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

