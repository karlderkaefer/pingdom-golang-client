# TracerouteData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **string** | Traceroute output | [optional] 
**Probeid** | Pointer to **int32** | Probe identifier | [optional] 
**Probedescription** | Pointer to **string** | Probe description | [optional] 

## Methods

### NewTracerouteData

`func NewTracerouteData() *TracerouteData`

NewTracerouteData instantiates a new TracerouteData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTracerouteDataWithDefaults

`func NewTracerouteDataWithDefaults() *TracerouteData`

NewTracerouteDataWithDefaults instantiates a new TracerouteData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *TracerouteData) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TracerouteData) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TracerouteData) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *TracerouteData) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetProbeid

`func (o *TracerouteData) GetProbeid() int32`

GetProbeid returns the Probeid field if non-nil, zero value otherwise.

### GetProbeidOk

`func (o *TracerouteData) GetProbeidOk() (*int32, bool)`

GetProbeidOk returns a tuple with the Probeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeid

`func (o *TracerouteData) SetProbeid(v int32)`

SetProbeid sets Probeid field to given value.

### HasProbeid

`func (o *TracerouteData) HasProbeid() bool`

HasProbeid returns a boolean if a field has been set.

### GetProbedescription

`func (o *TracerouteData) GetProbedescription() string`

GetProbedescription returns the Probedescription field if non-nil, zero value otherwise.

### GetProbedescriptionOk

`func (o *TracerouteData) GetProbedescriptionOk() (*string, bool)`

GetProbedescriptionOk returns a tuple with the Probedescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbedescription

`func (o *TracerouteData) SetProbedescription(v string)`

SetProbedescription sets Probedescription field to given value.

### HasProbedescription

`func (o *TracerouteData) HasProbedescription() bool`

HasProbedescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


