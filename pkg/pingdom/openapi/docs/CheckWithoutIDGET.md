# CheckWithoutIDGET

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Check status - active or inactive | [optional] 
**ContactIds** | Pointer to **[]int64** | Contacts to alert | [optional] 
**CreatedAt** | Pointer to **int64** | Timestamp when the check was created | [optional] 
**ModifiedAt** | Pointer to **int64** | Timestamp when the check was modified | [optional] 
**LastDowntimeStart** | Pointer to **int64** | Timestamp when the last downtime started. This field is optional | [optional] 
**LastDowntimeEnd** | Pointer to **int64** | Timestamp when the last downtime ended. This field is optional | [optional] 
**CustomMessage** | Pointer to **string** | Custom message that is part of the email and webhook alerts | [optional] 
**Interval** | Pointer to **int64** | TMS test intervals in minutes. Allowed intervals: 5,10,20,60,720,1440. The interval you&#39;re allowed to set may vary depending on your current plan. | [optional] 
**Name** | Pointer to **string** | Name of the check | [optional] 
**Region** | Pointer to **string** | Name of the region where the check is executed. Supported regions: us-east, us-west, eu, au | [optional] 
**SendNotificationWhenDown** | Pointer to **int64** | Send notification when down X times | [optional] 
**SeverityLevel** | Pointer to **string** | Check importance- how important are the alerts when the check fails. Allowed values: low, high | [optional] 
**Status** | Pointer to **string** | Whether the check is passing or failing at the moment (successful, failing, unknown) | [optional] 
**Steps** | Pointer to [**[]Step**](Step.md) | steps to be executed as part of the check | [optional] 
**TeamIds** | Pointer to **[]int64** | Teams to alert | [optional] 
**IntegrationIds** | Pointer to **[]int64** | Integration identifiers. | [optional] 
**Metadata** | Pointer to [**MetadataGET**](MetadataGET.md) |  | [optional] 
**Tags** | Pointer to **[]string** | List of tags for a check. The tag name may contain the characters &#39;A-Z&#39;, &#39;a-z&#39;, &#39;0-9&#39;, &#39;_&#39; and &#39;-&#39;. The maximum length of a tag is 64 characters. | [optional] 
**Type** | Pointer to **string** | Type of transaction check: \&quot;script\&quot; for regular TMS checks and \&quot;recording\&quot; for checks made using the Transaction Recorder | [optional] 

## Methods

### NewCheckWithoutIDGET

`func NewCheckWithoutIDGET() *CheckWithoutIDGET`

NewCheckWithoutIDGET instantiates a new CheckWithoutIDGET object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithoutIDGETWithDefaults

`func NewCheckWithoutIDGETWithDefaults() *CheckWithoutIDGET`

NewCheckWithoutIDGETWithDefaults instantiates a new CheckWithoutIDGET object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *CheckWithoutIDGET) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CheckWithoutIDGET) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CheckWithoutIDGET) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CheckWithoutIDGET) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetContactIds

`func (o *CheckWithoutIDGET) GetContactIds() []int64`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *CheckWithoutIDGET) GetContactIdsOk() (*[]int64, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *CheckWithoutIDGET) SetContactIds(v []int64)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *CheckWithoutIDGET) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CheckWithoutIDGET) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckWithoutIDGET) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckWithoutIDGET) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CheckWithoutIDGET) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *CheckWithoutIDGET) GetModifiedAt() int64`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *CheckWithoutIDGET) GetModifiedAtOk() (*int64, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *CheckWithoutIDGET) SetModifiedAt(v int64)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *CheckWithoutIDGET) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetLastDowntimeStart

`func (o *CheckWithoutIDGET) GetLastDowntimeStart() int64`

GetLastDowntimeStart returns the LastDowntimeStart field if non-nil, zero value otherwise.

### GetLastDowntimeStartOk

`func (o *CheckWithoutIDGET) GetLastDowntimeStartOk() (*int64, bool)`

GetLastDowntimeStartOk returns a tuple with the LastDowntimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDowntimeStart

`func (o *CheckWithoutIDGET) SetLastDowntimeStart(v int64)`

SetLastDowntimeStart sets LastDowntimeStart field to given value.

### HasLastDowntimeStart

`func (o *CheckWithoutIDGET) HasLastDowntimeStart() bool`

HasLastDowntimeStart returns a boolean if a field has been set.

### GetLastDowntimeEnd

`func (o *CheckWithoutIDGET) GetLastDowntimeEnd() int64`

GetLastDowntimeEnd returns the LastDowntimeEnd field if non-nil, zero value otherwise.

### GetLastDowntimeEndOk

`func (o *CheckWithoutIDGET) GetLastDowntimeEndOk() (*int64, bool)`

GetLastDowntimeEndOk returns a tuple with the LastDowntimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDowntimeEnd

`func (o *CheckWithoutIDGET) SetLastDowntimeEnd(v int64)`

SetLastDowntimeEnd sets LastDowntimeEnd field to given value.

### HasLastDowntimeEnd

`func (o *CheckWithoutIDGET) HasLastDowntimeEnd() bool`

HasLastDowntimeEnd returns a boolean if a field has been set.

### GetCustomMessage

`func (o *CheckWithoutIDGET) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *CheckWithoutIDGET) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *CheckWithoutIDGET) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *CheckWithoutIDGET) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.

