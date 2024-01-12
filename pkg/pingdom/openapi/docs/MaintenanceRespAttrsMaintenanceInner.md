# MaintenanceRespAttrsMaintenanceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Maintenance window identifier | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**From** | Pointer to **int32** | Initial maintenance window start. Format UNIX time. | [optional] 
**To** | Pointer to **int32** | Initial maintenance window end. Format UNIX time. | [optional] 
**Recurrencetype** | Pointer to **string** | Type of recurrence. | [optional] 
**Repeatevery** | Pointer to **int32** | Repeat every n-th day/week/month | [optional] 
**Effectiveto** | Pointer to **int32** | Recurrence end. Format UNIX time. | [optional] 
**Checks** | Pointer to [**MaintenanceRespAttrsMaintenanceInnerChecks**](MaintenanceRespAttrsMaintenanceInnerChecks.md) |  | [optional] 

## Methods

### NewMaintenanceRespAttrsMaintenanceInner

`func NewMaintenanceRespAttrsMaintenanceInner() *MaintenanceRespAttrsMaintenanceInner`

NewMaintenanceRespAttrsMaintenanceInner instantiates a new MaintenanceRespAttrsMaintenanceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceRespAttrsMaintenanceInnerWithDefaults

`func NewMaintenanceRespAttrsMaintenanceInnerWithDefaults() *MaintenanceRespAttrsMaintenanceInner`

NewMaintenanceRespAttrsMaintenanceInnerWithDefaults instantiates a new MaintenanceRespAttrsMaintenanceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaintenanceRespAttrsMaintenanceInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceRespAttrsMaintenanceInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceRespAttrsMaintenanceInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MaintenanceRespAttrsMaintenanceInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MaintenanceRespAttrsMaintenanceInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaintenanceRespAttrsMaintenanceInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaintenanceRespAttrsMaintenanceInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MaintenanceRespAttrsMaintenanceInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrom

`func (o *MaintenanceRespAttrsMaintenanceInner) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MaintenanceRespAttrsMaintenanceInner) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MaintenanceRespAttrsMaintenanceInner) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MaintenanceRespAttrsMaintenanceInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MaintenanceRespAttrsMaintenanceInner) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MaintenanceRespAttrsMaintenanceInner) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MaintenanceRespAttrsMaintenanceInner) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *MaintenanceRespAttrsMaintenanceInner) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetRecurrencetype

`func (o *MaintenanceRespAttrsMaintenanceInner) GetRecurrencetype() string`

GetRecurrencetype returns the Recurrencetype field if non-nil, zero value otherwise.

### GetRecurrencetypeOk

`func (o *MaintenanceRespAttrsMaintenanceInner) GetRecurrencetypeOk() (*string, bool)`

GetRecurrencetypeOk returns a tuple with the Recurrencetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrencetype

`func (o *MaintenanceRespAttrsMaintenanceInner) SetRecurrencetype(v string)`

SetRecurrencetype sets Recurrencetype field to given value.

### HasRecurrencetype

`func (o *MaintenanceRespAttrsMaintenanceInner) HasRecurrencetype() bool`

HasRecurrencetype returns a boolean if a field has been set.

### GetRepeatevery

`func (o *MaintenanceRespAttrsMaintenanceInner) GetRepeatevery() int32`

GetRepeatevery returns the Repeatevery field if non-nil, zero value otherwise.

### GetRepeateveryOk

`func (o *MaintenanceRespAttrsMaintenanceInner) GetRepeateveryOk() (*int32, bool)`

GetRepeateveryOk returns a tuple with the Repeatevery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatevery

`func (o *MaintenanceRespAttrsMaintenanceInner) SetRepeatevery(v int32)`

SetRepeatevery sets Repeatevery field to given value.

### HasRepeatevery

`func (o *MaintenanceRespAttrsMaintenanceInner) HasRepeatevery() bool`

HasRepeatevery returns a boolean if a field has been set.

### GetEffectiveto

`func (o *MaintenanceRespAttrsMaintenanceInner) GetEffectiveto() int32`

GetEffectiveto returns the Effectiveto field if non-nil, zero value otherwise.

### GetEffectivetoOk

`func (o *MaintenanceRespAttrsMaintenanceInner) GetEffectivetoOk() (*int32, bool)`

GetEffectivetoOk returns a tuple with the Effectiveto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveto

`func (o *MaintenanceRespAttrsMaintenanceInner) SetEffectiveto(v int32)`

SetEffectiveto sets Effectiveto field to given value.

### HasEffectiveto

`func (o *MaintenanceRespAttrsMaintenanceInner) HasEffectiveto() bool`

HasEffectiveto returns a boolean if a field has been set.

### GetChecks

`func (o *MaintenanceRespAttrsMaintenanceInner) GetChecks() MaintenanceRespAttrsMaintenanceInnerChecks`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *MaintenanceRespAttrsMaintenanceInner) GetChecksOk() (*MaintenanceRespAttrsMaintenanceInnerChecks, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *MaintenanceRespAttrsMaintenanceInner) SetChecks(v MaintenanceRespAttrsMaintenanceInnerChecks)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *MaintenanceRespAttrsMaintenanceInner) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


