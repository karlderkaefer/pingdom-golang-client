# ChecksPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paused** | Pointer to **bool** | Use value: true to pause the check(s) and value: false to unpause it(them). | [optional] 
**Resolution** | Pointer to **int32** |  | [optional] 
**Checkids** | Pointer to **string** | Comma-separated list of identifiers for checks to be modified. Invalid check identifiers will be ignored. Default: all checks  | [optional] 

## Methods

### NewChecksPutRequest

`func NewChecksPutRequest() *ChecksPutRequest`

NewChecksPutRequest instantiates a new ChecksPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksPutRequestWithDefaults

`func NewChecksPutRequestWithDefaults() *ChecksPutRequest`

NewChecksPutRequestWithDefaults instantiates a new ChecksPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaused

`func (o *ChecksPutRequest) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ChecksPutRequest) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ChecksPutRequest) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ChecksPutRequest) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetResolution

`func (o *ChecksPutRequest) GetResolution() int32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ChecksPutRequest) GetResolutionOk() (*int32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ChecksPutRequest) SetResolution(v int32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ChecksPutRequest) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetCheckids

`func (o *ChecksPutRequest) GetCheckids() string`

GetCheckids returns the Checkids field if non-nil, zero value otherwise.

### GetCheckidsOk

`func (o *ChecksPutRequest) GetCheckidsOk() (*string, bool)`

GetCheckidsOk returns a tuple with the Checkids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckids

`func (o *ChecksPutRequest) SetCheckids(v string)`

SetCheckids sets Checkids field to given value.

### HasCheckids

`func (o *ChecksPutRequest) HasCheckids() bool`

HasCheckids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


