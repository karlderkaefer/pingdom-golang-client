# \ChecksAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChecksCheckidDelete**](ChecksAPI.md#ChecksCheckidDelete) | **Delete** /checks/{checkid} | Deletes a check.
[**ChecksCheckidGet**](ChecksAPI.md#ChecksCheckidGet) | **Get** /checks/{checkid} | Returns a detailed description of a check.
[**ChecksCheckidPut**](ChecksAPI.md#ChecksCheckidPut) | **Put** /checks/{checkid} | Modify settings for a check.
[**ChecksDelete**](ChecksAPI.md#ChecksDelete) | **Delete** /checks | Deletes a list of checks.
[**ChecksGet**](ChecksAPI.md#ChecksGet) | **Get** /checks | 
[**ChecksPost**](ChecksAPI.md#ChecksPost) | **Post** /checks | Creates a new check.
[**ChecksPut**](ChecksAPI.md#ChecksPut) | **Put** /checks | Pause or change resolution for multiple checks.



## ChecksCheckidDelete

> ChecksCheckidDelete200Response ChecksCheckidDelete(ctx, checkid).Execute()

Deletes a check.



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
	checkid := int32(56) // int32 | Identifier of check to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChecksAPI.ChecksCheckidDelete(context.Background(), checkid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChecksAPI.ChecksCheckidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChecksCheckidDelete`: ChecksCheckidDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ChecksAPI.ChecksCheckidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** | Identifier of check to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksCheckidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChecksCheckidDelete200Response**](ChecksCheckidDelete200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksCheckidGet

> DetailedCheck ChecksCheckidGet(ctx, checkid).IncludeTeams(includeTeams).Execute()

Returns a detailed description of a check.



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
	checkid := int32(56) // int32 | Identifier of check to be retrieved
	includeTeams := true // bool | Include team connections for check. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChecksAPI.ChecksCheckidGet(context.Background(), checkid).IncludeTeams(includeTeams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChecksAPI.ChecksCheckidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChecksCheckidGet`: DetailedCheck
	fmt.Fprintf(os.Stdout, "Response from `ChecksAPI.ChecksCheckidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** | Identifier of check to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksCheckidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTeams** | **bool** | Include team connections for check. | [default to false]

### Return type

[**DetailedCheck**](DetailedCheck.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksCheckidPut

> ChecksCheckidPut200Response ChecksCheckidPut(ctx, checkid).ModifyCheckSettings(modifyCheckSettings).Execute()

Modify settings for a check.



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
	checkid := int32(56) // int32 | Identifier of check to be updated
	modifyCheckSettings := openapiclient.ModifyCheckSettings{DnsAttributes: openapiclient.NewDnsAttributes("8.8.8.8", "104.20.90.241")} // ModifyCheckSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChecksAPI.ChecksCheckidPut(context.Background(), checkid).ModifyCheckSettings(modifyCheckSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChecksAPI.ChecksCheckidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChecksCheckidPut`: ChecksCheckidPut200Response
	fmt.Fprintf(os.Stdout, "Response from `ChecksAPI.ChecksCheckidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkid** | **int32** | Identifier of check to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiChecksCheckidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifyCheckSettings** | [**ModifyCheckSettings**](ModifyCheckSettings.md) |  | 

### Return type

[**ChecksCheckidPut200Response**](ChecksCheckidPut200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksDelete

> ChecksDelete200Response ChecksDelete(ctx).Delcheckids(delcheckids).Body(body).Execute()

Deletes a list of checks.



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
	delcheckids := []int32{int32(123)} // []int32 | Comma-separated list of identifiers for checks to be deleted.
	body := "{"delcheckids":"1,2,3"}" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChecksAPI.ChecksDelete(context.Background()).Delcheckids(delcheckids).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChecksAPI.ChecksDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChecksDelete`: ChecksDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ChecksAPI.ChecksDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChecksDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **delcheckids** | **[]int32** | Comma-separated list of identifiers for checks to be deleted. | 
 **body** | **string** |  | 

### Return type

[**ChecksDelete200Response**](ChecksDelete200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksGet

> Checks ChecksGet(ctx).Limit(limit).Offset(offset).Showencryption(showencryption).IncludeTags(includeTags).IncludeSeverity(includeSeverity).Tags(tags).Execute()





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
	limit := int32(3) // int32 | Limits the number of returned probes to the specified quantity. (Max value is 25000) (optional) (default to 25000)
	offset := int32(0) // int32 | Offset for listing. (Requires limit.) (optional) (default to 0)
	showencryption := false // bool | If set, show encryption setting for each check (optional) (default to false)
	includeTags := false // bool | Include tag list for each check. Tags can be marked as \"a\" or \"u\", for auto tagged or user tagged. (optional) (default to false)
	includeSeverity := false // bool | Include severity level for each check. (optional) (default to false)
	tags := "nginx,apache,ssh" // string | Tag list separated by commas. As an example \"nginx,apache\" would filter out all responses except those tagged nginx or apache (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChecksAPI.ChecksGet(context.Background()).Limit(limit).Offset(offset).Showencryption(showencryption).IncludeTags(includeTags).IncludeSeverity(includeSeverity).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChecksAPI.ChecksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChecksGet`: Checks
	fmt.Fprintf(os.Stdout, "Response from `ChecksAPI.ChecksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChecksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limits the number of returned probes to the specified quantity. (Max value is 25000) | [default to 25000]
 **offset** | **int32** | Offset for listing. (Requires limit.) | [default to 0]
 **showencryption** | **bool** | If set, show encryption setting for each check | [default to false]
 **includeTags** | **bool** | Include tag list for each check. Tags can be marked as \&quot;a\&quot; or \&quot;u\&quot;, for auto tagged or user tagged. | [default to false]
 **includeSeverity** | **bool** | Include severity level for each check. | [default to false]
 **tags** | **string** | Tag list separated by commas. As an example \&quot;nginx,apache\&quot; would filter out all responses except those tagged nginx or apache | 

### Return type

[**Checks**](Checks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksPost

> ChecksPost200Response ChecksPost(ctx).CreateCheck(createCheck).Execute()

Creates a new check.



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
	createCheck := openapiclient.CreateCheck{DnsAttributes: openapiclient.NewDnsAttributes("8.8.8.8", "104.20.90.241")} // CreateCheck | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChecksAPI.ChecksPost(context.Background()).CreateCheck(createCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChecksAPI.ChecksPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChecksPost`: ChecksPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ChecksAPI.ChecksPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChecksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCheck** | [**CreateCheck**](CreateCheck.md) |  | 

### Return type

[**ChecksPost200Response**](ChecksPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChecksPut

> ChecksPut200Response ChecksPut(ctx).ChecksPutRequest(checksPutRequest).Execute()

Pause or change resolution for multiple checks.



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
	checksPutRequest := *openapiclient.NewChecksPutRequest() // ChecksPutRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChecksAPI.ChecksPut(context.Background()).ChecksPutRequest(checksPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChecksAPI.ChecksPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChecksPut`: ChecksPut200Response
	fmt.Fprintf(os.Stdout, "Response from `ChecksAPI.ChecksPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChecksPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checksPutRequest** | [**ChecksPutRequest**](ChecksPutRequest.md) |  | 

### Return type

[**ChecksPut200Response**](ChecksPut200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

