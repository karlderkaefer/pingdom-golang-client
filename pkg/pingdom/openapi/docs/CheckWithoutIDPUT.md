# CheckWithoutIDPUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Check status: active or inactive | [optional] 
**ContactIds** | Pointer to **[]int64** | Contacts to alert | [optional] 
**CustomMessage** | Pointer to **string** | Custom message that is part of the email and webhook alerts | [optional] 
**Interval** | Pointer to **int64** | TMS test intervals in minutes. Allowed intervals: 5,10,20,60,720,1440. The interval you&#39;re allowed to set may vary depending on your current plan. | [optional] 
**Name** | Pointer to **string** | Name of the check | [optional] 
**Region** | Pointer to **string** | Name of the region where the check is executed. Supported regions: us-east, us-west, eu, au | [optional] 
**SendNotificationWhenDown** | Pointer to **int64** | Send notification when down X times | [optional] 
**SeverityLevel** | Pointer to **string** | Check importance- how important are the alerts when the check fails. Allowed values: low, high | [optional] 
**Steps** | Pointer to [**[]Step**](Step.md) | steps to be executed as part of the check | [optional] 
**TeamIds** | Pointer to **[]int64** | Teams to alert | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Tags** | Pointer to **[]string** | List of tags for a check. The tag name may contain the characters &#39;A-Z&#39;, &#39;a-z&#39;, &#39;0-9&#39;, &#39;_&#39; and &#39;-&#39;. The maximum length of a tag is 64 characters. | [optional] 
**IntegrationIds** | Pointer to **[]int32** | Integration identifiers as a list of integers. | [optional] 

## Methods

### NewCheckWithoutIDPUT

`func NewCheckWithoutIDPUT() *CheckWithoutIDPUT`

NewCheckWithoutIDPUT instantiates a new CheckWithoutIDPUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithoutIDPUTWithDefaults

`func NewCheckWithoutIDPUTWithDefaults() *CheckWithoutIDPUT`

NewCheckWithoutIDPUTWithDefaults instantiates a new CheckWithoutIDPUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CheckWithoutIDPUT) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CheckWithoutIDPUT) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CheckWithoutIDPUT) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CheckWithoutIDPUT) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetContactIds

`func (o *CheckWithoutIDPUT) GetContactIds() []int64`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *CheckWithoutIDPUT) GetContactIdsOk() (*[]int64, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *CheckWithoutIDPUT) SetContactIds(v []int64)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *CheckWithoutIDPUT) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### GetCustomMessage

`func (o *CheckWithoutIDPUT) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *CheckWithoutIDPUT) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *CheckWithoutIDPUT) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *CheckWithoutIDPUT) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetInterval

`func (o *CheckWithoutIDPUT) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CheckWithoutIDPUT) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CheckWithoutIDPUT) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CheckWithoutIDPUT) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *CheckWithoutIDPUT) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckWithoutIDPUT) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckWithoutIDPUT) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckWithoutIDPUT) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CheckWithoutIDPUT) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CheckWithoutIDPUT) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CheckWithoutIDPUT) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CheckWithoutIDPUT) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSendNotificationWhenDown

`func (o *CheckWithoutIDPUT) GetSendNotificationWhenDown() int64`

GetSendNotificationWhenDown returns the SendNotificationWhenDown field if non-nil, zero value otherwise.

### GetSendNotificationWhenDownOk

`func (o *CheckWithoutIDPUT) GetSendNotificationWhenDownOk() (*int64, bool)`

GetSendNotificationWhenDownOk returns a tuple with the SendNotificationWhenDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotificationWhenDown

`func (o *CheckWithoutIDPUT) SetSendNotificationWhenDown(v int64)`

SetSendNotificationWhenDown sets SendNotificationWhenDown field to given value.

### HasSendNotificationWhenDown

`func (o *CheckWithoutIDPUT) HasSendNotificationWhenDown() bool`

HasSendNotificationWhenDown returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *CheckWithoutIDPUT) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *CheckWithoutIDPUT) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *CheckWithoutIDPUT) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *CheckWithoutIDPUT) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetSteps

`func (o *CheckWithoutIDPUT) GetSteps() []Step`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CheckWithoutIDPUT) GetStepsOk() (*[]Step, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CheckWithoutIDPUT) SetSteps(v []Step)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CheckWithoutIDPUT) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTeamIds

`func (o *CheckWithoutIDPUT) GetTeamIds() []int64`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *CheckWithoutIDPUT) GetTeamIdsOk() (*[]int64, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *CheckWithoutIDPUT) SetTeamIds(v []int64)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *CheckWithoutIDPUT) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetMetadata

`func (o *CheckWithoutIDPUT) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckWithoutIDPUT) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckWithoutIDPUT) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckWithoutIDPUT) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTags

`func (o *CheckWithoutIDPUT) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckWithoutIDPUT) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckWithoutIDPUT) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckWithoutIDPUT) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIntegrationIds

`func (o *CheckWithoutIDPUT) GetIntegrationIds() []int32`

GetIntegrationIds returns the IntegrationIds field if non-nil, zero value otherwise.

### GetIntegrationIdsOk

`func (o *CheckWithoutIDPUT) GetIntegrationIdsOk() (*[]int32, bool)`

GetIntegrationIdsOk returns a tuple with the IntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIds

`func (o *CheckWithoutIDPUT) SetIntegrationIds(v []int32)`

SetIntegrationIds sets IntegrationIds field to given value.

### HasIntegrationIds

`func (o *CheckWithoutIDPUT) HasIntegrationIds() bool`

HasIntegrationIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