### GetInterval

`func (o *CheckWithoutIDGET) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CheckWithoutIDGET) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CheckWithoutIDGET) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *CheckWithoutIDGET) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *CheckWithoutIDGET) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckWithoutIDGET) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckWithoutIDGET) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CheckWithoutIDGET) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegion

`func (o *CheckWithoutIDGET) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CheckWithoutIDGET) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CheckWithoutIDGET) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CheckWithoutIDGET) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSendNotificationWhenDown

`func (o *CheckWithoutIDGET) GetSendNotificationWhenDown() int64`

GetSendNotificationWhenDown returns the SendNotificationWhenDown field if non-nil, zero value otherwise.

### GetSendNotificationWhenDownOk

`func (o *CheckWithoutIDGET) GetSendNotificationWhenDownOk() (*int64, bool)`

GetSendNotificationWhenDownOk returns a tuple with the SendNotificationWhenDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendNotificationWhenDown

`func (o *CheckWithoutIDGET) SetSendNotificationWhenDown(v int64)`

SetSendNotificationWhenDown sets SendNotificationWhenDown field to given value.

### HasSendNotificationWhenDown

`func (o *CheckWithoutIDGET) HasSendNotificationWhenDown() bool`

HasSendNotificationWhenDown returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *CheckWithoutIDGET) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *CheckWithoutIDGET) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *CheckWithoutIDGET) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *CheckWithoutIDGET) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetStatus

`func (o *CheckWithoutIDGET) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckWithoutIDGET) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckWithoutIDGET) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckWithoutIDGET) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSteps

`func (o *CheckWithoutIDGET) GetSteps() []Step`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CheckWithoutIDGET) GetStepsOk() (*[]Step, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CheckWithoutIDGET) SetSteps(v []Step)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CheckWithoutIDGET) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTeamIds

`func (o *CheckWithoutIDGET) GetTeamIds() []int64`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *CheckWithoutIDGET) GetTeamIdsOk() (*[]int64, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *CheckWithoutIDGET) SetTeamIds(v []int64)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *CheckWithoutIDGET) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.

### GetIntegrationIds

`func (o *CheckWithoutIDGET) GetIntegrationIds() []int64`

GetIntegrationIds returns the IntegrationIds field if non-nil, zero value otherwise.

### GetIntegrationIdsOk

`func (o *CheckWithoutIDGET) GetIntegrationIdsOk() (*[]int64, bool)`

GetIntegrationIdsOk returns a tuple with the IntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationIds

`func (o *CheckWithoutIDGET) SetIntegrationIds(v []int64)`

SetIntegrationIds sets IntegrationIds field to given value.

### HasIntegrationIds

`func (o *CheckWithoutIDGET) HasIntegrationIds() bool`

HasIntegrationIds returns a boolean if a field has been set.

### GetMetadata

`func (o *CheckWithoutIDGET) GetMetadata() MetadataGET`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CheckWithoutIDGET) GetMetadataOk() (*MetadataGET, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CheckWithoutIDGET) SetMetadata(v MetadataGET)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CheckWithoutIDGET) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTags

`func (o *CheckWithoutIDGET) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckWithoutIDGET) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckWithoutIDGET) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckWithoutIDGET) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *CheckWithoutIDGET) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckWithoutIDGET) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckWithoutIDGET) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CheckWithoutIDGET) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


