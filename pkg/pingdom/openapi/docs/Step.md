# Step

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to [**StepArgs**](StepArgs.md) |  | [optional] 
**Fn** | Pointer to **string** | Operation to be done | [optional] 

## Methods

### NewStep

`func NewStep() *Step`

NewStep instantiates a new Step object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStepWithDefaults

`func NewStepWithDefaults() *Step`

NewStepWithDefaults instantiates a new Step object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *Step) GetArgs() StepArgs`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *Step) GetArgsOk() (*StepArgs, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *Step) SetArgs(v StepArgs)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *Step) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetFn

`func (o *Step) GetFn() string`

GetFn returns the Fn field if non-nil, zero value otherwise.

### GetFnOk

`func (o *Step) GetFnOk() (*string, bool)`

GetFnOk returns a tuple with the Fn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFn

`func (o *Step) SetFn(v string)`

SetFn sets Fn field to given value.

### HasFn

`func (o *Step) HasFn() bool`

HasFn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


