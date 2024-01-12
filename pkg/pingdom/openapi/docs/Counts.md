# Counts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of checks | [optional] 
**Limited** | Pointer to **int32** | Number of checks after tags filter was applied | [optional] 
**Filtered** | Pointer to **int32** | Number of checks after limit was applied | [optional] 

## Methods

### NewCounts

`func NewCounts() *Counts`

NewCounts instantiates a new Counts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountsWithDefaults

`func NewCountsWithDefaults() *Counts`

NewCountsWithDefaults instantiates a new Counts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *Counts) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Counts) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Counts) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Counts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLimited

`func (o *Counts) GetLimited() int32`

GetLimited returns the Limited field if non-nil, zero value otherwise.

### GetLimitedOk

`func (o *Counts) GetLimitedOk() (*int32, bool)`

GetLimitedOk returns a tuple with the Limited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimited

`func (o *Counts) SetLimited(v int32)`

SetLimited sets Limited field to given value.

### HasLimited

`func (o *Counts) HasLimited() bool`

HasLimited returns a boolean if a field has been set.

### GetFiltered

`func (o *Counts) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *Counts) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *Counts) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### HasFiltered

`func (o *Counts) HasFiltered() bool`

HasFiltered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


