# MaintenanceRespAttrs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Maintenance** | Pointer to [**[]MaintenanceRespAttrsMaintenanceInner**](MaintenanceRespAttrsMaintenanceInner.md) | A list of maintenance windows | [optional] 

## Methods

### NewMaintenanceRespAttrs

`func NewMaintenanceRespAttrs() *MaintenanceRespAttrs`

NewMaintenanceRespAttrs instantiates a new MaintenanceRespAttrs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceRespAttrsWithDefaults

`func NewMaintenanceRespAttrsWithDefaults() *MaintenanceRespAttrs`

NewMaintenanceRespAttrsWithDefaults instantiates a new MaintenanceRespAttrs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaintenance

`func (o *MaintenanceRespAttrs) GetMaintenance() []MaintenanceRespAttrsMaintenanceInner`

GetMaintenance returns the Maintenance field if non-nil, zero value otherwise.

### GetMaintenanceOk

`func (o *MaintenanceRespAttrs) GetMaintenanceOk() (*[]MaintenanceRespAttrsMaintenanceInner, bool)`

GetMaintenanceOk returns a tuple with the Maintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenance

`func (o *MaintenanceRespAttrs) SetMaintenance(v []MaintenanceRespAttrsMaintenanceInner)`

SetMaintenance sets Maintenance field to given value.

### HasMaintenance

`func (o *MaintenanceRespAttrs) HasMaintenance() bool`

HasMaintenance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


