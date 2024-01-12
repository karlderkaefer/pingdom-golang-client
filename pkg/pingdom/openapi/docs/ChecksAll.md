# ChecksAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to [**[]CheckGeneral**](CheckGeneral.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 

## Methods

### NewChecksAll

`func NewChecksAll() *ChecksAll`

NewChecksAll instantiates a new ChecksAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChecksAllWithDefaults

`func NewChecksAllWithDefaults() *ChecksAll`

NewChecksAllWithDefaults instantiates a new ChecksAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *ChecksAll) GetChecks() []CheckGeneral`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ChecksAll) GetChecksOk() (*[]CheckGeneral, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ChecksAll) SetChecks(v []CheckGeneral)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ChecksAll) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetLimit

`func (o *ChecksAll) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ChecksAll) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ChecksAll) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ChecksAll) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ChecksAll) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ChecksAll) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ChecksAll) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ChecksAll) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


