# \ActionsAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionsGet**](ActionsAPI.md#ActionsGet) | **Get** /actions | Returns a list of actions alerts.



## ActionsGet

> ActionsAlertsEntry ActionsGet(ctx).From(from).To(to).Limit(limit).Offset(offset).Checkids(checkids).Userids(userids).Status(status).Via(via).Execute()

Returns a list of actions alerts.



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
	from := int32(56) // int32 | Only include actions generated later than this timestamp. Format is UNIX time. (optional)
	to := int32(56) // int32 | Only include actions generated prior to this timestamp. Format is UNIX time. (optional)
	limit := int32(56) // int32 | Limits the number of returned results to the specified quantity. (optional) (default to 100)
	offset := int32(56) // int32 | Offset for listing. (optional) (default to 0)
	checkids := "checkids_example" // string | Comma-separated list of check identifiers. Limit results to actions generated from these checks. Default: all checks. (optional)
	userids := "userids_example" // string | Comma-separated list of user identifiers. Limit results to actions sent to these users. Default: all users. (optional)
	status := "status_example" // string | Comma-separated list of statuses. Limit results to actions with these statuses. Default: all statuses. (optional)
	via := "via_example" // string | Comma-separated list of via mediums. Limit results to actions with these mediums. Default: all mediums. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActionsAPI.ActionsGet(context.Background()).From(from).To(to).Limit(limit).Offset(offset).Checkids(checkids).Userids(userids).Status(status).Via(via).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActionsAPI.ActionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionsGet`: ActionsAlertsEntry
	fmt.Fprintf(os.Stdout, "Response from `ActionsAPI.ActionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int32** | Only include actions generated later than this timestamp. Format is UNIX time. | 
 **to** | **int32** | Only include actions generated prior to this timestamp. Format is UNIX time. | 
 **limit** | **int32** | Limits the number of returned results to the specified quantity. | [default to 100]
 **offset** | **int32** | Offset for listing. | [default to 0]
 **checkids** | **string** | Comma-separated list of check identifiers. Limit results to actions generated from these checks. Default: all checks. | 
 **userids** | **string** | Comma-separated list of user identifiers. Limit results to actions sent to these users. Default: all users. | 
 **status** | **string** | Comma-separated list of statuses. Limit results to actions with these statuses. Default: all statuses. | 
 **via** | **string** | Comma-separated list of via mediums. Limit results to actions with these mediums. Default: all mediums. | 

### Return type

[**ActionsAlertsEntry**](ActionsAlertsEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

