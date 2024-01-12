# \TeamsAPI

All URIs are relative to *https://api.pingdom.com/api/3.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertingTeamsGet**](TeamsAPI.md#AlertingTeamsGet) | **Get** /alerting/teams | 
[**AlertingTeamsPost**](TeamsAPI.md#AlertingTeamsPost) | **Post** /alerting/teams | Creates a new team
[**AlertingTeamsTeamidDelete**](TeamsAPI.md#AlertingTeamsTeamidDelete) | **Delete** /alerting/teams/{teamid} | 
[**AlertingTeamsTeamidGet**](TeamsAPI.md#AlertingTeamsTeamidGet) | **Get** /alerting/teams/{teamid} | 
[**AlertingTeamsTeamidPut**](TeamsAPI.md#AlertingTeamsTeamidPut) | **Put** /alerting/teams/{teamid} | 



## AlertingTeamsGet

> Teams AlertingTeamsGet(ctx).Execute()





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
	resp, r, err := apiClient.TeamsAPI.AlertingTeamsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AlertingTeamsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingTeamsGet`: Teams
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AlertingTeamsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingTeamsGetRequest struct via the builder pattern


### Return type

[**Teams**](Teams.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingTeamsPost

> AlertingTeamsPost200Response AlertingTeamsPost(ctx).CreateTeam(createTeam).Execute()

Creates a new team



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
	createTeam := *openapiclient.NewCreateTeam("Name_example", []int32{int32(123)}) // CreateTeam | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.AlertingTeamsPost(context.Background()).CreateTeam(createTeam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AlertingTeamsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingTeamsPost`: AlertingTeamsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AlertingTeamsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertingTeamsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeam** | [**CreateTeam**](CreateTeam.md) |  | 

### Return type

[**AlertingTeamsPost200Response**](AlertingTeamsPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingTeamsTeamidDelete

> AlertingTeamsTeamidDelete200Response AlertingTeamsTeamidDelete(ctx, teamid).Execute()





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
	teamid := int32(56) // int32 | ID of the team to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.AlertingTeamsTeamidDelete(context.Background(), teamid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AlertingTeamsTeamidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingTeamsTeamidDelete`: AlertingTeamsTeamidDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AlertingTeamsTeamidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamid** | **int32** | ID of the team to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingTeamsTeamidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertingTeamsTeamidDelete200Response**](AlertingTeamsTeamidDelete200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingTeamsTeamidGet

> TeamID AlertingTeamsTeamidGet(ctx, teamid).Execute()





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
	teamid := int32(56) // int32 | ID of the team to be retrieved

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.AlertingTeamsTeamidGet(context.Background(), teamid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AlertingTeamsTeamidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingTeamsTeamidGet`: TeamID
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AlertingTeamsTeamidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamid** | **int32** | ID of the team to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingTeamsTeamidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TeamID**](TeamID.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertingTeamsTeamidPut

> TeamID AlertingTeamsTeamidPut(ctx, teamid).UpdateTeam(updateTeam).Execute()





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
	teamid := int32(56) // int32 | ID of the team to be retrieved
	updateTeam := *openapiclient.NewUpdateTeam("Name_example", []int64{int64(123)}) // UpdateTeam | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.AlertingTeamsTeamidPut(context.Background(), teamid).UpdateTeam(updateTeam).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AlertingTeamsTeamidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AlertingTeamsTeamidPut`: TeamID
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AlertingTeamsTeamidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamid** | **int32** | ID of the team to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertingTeamsTeamidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTeam** | [**UpdateTeam**](UpdateTeam.md) |  | 

### Return type

[**TeamID**](TeamID.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

