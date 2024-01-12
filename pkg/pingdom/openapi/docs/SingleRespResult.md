# SingleRespResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Test result status (up, down) | [optional] 
**Responsetime** | Pointer to **int32** | Response time in milliseconds | [optional] 
**Statusdesc** | Pointer to **string** | Short status description | [optional] 
**Statusdesclong** | Pointer to **string** | Long status description | [optional] 
**Probeid** | Pointer to **int32** | Probe identifier | [optional] 
**Probedesc** | Pointer to **string** | Probe description | [optional] 

## Methods

### NewSingleRespResult

`func NewSingleRespResult() *SingleRespResult`

NewSingleRespResult instantiates a new SingleRespResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleRespResultWithDefaults

`func NewSingleRespResultWithDefaults() *SingleRespResult`

NewSingleRespResultWithDefaults instantiates a new SingleRespResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SingleRespResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SingleRespResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SingleRespResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SingleRespResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResponsetime

`func (o *SingleRespResult) GetResponsetime() int32`

GetResponsetime returns the Responsetime field if non-nil, zero value otherwise.

### GetResponsetimeOk

`func (o *SingleRespResult) GetResponsetimeOk() (*int32, bool)`

GetResponsetimeOk returns a tuple with the Responsetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetime

`func (o *SingleRespResult) SetResponsetime(v int32)`

SetResponsetime sets Responsetime field to given value.

### HasResponsetime

`func (o *SingleRespResult) HasResponsetime() bool`

HasResponsetime returns a boolean if a field has been set.

### GetStatusdesc

`func (o *SingleRespResult) GetStatusdesc() string`

GetStatusdesc returns the Statusdesc field if non-nil, zero value otherwise.

### GetStatusdescOk

`func (o *SingleRespResult) GetStatusdescOk() (*string, bool)`

GetStatusdescOk returns a tuple with the Statusdesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusdesc

`func (o *SingleRespResult) SetStatusdesc(v string)`

SetStatusdesc sets Statusdesc field to given value.

### HasStatusdesc

`func (o *SingleRespResult) HasStatusdesc() bool`

HasStatusdesc returns a boolean if a field has been set.

### GetStatusdesclong

`func (o *SingleRespResult) GetStatusdesclong() string`

GetStatusdesclong returns the Statusdesclong field if non-nil, zero value otherwise.

### GetStatusdesclongOk

`func (o *SingleRespResult) GetStatusdesclongOk() (*string, bool)`

GetStatusdesclongOk returns a tuple with the Statusdesclong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusdesclong

`func (o *SingleRespResult) SetStatusdesclong(v string)`

SetStatusdesclong sets Statusdesclong field to given value.

### HasStatusdesclong

`func (o *SingleRespResult) HasStatusdesclong() bool`

HasStatusdesclong returns a boolean if a field has been set.

### GetProbeid

`func (o *SingleRespResult) GetProbeid() int32`

GetProbeid returns the Probeid field if non-nil, zero value otherwise.

### GetProbeidOk

`func (o *SingleRespResult) GetProbeidOk() (*int32, bool)`

GetProbeidOk returns a tuple with the Probeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeid

`func (o *SingleRespResult) SetProbeid(v int32)`

SetProbeid sets Probeid field to given value.

### HasProbeid

`func (o *SingleRespResult) HasProbeid() bool`

HasProbeid returns a boolean if a field has been set.

### GetProbedesc

`func (o *SingleRespResult) GetProbedesc() string`

GetProbedesc returns the Probedesc field if non-nil, zero value otherwise.

### GetProbedescOk

`func (o *SingleRespResult) GetProbedescOk() (*string, bool)`

GetProbedescOk returns a tuple with the Probedesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbedesc

`func (o *SingleRespResult) SetProbedesc(v string)`

SetProbedesc sets Probedesc field to given value.

### HasProbedesc

`func (o *SingleRespResult) HasProbedesc() bool`

HasProbedesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


