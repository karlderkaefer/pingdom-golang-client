# \MaintenanceAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaintenanceDelete**](MaintenanceAPI.md#MaintenanceDelete) | **Delete** /maintenance | Delete multiple maintenance windows.
[**MaintenanceGet**](MaintenanceAPI.md#MaintenanceGet) | **Get** /maintenance | 
[**MaintenanceIdDelete**](MaintenanceAPI.md#MaintenanceIdDelete) | **Delete** /maintenance/{id} | Delete the maintenance window.
[**MaintenanceIdGet**](MaintenanceAPI.md#MaintenanceIdGet) | **Get** /maintenance/{id} | 
[**MaintenanceIdPut**](MaintenanceAPI.md#MaintenanceIdPut) | **Put** /maintenance/{id} | 
[**MaintenancePost**](MaintenanceAPI.md#MaintenancePost) | **Post** /maintenance | 



## MaintenanceDelete

> MaintenanceDeleteRespAttrs MaintenanceDelete(ctx).Maintenanceids(maintenanceids).Execute()

Delete multiple maintenance windows.



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
	maintenanceids := []int32{int32(123)} // []int32 | Comma-separated list of identifiers of maintenance windows to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceAPI.MaintenanceDelete(context.Background()).Maintenanceids(maintenanceids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.MaintenanceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceDelete`: MaintenanceDeleteRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceAPI.MaintenanceDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceids** | **[]int32** | Comma-separated list of identifiers of maintenance windows to be deleted. | 

### Return type

[**MaintenanceDeleteRespAttrs**](MaintenanceDeleteRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceGet

> MaintenanceRespAttrs MaintenanceGet(ctx).Limit(limit).Offset(offset).Orderby(orderby).Order(order).Execute()





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
	limit := int32(56) // int32 | Count of items to list. (optional)
	offset := int32(56) // int32 | Offset of the list. (optional)
	orderby := openapiclient.maintenance_orderby("description") // MaintenanceOrderby | Order by the specific property of the maintenance window. (optional)
	order := openapiclient.maintenance_order("asc") // MaintenanceOrder | Order a-z for asc z-a for desc. Works only if orderby is specified. (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceAPI.MaintenanceGet(context.Background()).Limit(limit).Offset(offset).Orderby(orderby).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.MaintenanceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceGet`: MaintenanceRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceAPI.MaintenanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Count of items to list. | 
 **offset** | **int32** | Offset of the list. | 
 **orderby** | [**MaintenanceOrderby**](MaintenanceOrderby.md) | Order by the specific property of the maintenance window. | 
 **order** | [**MaintenanceOrder**](MaintenanceOrder.md) | Order a-z for asc z-a for desc. Works only if orderby is specified. | [default to &quot;asc&quot;]

### Return type

[**MaintenanceRespAttrs**](MaintenanceRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceIdDelete

> MaintenanceIdDeleteRespAttrs MaintenanceIdDelete(ctx, id).Execute()

Delete the maintenance window.



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
	id := int32(3) // int32 | id of maintenance window

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceAPI.MaintenanceIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.MaintenanceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceIdDelete`: MaintenanceIdDeleteRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceAPI.MaintenanceIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id of maintenance window | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintenanceIdDeleteRespAttrs**](MaintenanceIdDeleteRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceIdGet

> MaintenanceIdRespAttrs MaintenanceIdGet(ctx, id).Execute()





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
	id := int32(3) // int32 | id of maintenance window

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceAPI.MaintenanceIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.MaintenanceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceIdGet`: MaintenanceIdRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceAPI.MaintenanceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id of maintenance window | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintenanceIdRespAttrs**](MaintenanceIdRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceIdPut

> MaintenanceIdPutRespAttrs MaintenanceIdPut(ctx, id).MaintenanceIdPut(maintenanceIdPut).Execute()





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
	id := int32(3) // int32 | id of maintenance window
	maintenanceIdPut := *openapiclient.NewMaintenanceIdPut() // MaintenanceIdPut | Description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceAPI.MaintenanceIdPut(context.Background(), id).MaintenanceIdPut(maintenanceIdPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.MaintenanceIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceIdPut`: MaintenanceIdPutRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceAPI.MaintenanceIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | id of maintenance window | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceIdPut** | [**MaintenanceIdPut**](MaintenanceIdPut.md) | Description | 

### Return type

[**MaintenanceIdPutRespAttrs**](MaintenanceIdPutRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenancePost

> MaintenancePostRespAttrs MaintenancePost(ctx).MaintenancePost(maintenancePost).Execute()





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
	maintenancePost := *openapiclient.NewMaintenancePost("New maintenance window.", int32(1500471702), int32(1500475302)) // MaintenancePost | Description

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceAPI.MaintenancePost(context.Background()).MaintenancePost(maintenancePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceAPI.MaintenancePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenancePost`: MaintenancePostRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceAPI.MaintenancePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaintenancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenancePost** | [**MaintenancePost**](MaintenancePost.md) | Description | 

### Return type

[**MaintenancePostRespAttrs**](MaintenancePostRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

