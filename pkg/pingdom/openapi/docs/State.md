# State

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Interval status | [optional] 
**From** | Pointer to **time.Time** | Interval start. The format is &#x60;RFC 3339&#x60; | [optional] 
**To** | Pointer to **time.Time** | Interval end. The format is &#x60;RFC 3339&#x60; | [optional] 
**ErrorInStep** | Pointer to **int32** | Number of step in which an error has occured (only if &#x60;status&#x60; is &#x60;down&#x60;) | [optional] 
**Message** | Pointer to **string** | Error message for the step that is failing (only if &#x60;status&#x60; is &#x60;down&#x60;) | [optional] 

## Methods

### NewState

`func NewState() *State`

NewState instantiates a new State object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateWithDefaults

`func NewStateWithDefaults() *State`

NewStateWithDefaults instantiates a new State object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *State) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *State) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *State) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *State) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFrom

`func (o *State) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *State) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *State) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *State) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *State) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *State) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *State) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *State) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetErrorInStep

`func (o *State) GetErrorInStep() int32`

GetErrorInStep returns the ErrorInStep field if non-nil, zero value otherwise.

### GetErrorInStepOk

`func (o *State) GetErrorInStepOk() (*int32, bool)`

GetErrorInStepOk returns a tuple with the ErrorInStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorInStep

`func (o *State) SetErrorInStep(v int32)`

SetErrorInStep sets ErrorInStep field to given value.

### HasErrorInStep

`func (o *State) HasErrorInStep() bool`

HasErrorInStep returns a boolean if a field has been set.

### GetMessage

`func (o *State) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *State) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *State) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *State) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


