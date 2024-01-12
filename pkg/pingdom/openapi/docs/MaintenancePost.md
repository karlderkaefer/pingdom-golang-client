# MaintenancePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | Description | 
**From** | **int32** | Initial maintenance window start. Format UNIX time. (Only future allowed. Use 1 for the current timestamp.) | 
**To** | **int32** | Initial maintenance window end. Format UNIX time. (Only future allowed. Use 1 for the current timestamp.) | 
**Recurrencetype** | Pointer to **string** | Type of recurrence. | [optional] [default to "none"]
**Repeatevery** | Pointer to **int32** | Repeat every n-th day/week/month | [optional] [default to 0]
**Effectiveto** | Pointer to **int32** | Recurrence end. Format UNIX time. Default: equal to &#x60;to&#x60;. (Only future allowed. Use 1 for the current timestamp.) | [optional] 
**Uptimeids** | Pointer to **[]int32** | Identifiers of uptime checks to assign to the maintenance window - Comma separated Integers | [optional] 
**Tmsids** | Pointer to **[]int32** | Identifiers of transaction checks to assign to the maintenance window - Comma separated Integers | [optional] 

## Methods

### NewMaintenancePost

`func NewMaintenancePost(description string, from int32, to int32, ) *MaintenancePost`

NewMaintenancePost instantiates a new MaintenancePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenancePostWithDefaults

`func NewMaintenancePostWithDefaults() *MaintenancePost`

NewMaintenancePostWithDefaults instantiates a new MaintenancePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MaintenancePost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaintenancePost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaintenancePost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFrom

`func (o *MaintenancePost) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MaintenancePost) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MaintenancePost) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetTo

`func (o *MaintenancePost) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MaintenancePost) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MaintenancePost) SetTo(v int32)`

SetTo sets To field to given value.


### GetRecurrencetype

`func (o *MaintenancePost) GetRecurrencetype() string`

GetRecurrencetype returns the Recurrencetype field if non-nil, zero value otherwise.

### GetRecurrencetypeOk

`func (o *MaintenancePost) GetRecurrencetypeOk() (*string, bool)`

GetRecurrencetypeOk returns a tuple with the Recurrencetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrencetype

`func (o *MaintenancePost) SetRecurrencetype(v string)`

SetRecurrencetype sets Recurrencetype field to given value.

### HasRecurrencetype

`func (o *MaintenancePost) HasRecurrencetype() bool`

HasRecurrencetype returns a boolean if a field has been set.

### GetRepeatevery

`func (o *MaintenancePost) GetRepeatevery() int32`

GetRepeatevery returns the Repeatevery field if non-nil, zero value otherwise.

### GetRepeateveryOk

`func (o *MaintenancePost) GetRepeateveryOk() (*int32, bool)`

GetRepeateveryOk returns a tuple with the Repeatevery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatevery

`func (o *MaintenancePost) SetRepeatevery(v int32)`

SetRepeatevery sets Repeatevery field to given value.

### HasRepeatevery

`func (o *MaintenancePost) HasRepeatevery() bool`

HasRepeatevery returns a boolean if a field has been set.

### GetEffectiveto

`func (o *MaintenancePost) GetEffectiveto() int32`

GetEffectiveto returns the Effectiveto field if non-nil, zero value otherwise.

### GetEffectivetoOk

`func (o *MaintenancePost) GetEffectivetoOk() (*int32, bool)`

GetEffectivetoOk returns a tuple with the Effectiveto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveto

`func (o *MaintenancePost) SetEffectiveto(v int32)`

SetEffectiveto sets Effectiveto field to given value.

### HasEffectiveto

`func (o *MaintenancePost) HasEffectiveto() bool`

HasEffectiveto returns a boolean if a field has been set.

### GetUptimeids

`func (o *MaintenancePost) GetUptimeids() []int32`

GetUptimeids returns the Uptimeids field if non-nil, zero value otherwise.

### GetUptimeidsOk

`func (o *MaintenancePost) GetUptimeidsOk() (*[]int32, bool)`

GetUptimeidsOk returns a tuple with the Uptimeids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptimeids

`func (o *MaintenancePost) SetUptimeids(v []int32)`

SetUptimeids sets Uptimeids field to given value.

### HasUptimeids

`func (o *MaintenancePost) HasUptimeids() bool`

HasUptimeids returns a boolean if a field has been set.

### GetTmsids

`func (o *MaintenancePost) GetTmsids() []int32`

GetTmsids returns the Tmsids field if non-nil, zero value otherwise.

### GetTmsidsOk

`func (o *MaintenancePost) GetTmsidsOk() (*[]int32, bool)`

GetTmsidsOk returns a tuple with the Tmsids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmsids

`func (o *MaintenancePost) SetTmsids(v []int32)`

SetTmsids sets Tmsids field to given value.

### HasTmsids

`func (o *MaintenancePost) HasTmsids() bool`

HasTmsids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


