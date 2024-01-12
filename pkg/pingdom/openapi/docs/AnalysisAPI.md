# \AnalysisAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalysisCheckidAnalysisidGet**](AnalysisAPI.md#AnalysisCheckidAnalysisidGet) | **Get** /analysis/{checkid}/{analysisid} | Returns the raw result for a specified analysis.
[**AnalysisCheckidGet**](AnalysisAPI.md#AnalysisCheckidGet) | **Get** /analysis/{checkid} | Returns a list of the latest root cause analysis



## AnalysisCheckidAnalysisidGet

> AnalysisCheckidAnalysisidGet(ctx, checkid, analysisid).Execute()

Returns the raw result for a specified analysis.



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
	analysisid := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnalysisAPI.AnalysisCheckidAnalysisidGet(context.Background(), checkid, analysisid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.AnalysisCheckidAnalysisidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** |  | 
**analysisid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalysisCheckidAnalysisidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AnalysisCheckidGet

> AnalysisRespAttrs AnalysisCheckidGet(ctx, checkid).Limit(limit).Offset(offset).From(from).To(to).Execute()

Returns a list of the latest root cause analysis



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
	limit := int32(56) // int32 | Limits the number of returned results to the specified quantity. (optional) (default to 100)
	offset := int32(56) // int32 | Offset for listing. (Requires limit.) (optional) (default to 0)
	from := int32(56) // int32 | Return only results with timestamp of first test greater or equal to this value. Format is UNIX timestamp. (optional) (default to 0)
	to := int32(56) // int32 | Return only results with timestamp of first test less or equal to this value. Format is UNIX timestamp. Default: current timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnalysisAPI.AnalysisCheckidGet(context.Background(), checkid).Limit(limit).Offset(offset).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnalysisAPI.AnalysisCheckidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalysisCheckidGet`: AnalysisRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `AnalysisAPI.AnalysisCheckidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAnalysisCheckidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limits the number of returned results to the specified quantity. | [default to 100]
 **offset** | **int32** | Offset for listing. (Requires limit.) | [default to 0]
 **from** | **int32** | Return only results with timestamp of first test greater or equal to this value. Format is UNIX timestamp. | [default to 0]
 **to** | **int32** | Return only results with timestamp of first test less or equal to this value. Format is UNIX timestamp. Default: current timestamp | 

### Return type

[**AnalysisRespAttrs**](AnalysisRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

