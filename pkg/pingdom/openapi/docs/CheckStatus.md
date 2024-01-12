# CheckStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckId** | Pointer to **int64** | ID of the check | [optional] 
**Name** | Pointer to **string** | Name of the check | [optional] 
**States** | Pointer to [**[]State**](State.md) | Intervals when the check had a specific status (&#x60;up&#x60;/&#x60;down&#x60;). | [optional] 

## Methods

### NewCheckStatus

`func NewCheckStatus() *CheckStatus`

NewCheckStatus instantiates a new CheckStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckStatusWithDefaults

`func NewCheckStatusWithDefaults() *CheckStatus`

NewCheckStatusWithDefaults instantiates a new CheckStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckId

`func (o *CheckStatus) GetCheckId() int64`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *CheckStatus) GetCheckIdOk() (*int64, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *CheckStatus) SetCheckId(v int64)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *CheckStatus) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetName

`func (o *CheckStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStates

`func (o *CheckStatus) GetStates() []State`

GetStates returns the States field if non-nil, zero value otherwise.

### GetStatesOk

`func (o *CheckStatus) GetStatesOk() (*[]State, bool)`

GetStatesOk returns a tuple with the States field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStates

`func (o *CheckStatus) SetStates(v []State)`

SetStates sets States field to given value.

### HasStates

`func (o *CheckStatus) HasStates() bool`

HasStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


