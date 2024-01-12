# \MaintenanceOccurrencesAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaintenanceOccurrencesDelete**](MaintenanceOccurrencesAPI.md#MaintenanceOccurrencesDelete) | **Delete** /maintenance.occurrences | Deletes multiple maintenance occurrences
[**MaintenanceOccurrencesGet**](MaintenanceOccurrencesAPI.md#MaintenanceOccurrencesGet) | **Get** /maintenance.occurrences | Returns a list of maintenance occurrences.
[**MaintenanceOccurrencesIdDelete**](MaintenanceOccurrencesAPI.md#MaintenanceOccurrencesIdDelete) | **Delete** /maintenance.occurrences/{id} | Deletes the maintenance occurrence
[**MaintenanceOccurrencesIdGet**](MaintenanceOccurrencesAPI.md#MaintenanceOccurrencesIdGet) | **Get** /maintenance.occurrences/{id} | Gets a maintenance occurrence details
[**MaintenanceOccurrencesIdPut**](MaintenanceOccurrencesAPI.md#MaintenanceOccurrencesIdPut) | **Put** /maintenance.occurrences/{id} | Modifies a maintenance occurrence



## MaintenanceOccurrencesDelete

> MaintenanceOccurrencesDeleteRespAttrs MaintenanceOccurrencesDelete(ctx).Occurenceids(occurenceids).Execute()

Deletes multiple maintenance occurrences



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
	occurenceids := *openapiclient.NewMaintenanceOccurrencesDelete([]int32{int32(123)}) // MaintenanceOccurrencesDelete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceOccurrencesAPI.MaintenanceOccurrencesDelete(context.Background()).Occurenceids(occurenceids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceOccurrencesAPI.MaintenanceOccurrencesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceOccurrencesDelete`: MaintenanceOccurrencesDeleteRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceOccurrencesAPI.MaintenanceOccurrencesDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceOccurrencesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **occurenceids** | [**MaintenanceOccurrencesDelete**](MaintenanceOccurrencesDelete.md) |  | 

### Return type

[**MaintenanceOccurrencesDeleteRespAttrs**](MaintenanceOccurrencesDeleteRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceOccurrencesGet

> MaintenanceOccurrencesRespAttrs MaintenanceOccurrencesGet(ctx).Maintenanceid(maintenanceid).From(from).To(to).Execute()

Returns a list of maintenance occurrences.



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
	maintenanceid := int32(56) // int32 | Maintenance window identifier. (List only occurrences of a specific maintenance window.)   (optional)
	from := int32(56) // int32 | Effective from (unix timestamp). (List occurrences which are effective from the specified unix timestamp. If not specified, current timestamp is used.) (optional)
	to := int32(56) // int32 | Effective to (unix timestamp). (List occurrences which are effective to the specified unix timestamp.) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceOccurrencesAPI.MaintenanceOccurrencesGet(context.Background()).Maintenanceid(maintenanceid).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceOccurrencesAPI.MaintenanceOccurrencesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceOccurrencesGet`: MaintenanceOccurrencesRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceOccurrencesAPI.MaintenanceOccurrencesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceOccurrencesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceid** | **int32** | Maintenance window identifier. (List only occurrences of a specific maintenance window.)   | 
 **from** | **int32** | Effective from (unix timestamp). (List occurrences which are effective from the specified unix timestamp. If not specified, current timestamp is used.) | 
 **to** | **int32** | Effective to (unix timestamp). (List occurrences which are effective to the specified unix timestamp.) | 

### Return type

[**MaintenanceOccurrencesRespAttrs**](MaintenanceOccurrencesRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceOccurrencesIdDelete

> MaintenanceOccurrencesIdDeleteRespAttrs MaintenanceOccurrencesIdDelete(ctx, id).Execute()

Deletes the maintenance occurrence



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceOccurrencesAPI.MaintenanceOccurrencesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceOccurrencesAPI.MaintenanceOccurrencesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceOccurrencesIdDelete`: MaintenanceOccurrencesIdDeleteRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceOccurrencesAPI.MaintenanceOccurrencesIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceOccurrencesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintenanceOccurrencesIdDeleteRespAttrs**](MaintenanceOccurrencesIdDeleteRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceOccurrencesIdGet

> MaintenanceOccurrencesIdRespAttrs MaintenanceOccurrencesIdGet(ctx, id).Execute()

Gets a maintenance occurrence details



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceOccurrencesAPI.MaintenanceOccurrencesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceOccurrencesAPI.MaintenanceOccurrencesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceOccurrencesIdGet`: MaintenanceOccurrencesIdRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceOccurrencesAPI.MaintenanceOccurrencesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceOccurrencesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintenanceOccurrencesIdRespAttrs**](MaintenanceOccurrencesIdRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MaintenanceOccurrencesIdPut

> MaintenanceOccurrencesIdPutRespAttrs MaintenanceOccurrencesIdPut(ctx, id).MaintenanceOccurrencesIdPut(maintenanceOccurrencesIdPut).Execute()

Modifies a maintenance occurrence



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
	id := int32(56) // int32 | 
	maintenanceOccurrencesIdPut := *openapiclient.NewMaintenanceOccurrencesIdPut() // MaintenanceOccurrencesIdPut | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MaintenanceOccurrencesAPI.MaintenanceOccurrencesIdPut(context.Background(), id).MaintenanceOccurrencesIdPut(maintenanceOccurrencesIdPut).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceOccurrencesAPI.MaintenanceOccurrencesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MaintenanceOccurrencesIdPut`: MaintenanceOccurrencesIdPutRespAttrs
	fmt.Fprintf(os.Stdout, "Response from `MaintenanceOccurrencesAPI.MaintenanceOccurrencesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMaintenanceOccurrencesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **maintenanceOccurrencesIdPut** | [**MaintenanceOccurrencesIdPut**](MaintenanceOccurrencesIdPut.md) |  | 

### Return type

[**MaintenanceOccurrencesIdPutRespAttrs**](MaintenanceOccurrencesIdPutRespAttrs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

