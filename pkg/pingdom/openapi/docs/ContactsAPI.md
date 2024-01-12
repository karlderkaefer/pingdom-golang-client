# \ContactsAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertingContactsContactidDelete**](ContactsAPI.md#AlertingContactsContactidDelete) | **Delete** /alerting/contacts/{contactid} | Deletes a contact with its contact methods
[**AlertingContactsContactidGet**](ContactsAPI.md#AlertingContactsContactidGet) | **Get** /alerting/contacts/{contactid} | Returns a contact with its contact methods
[**AlertingContactsContactidPut**](ContactsAPI.md#AlertingContactsContactidPut) | **Put** /alerting/contacts/{contactid} | Update a contact
[**AlertingContactsGet**](ContactsAPI.md#AlertingContactsGet) | **Get** /alerting/contacts | Returns a list of all contacts
[**AlertingContactsPost**](ContactsAPI.md#AlertingContactsPost) | **Post** /alerting/contacts | Creates a new contact



## AlertingContactsContactidDelete

> AlertingContactsContactidDelete200Response AlertingContactsContactidDelete(ctx, contactid).Execute()

Deletes a contact with its contact methods



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
	contactid := int32(56) // int32 | ID of contact to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.AlertingContactsContactidDelete(context.Background(), contactid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.AlertingContactsContactidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingContactsContactidDelete`: AlertingContactsContactidDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.AlertingContactsContactidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactid** | **int32** | ID of contact to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingContactsContactidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertingContactsContactidDelete200Response**](AlertingContactsContactidDelete200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingContactsContactidGet

> Contact AlertingContactsContactidGet(ctx, contactid).Execute()

Returns a contact with its contact methods



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
	contactid := int32(56) // int32 | ID of contact to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.AlertingContactsContactidGet(context.Background(), contactid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.AlertingContactsContactidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingContactsContactidGet`: Contact
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.AlertingContactsContactidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactid** | **int32** | ID of contact to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingContactsContactidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingContactsContactidPut

> Contact AlertingContactsContactidPut(ctx, contactid).UpdateContact(updateContact).Execute()

Update a contact



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
	contactid := int32(56) // int32 | ID of contact to be updated
	updateContact := *openapiclient.NewUpdateContact("Name_example", false, *openapiclient.NewCreateContactNotificationTargets()) // UpdateContact | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.AlertingContactsContactidPut(context.Background(), contactid).UpdateContact(updateContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.AlertingContactsContactidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingContactsContactidPut`: Contact
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.AlertingContactsContactidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contactid** | **int32** | ID of contact to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingContactsContactidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContact** | [**UpdateContact**](UpdateContact.md) |  | 

### Return type

[**Contact**](Contact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingContactsGet

> ContactsList AlertingContactsGet(ctx).Execute()

Returns a list of all contacts



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.AlertingContactsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.AlertingContactsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingContactsGet`: ContactsList
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.AlertingContactsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingContactsGetRequest struct via the builder pattern


### Return type

[**ContactsList**](ContactsList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingContactsPost

> AlertingContactsPost200Response AlertingContactsPost(ctx).CreateContact(createContact).Execute()

Creates a new contact



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
	createContact := *openapiclient.NewCreateContact("Name_example", *openapiclient.NewCreateContactNotificationTargets()) // CreateContact | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContactsAPI.AlertingContactsPost(context.Background()).CreateContact(createContact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContactsAPI.AlertingContactsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingContactsPost`: AlertingContactsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ContactsAPI.AlertingContactsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingContactsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContact** | [**CreateContact**](CreateContact.md) |  | 

### Return type

[**AlertingContactsPost200Response**](AlertingContactsPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

