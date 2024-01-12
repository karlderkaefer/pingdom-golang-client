# \SummaryAverageAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SummaryAverageCheckidGet**](SummaryAverageAPI.md#SummaryAverageCheckidGet) | **Get** /summary.average/{checkid} | Get the average time/uptime value for a specified



## SummaryAverageCheckidGet

> SummaryRespAttrs SummaryAverageCheckidGet(ctx, checkid).From(from).To(to).Probes(probes).Includeuptime(includeuptime).Bycountry(bycountry).Byprobe(byprobe).Execute()

Get the average time/uptime value for a specified



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
	from := int32(56) // int32 | Start time of period. Format is UNIX timestamp (optional) (default to 0)
	to := int32(56) // int32 | End time of period. Format is UNIX timestamp. Default is the current time (optional)
	probes := "probes_example" // string | Filter to only use results from a list of probes. Format is a comma separated list of probe identifiers. By default result from all probes are shown. (optional)
	includeuptime := true // bool | Include uptime information (optional) (default to false)
	bycountry := true // bool | Split response times into country groups (optional) (default to false)
	byprobe := true // bool | Split response times into probe groups (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SummaryAverageAPI.SummaryAverageCheckidGet(context.Background(), checkid).From(from).To(to).Probes(probes).Includeuptime(includeuptime).Bycountry(bycountry).Byprobe(byprobe).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SummaryAverageAPI.SummaryAverageCheckidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummaryAverageCheckidGet`: SummaryRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `SummaryAverageAPI.SummaryAverageCheckidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSummaryAverageCheckidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | Start time of period. Format is UNIX timestamp | [default to 0]
 **to** | **int32** | End time of period. Format is UNIX timestamp. Default is the current time | 
 **probes** | **string** | Filter to only use results from a list of probes. Format is a comma separated list of probe identifiers. By default result from all probes are shown. | 
 **includeuptime** | **bool** | Include uptime information | [default to false]
 **bycountry** | **bool** | Split response times into country groups | [default to false]
 **byprobe** | **bool** | Split response times into probe groups | [default to false]

### Return type

[**SummaryRespAttrs**](SummaryRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

