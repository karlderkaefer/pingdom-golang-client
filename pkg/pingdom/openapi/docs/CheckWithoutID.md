# CheckWithoutID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Check status: active or inactive | [optional] [default to true]
**ContactIds** | Pointer to **[]int64** | Contacts to alert | [optional] 
**CustomMessage** | Pointer to **string** | Custom message that is part of the email and webhook alerts | [optional] 
**Interval** | Pointer to **int64** | TMS test intervals in minutes. Allowed intervals: 5,10,20,60,720,1440. The interval you&#39;re allowed to set may vary depending on your current plan. | [optional] [default to 10]
**Name** | **string** | Name of the check | 
**Region** | Pointer to **string** | Name of the region where the check is executed. Supported regions: us-east, us-west, eu, au | [optional] [default to "us-east"]
**SendNotificationWhenDown** | Pointer to **int64** | Send notification when down X times | [optional] [default to 1]
**SeverityLevel** | Pointer to **string** | Check importance- how important are the alerts when the check fails. Allowed values: low, high | [optional] [default to "high"]
**Steps** | [**[]Step**](Step.md) | steps to be executed as part of the check | 
**TeamIds** | Pointer to **[]int64** | Teams to alert | [optional] 
**IntegrationIds** | Pointer to **[]int64** | Integration identifiers. | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Tags** | Pointer to **[]string** | List of tags for a check. The tag name may contain the characters &#39;A-Z&#39;, &#39;a-z&#39;, &#39;0-9&#39;, &#39;_&#39; and &#39;-&#39;. The maximum length of a tag is 64 characters. | [optional] 

## Methods

### NewCheckWithoutID

`func NewCheckWithoutID(name string, steps []Step, ) *CheckWithoutID`

NewCheckWithoutID instantiates a new CheckWithoutID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithoutIDWithDefaults

`func NewCheckWithoutIDWithDefaults() *CheckWithoutID`

NewCheckWithoutIDWithDefaults instantiates a new CheckWithoutID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CheckWithoutID) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CheckWithoutID) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CheckWithoutID) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CheckWithoutID) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetContactIds

`func (o *CheckWithoutID) GetContactIds() []int64`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *CheckWithoutID) GetContactIdsOk() (*[]int64, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *CheckWithoutID) SetContactIds(v []int64)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *CheckWithoutID) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### GetCustomMessage

`func (o *CheckWithoutID) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *CheckWithoutID) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *CheckWithoutID) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *CheckWithoutID) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetInterval

`func (o *CheckWithoutID) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CheckWithoutID) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CheckWithoutID) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CheckWithoutID) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *CheckWithoutID) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckWithoutID) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckWithoutID) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *CheckWithoutID) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CheckWithoutID) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CheckWithoutID) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CheckWithoutID) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSendNotificationWhenDown

`func (o *CheckWithoutID) GetSendNotificationWhenDown() int64`

GetSendNotificationWhenDown returns the SendNotificationWhenDown field if non-nil, zero value otherwise.

### GetSendNotificationWhenDownOk

`func (o *CheckWithoutID) GetSendNotificationWhenDownOk() (*int64, bool)`

GetSendNotificationWhenDownOk returns a tuple with the SendNotificationWhenDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotificationWhenDown

`func (o *CheckWithoutID) SetSendNotificationWhenDown(v int64)`

SetSendNotificationWhenDown sets SendNotificationWhenDown field to given value.

### HasSendNotificationWhenDown

`func (o *CheckWithoutID) HasSendNotificationWhenDown() bool`

HasSendNotificationWhenDown returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *CheckWithoutID) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *CheckWithoutID) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *CheckWithoutID) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *CheckWithoutID) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetSteps

`func (o *CheckWithoutID) GetSteps() []Step`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CheckWithoutID) GetStepsOk() (*[]Step, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CheckWithoutID) SetSteps(v []Step)`

SetSteps sets Steps field to given value.


### GetTeamIds

`func (o *CheckWithoutID) GetTeamIds() []int64`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *CheckWithoutID) GetTeamIdsOk() (*[]int64, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *CheckWithoutID) SetTeamIds(v []int64)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *CheckWithoutID) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetIntegrationIds

`func (o *CheckWithoutID) GetIntegrationIds() []int64`

GetIntegrationIds returns the IntegrationIds field if non-nil, zero value otherwise.

### GetIntegrationIdsOk

`func (o *CheckWithoutID) GetIntegrationIdsOk() (*[]int64, bool)`

GetIntegrationIdsOk returns a tuple with the IntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIds

`func (o *CheckWithoutID) SetIntegrationIds(v []int64)`

SetIntegrationIds sets IntegrationIds field to given value.

### HasIntegrationIds

`func (o *CheckWithoutID) HasIntegrationIds() bool`

HasIntegrationIds returns a boolean if a field has been set.

### GetMetadata

`func (o *CheckWithoutID) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckWithoutID) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckWithoutID) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckWithoutID) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTags

`func (o *CheckWithoutID) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckWithoutID) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckWithoutID) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckWithoutID) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


