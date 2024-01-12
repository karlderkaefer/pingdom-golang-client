# AlertingTeams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Team identifier | [optional] 
**Name** | Pointer to **string** | Team name | [optional] 
**Members** | Pointer to [**[]Members**](Members.md) |  | [optional] 

## Methods

### NewAlertingTeams

`func NewAlertingTeams() *AlertingTeams`

NewAlertingTeams instantiates a new AlertingTeams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertingTeamsWithDefaults

`func NewAlertingTeamsWithDefaults() *AlertingTeams`

NewAlertingTeamsWithDefaults instantiates a new AlertingTeams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertingTeams) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertingTeams) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertingTeams) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertingTeams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AlertingTeams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertingTeams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertingTeams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertingTeams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMembers

`func (o *AlertingTeams) GetMembers() []Members`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AlertingTeams) GetMembersOk() (*[]Members, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AlertingTeams) SetMembers(v []Members)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *AlertingTeams) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


