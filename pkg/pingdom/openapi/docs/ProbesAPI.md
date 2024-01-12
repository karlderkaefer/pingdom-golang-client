# \ProbesAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProbesGet**](ProbesAPI.md#ProbesGet) | **Get** /probes | Returns a list of Pingdom probe servers



## ProbesGet

> Probes ProbesGet(ctx).Limit(limit).Offset(offset).Onlyactive(onlyactive).Includedeleted(includedeleted).Execute()

Returns a list of Pingdom probe servers



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
	limit := int32(3) // int32 | Limits the number of returned probes to the specified quantity. (optional)
	offset := int32(0) // int32 | Offset for listing. (Requires limit.) (optional) (default to 0)
	onlyactive := false // bool | Return only active probes. (optional) (default to false)
	includedeleted := false // bool | Include old probes that are no longer in use. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProbesAPI.ProbesGet(context.Background()).Limit(limit).Offset(offset).Onlyactive(onlyactive).Includedeleted(includedeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProbesAPI.ProbesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProbesGet`: Probes
	fmt.Fprintf(os.Stdout, "Response from `ProbesAPI.ProbesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProbesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limits the number of returned probes to the specified quantity. | 
 **offset** | **int32** | Offset for listing. (Requires limit.) | [default to 0]
 **onlyactive** | **bool** | Return only active probes. | [default to false]
 **includedeleted** | **bool** | Include old probes that are no longer in use. | [default to false]

### Return type

[**Probes**](Probes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

