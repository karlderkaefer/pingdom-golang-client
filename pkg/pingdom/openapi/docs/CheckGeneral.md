# CheckGeneral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Check status: active or inactive | [optional] 
**CreatedAt** | Pointer to **int64** | Timestamp when the check was created | [optional] 
**Id** | Pointer to **int64** | Id of the check | [optional] 
**Interval** | Pointer to **int64** | TMS test intervals in minutes. Allowed intervals: 5,10,20,60,720,1440. The interval you&#39;re allowed to set may vary depending on your current plan. | [optional] 
**ModifiedAt** | Pointer to **int64** | Timestamp when the check was modified | [optional] 
**LastDowntimeStart** | Pointer to **int64** | Timestamp when the last downtime started. This field is optional | [optional] 
**LastDowntimeEnd** | Pointer to **int64** | Timestamp when the last downtime ended. This field is optional | [optional] 
**Name** | Pointer to **string** | Name of the check | [optional] 
**Region** | Pointer to **string** | Name of the region where the check is executed. Supported regions: us-east, us-west, eu, au | [optional] 
**Status** | Pointer to **string** | Whether the check is passing or failing at the moment (successful, failing, unknown) | [optional] 
**Tags** | Pointer to **[]string** | List of tags for a check. The tag name may contain the characters &#39;A-Z&#39;, &#39;a-z&#39;, &#39;0-9&#39;, &#39;_&#39; and &#39;-&#39;. The maximum length of a tag is 64 characters. | [optional] 
**Type** | Pointer to **string** | Type of transaction check: \&quot;script\&quot; for regular TMS checks and \&quot;recording\&quot; for checks made using the Transaction Recorder | [optional] 

## Methods

### NewCheckGeneral

`func NewCheckGeneral() *CheckGeneral`

NewCheckGeneral instantiates a new CheckGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckGeneralWithDefaults

`func NewCheckGeneralWithDefaults() *CheckGeneral`

NewCheckGeneralWithDefaults instantiates a new CheckGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CheckGeneral) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CheckGeneral) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CheckGeneral) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CheckGeneral) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CheckGeneral) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckGeneral) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckGeneral) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CheckGeneral) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CheckGeneral) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckGeneral) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckGeneral) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CheckGeneral) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterval

`func (o *CheckGeneral) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CheckGeneral) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CheckGeneral) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CheckGeneral) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetModifiedAt

`func (o *CheckGeneral) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *CheckGeneral) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *CheckGeneral) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *CheckGeneral) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetLastDowntimeStart

`func (o *CheckGeneral) GetLastDowntimeStart() int64`

GetLastDowntimeStart returns the LastDowntimeStart field if non-nil, zero value otherwise.

### GetLastDowntimeStartOk

`func (o *CheckGeneral) GetLastDowntimeStartOk() (*int64, bool)`

GetLastDowntimeStartOk returns a tuple with the LastDowntimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDowntimeStart

`func (o *CheckGeneral) SetLastDowntimeStart(v int64)`

SetLastDowntimeStart sets LastDowntimeStart field to given value.

### HasLastDowntimeStart

`func (o *CheckGeneral) HasLastDowntimeStart() bool`

HasLastDowntimeStart returns a boolean if a field has been set.

### GetLastDowntimeEnd

`func (o *CheckGeneral) GetLastDowntimeEnd() int64`

GetLastDowntimeEnd returns the LastDowntimeEnd field if non-nil, zero value otherwise.

### GetLastDowntimeEndOk

`func (o *CheckGeneral) GetLastDowntimeEndOk() (*int64, bool)`

GetLastDowntimeEndOk returns a tuple with the LastDowntimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDowntimeEnd

`func (o *CheckGeneral) SetLastDowntimeEnd(v int64)`

SetLastDowntimeEnd sets LastDowntimeEnd field to given value.

### HasLastDowntimeEnd

`func (o *CheckGeneral) HasLastDowntimeEnd() bool`

HasLastDowntimeEnd returns a boolean if a field has been set.

### GetName

`func (o *CheckGeneral) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckGeneral) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckGeneral) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckGeneral) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CheckGeneral) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CheckGeneral) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CheckGeneral) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CheckGeneral) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *CheckGeneral) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckGeneral) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckGeneral) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckGeneral) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *CheckGeneral) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckGeneral) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckGeneral) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckGeneral) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *CheckGeneral) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckGeneral) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckGeneral) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckGeneral) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


