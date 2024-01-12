# \SingleAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SingleGet**](SingleAPI.md#SingleGet) | **Get** /single | Performs a single check.



## SingleGet

> SingleResp SingleGet(ctx).QueryParameters(queryParameters).Execute()

Performs a single check.



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
	queryParameters := openapiclient._single_get_Query_Parameters_parameter{DNS: openapiclient.NewDNS("pingdom.com", "Type_example", "Expectedip_example", "Nameserver_example")} // SingleGetQueryParametersParameter | Query Parameters to chose (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SingleAPI.SingleGet(context.Background()).QueryParameters(queryParameters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SingleAPI.SingleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SingleGet`: SingleResp
	fmt.Fprintf(os.Stdout, "Response from `SingleAPI.SingleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSingleGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryParameters** | [**SingleGetQueryParametersParameter**](SingleGetQueryParametersParameter.md) | Query Parameters to chose | 

### Return type

[**SingleResp**](SingleResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

