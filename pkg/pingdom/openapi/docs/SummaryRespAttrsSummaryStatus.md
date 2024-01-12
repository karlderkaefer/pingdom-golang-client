# SummaryRespAttrsSummaryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Totalup** | Pointer to **int32** | Total uptime in seconds (Please note that the accuracy of this value is depending on your check resolution). | [optional] 
**Totaldown** | Pointer to **int32** | Total downtime in seconds (Please note that the accuracy of this value is depending on your check resolution). | [optional] 
**Totalunknown** | Pointer to **int32** | Total unknown/unmonitored/paused in seconds (Please note that the accuracy of this value is depending on your check resolution. Also note that time before the check was created counts as unknown). | [optional] 

## Methods

### NewSummaryRespAttrsSummaryStatus

`func NewSummaryRespAttrsSummaryStatus() *SummaryRespAttrsSummaryStatus`

NewSummaryRespAttrsSummaryStatus instantiates a new SummaryRespAttrsSummaryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryRespAttrsSummaryStatusWithDefaults

`func NewSummaryRespAttrsSummaryStatusWithDefaults() *SummaryRespAttrsSummaryStatus`

NewSummaryRespAttrsSummaryStatusWithDefaults instantiates a new SummaryRespAttrsSummaryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalup

`func (o *SummaryRespAttrsSummaryStatus) GetTotalup() int32`

GetTotalup returns the Totalup field if non-nil, zero value otherwise.

### GetTotalupOk

`func (o *SummaryRespAttrsSummaryStatus) GetTotalupOk() (*int32, bool)`

GetTotalupOk returns a tuple with the Totalup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalup

`func (o *SummaryRespAttrsSummaryStatus) SetTotalup(v int32)`

SetTotalup sets Totalup field to given value.

### HasTotalup

`func (o *SummaryRespAttrsSummaryStatus) HasTotalup() bool`

HasTotalup returns a boolean if a field has been set.

### GetTotaldown

`func (o *SummaryRespAttrsSummaryStatus) GetTotaldown() int32`

GetTotaldown returns the Totaldown field if non-nil, zero value otherwise.

### GetTotaldownOk

`func (o *SummaryRespAttrsSummaryStatus) GetTotaldownOk() (*int32, bool)`

GetTotaldownOk returns a tuple with the Totaldown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotaldown

`func (o *SummaryRespAttrsSummaryStatus) SetTotaldown(v int32)`

SetTotaldown sets Totaldown field to given value.

### HasTotaldown

`func (o *SummaryRespAttrsSummaryStatus) HasTotaldown() bool`

HasTotaldown returns a boolean if a field has been set.

### GetTotalunknown

`func (o *SummaryRespAttrsSummaryStatus) GetTotalunknown() int32`

GetTotalunknown returns the Totalunknown field if non-nil, zero value otherwise.

### GetTotalunknownOk

`func (o *SummaryRespAttrsSummaryStatus) GetTotalunknownOk() (*int32, bool)`

GetTotalunknownOk returns a tuple with the Totalunknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalunknown

`func (o *SummaryRespAttrsSummaryStatus) SetTotalunknown(v int32)`

SetTotalunknown sets Totalunknown field to given value.

### HasTotalunknown

`func (o *SummaryRespAttrsSummaryStatus) HasTotalunknown() bool`

HasTotalunknown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


