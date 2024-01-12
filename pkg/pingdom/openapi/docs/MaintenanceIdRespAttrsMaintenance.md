# MaintenanceIdRespAttrsMaintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | Maintenance window identifier | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**From** | Pointer to **float32** | Initial maintenance window start. Format UNIX time. | [optional] 
**To** | Pointer to **float32** | Initial maintenance window end. Format UNIX time. | [optional] 
**Recurrencetype** | Pointer to **string** | Type of recurrence. | [optional] 
**Repeatevery** | Pointer to **float32** | Repeat every n-th day/week/month | [optional] 
**Effectiveto** | Pointer to **float32** | Recurrence end. Format UNIX time. | [optional] 
**Checks** | Pointer to [**MaintenanceIdRespAttrsMaintenanceChecks**](MaintenanceIdRespAttrsMaintenanceChecks.md) |  | [optional] 

## Methods

### NewMaintenanceIdRespAttrsMaintenance

`func NewMaintenanceIdRespAttrsMaintenance() *MaintenanceIdRespAttrsMaintenance`

NewMaintenanceIdRespAttrsMaintenance instantiates a new MaintenanceIdRespAttrsMaintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceIdRespAttrsMaintenanceWithDefaults

`func NewMaintenanceIdRespAttrsMaintenanceWithDefaults() *MaintenanceIdRespAttrsMaintenance`

NewMaintenanceIdRespAttrsMaintenanceWithDefaults instantiates a new MaintenanceIdRespAttrsMaintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaintenanceIdRespAttrsMaintenance) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaintenanceIdRespAttrsMaintenance) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaintenanceIdRespAttrsMaintenance) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *MaintenanceIdRespAttrsMaintenance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *MaintenanceIdRespAttrsMaintenance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MaintenanceIdRespAttrsMaintenance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MaintenanceIdRespAttrsMaintenance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MaintenanceIdRespAttrsMaintenance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrom

`func (o *MaintenanceIdRespAttrsMaintenance) GetFrom() float32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MaintenanceIdRespAttrsMaintenance) GetFromOk() (*float32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MaintenanceIdRespAttrsMaintenance) SetFrom(v float32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MaintenanceIdRespAttrsMaintenance) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MaintenanceIdRespAttrsMaintenance) GetTo() float32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MaintenanceIdRespAttrsMaintenance) GetToOk() (*float32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MaintenanceIdRespAttrsMaintenance) SetTo(v float32)`

SetTo sets To field to given value.

### HasTo

`func (o *MaintenanceIdRespAttrsMaintenance) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetRecurrencetype

`func (o *MaintenanceIdRespAttrsMaintenance) GetRecurrencetype() string`

GetRecurrencetype returns the Recurrencetype field if non-nil, zero value otherwise.

### GetRecurrencetypeOk

`func (o *MaintenanceIdRespAttrsMaintenance) GetRecurrencetypeOk() (*string, bool)`

GetRecurrencetypeOk returns a tuple with the Recurrencetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrencetype

`func (o *MaintenanceIdRespAttrsMaintenance) SetRecurrencetype(v string)`

SetRecurrencetype sets Recurrencetype field to given value.

### HasRecurrencetype

`func (o *MaintenanceIdRespAttrsMaintenance) HasRecurrencetype() bool`

HasRecurrencetype returns a boolean if a field has been set.

### GetRepeatevery

`func (o *MaintenanceIdRespAttrsMaintenance) GetRepeatevery() float32`

GetRepeatevery returns the Repeatevery field if non-nil, zero value otherwise.

### GetRepeateveryOk

`func (o *MaintenanceIdRespAttrsMaintenance) GetRepeateveryOk() (*float32, bool)`

GetRepeateveryOk returns a tuple with the Repeatevery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatevery

`func (o *MaintenanceIdRespAttrsMaintenance) SetRepeatevery(v float32)`

SetRepeatevery sets Repeatevery field to given value.

### HasRepeatevery

`func (o *MaintenanceIdRespAttrsMaintenance) HasRepeatevery() bool`

HasRepeatevery returns a boolean if a field has been set.

### GetEffectiveto

`func (o *MaintenanceIdRespAttrsMaintenance) GetEffectiveto() float32`

GetEffectiveto returns the Effectiveto field if non-nil, zero value otherwise.

### GetEffectivetoOk

`func (o *MaintenanceIdRespAttrsMaintenance) GetEffectivetoOk() (*float32, bool)`

GetEffectivetoOk returns a tuple with the Effectiveto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveto

`func (o *MaintenanceIdRespAttrsMaintenance) SetEffectiveto(v float32)`

SetEffectiveto sets Effectiveto field to given value.

### HasEffectiveto

`func (o *MaintenanceIdRespAttrsMaintenance) HasEffectiveto() bool`

HasEffectiveto returns a boolean if a field has been set.

### GetChecks

`func (o *MaintenanceIdRespAttrsMaintenance) GetChecks() MaintenanceIdRespAttrsMaintenanceChecks`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *MaintenanceIdRespAttrsMaintenance) GetChecksOk() (*MaintenanceIdRespAttrsMaintenanceChecks, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *MaintenanceIdRespAttrsMaintenance) SetChecks(v MaintenanceIdRespAttrsMaintenanceChecks)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *MaintenanceIdRespAttrsMaintenance) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


