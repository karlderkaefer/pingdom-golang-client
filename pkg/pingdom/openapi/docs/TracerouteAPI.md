# \TracerouteAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TracerouteGet**](TracerouteAPI.md#TracerouteGet) | **Get** /traceroute | Perform a traceroute



## TracerouteGet

> Traceroute TracerouteGet(ctx).Host(host).Probeid(probeid).Execute()

Perform a traceroute



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
	host := "pingdom.com" // string | Target host.
	probeid := int32(23) // int32 | Probe identifier. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracerouteAPI.TracerouteGet(context.Background()).Host(host).Probeid(probeid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracerouteAPI.TracerouteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TracerouteGet`: Traceroute
	fmt.Fprintf(os.Stdout, "Response from `TracerouteAPI.TracerouteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTracerouteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** | Target host. | 
 **probeid** | **int32** | Probe identifier. | 

### Return type

[**Traceroute**](Traceroute.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

