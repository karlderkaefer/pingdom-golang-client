# MaintenanceOccurrencesIdPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **int32** | Beginning of the maintenance occurrence. Format UNIX time. (Only future allowed. Use 1 for the current timestamp.) | [optional] 
**To** | Pointer to **int32** | End of the maintenance occurrence. Format UNIX time. (Only future allowed. Use 1 for the current timestamp.) | [optional] 

## Methods

### NewMaintenanceOccurrencesIdPut

`func NewMaintenanceOccurrencesIdPut() *MaintenanceOccurrencesIdPut`

NewMaintenanceOccurrencesIdPut instantiates a new MaintenanceOccurrencesIdPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceOccurrencesIdPutWithDefaults

`func NewMaintenanceOccurrencesIdPutWithDefaults() *MaintenanceOccurrencesIdPut`

NewMaintenanceOccurrencesIdPutWithDefaults instantiates a new MaintenanceOccurrencesIdPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *MaintenanceOccurrencesIdPut) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *MaintenanceOccurrencesIdPut) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *MaintenanceOccurrencesIdPut) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *MaintenanceOccurrencesIdPut) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *MaintenanceOccurrencesIdPut) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *MaintenanceOccurrencesIdPut) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *MaintenanceOccurrencesIdPut) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *MaintenanceOccurrencesIdPut) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


