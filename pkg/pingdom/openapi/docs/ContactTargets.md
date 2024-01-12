# ContactTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Contact ID | [optional] 
**Name** | Pointer to **string** | Contact name | [optional] 
**Paused** | Pointer to **bool** | Describes whether alerts are paused for this contact | [optional] 
**Type** | Pointer to **string** | Type defines whether this is a user (login user) or a contact only | [optional] 
**Owner** | Pointer to **bool** | Indicates whether the contact is the owner of the organization | [optional] 
**NotificationTargets** | Pointer to [**ContactTargetsNotificationTargets**](ContactTargetsNotificationTargets.md) |  | [optional] 
**Teams** | Pointer to [**[]ContactTargetsTeamsInner**](ContactTargetsTeamsInner.md) |  | [optional] 

## Methods

### NewContactTargets

`func NewContactTargets() *ContactTargets`

NewContactTargets instantiates a new ContactTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactTargetsWithDefaults

`func NewContactTargetsWithDefaults() *ContactTargets`

NewContactTargetsWithDefaults instantiates a new ContactTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactTargets) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactTargets) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactTargets) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactTargets) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ContactTargets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactTargets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactTargets) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactTargets) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaused

`func (o *ContactTargets) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ContactTargets) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ContactTargets) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ContactTargets) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetType

`func (o *ContactTargets) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactTargets) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactTargets) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContactTargets) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwner

`func (o *ContactTargets) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ContactTargets) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ContactTargets) SetOwner(v bool)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ContactTargets) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetNotificationTargets

`func (o *ContactTargets) GetNotificationTargets() ContactTargetsNotificationTargets`

GetNotificationTargets returns the NotificationTargets field if non-nil, zero value otherwise.

### GetNotificationTargetsOk

`func (o *ContactTargets) GetNotificationTargetsOk() (*ContactTargetsNotificationTargets, bool)`

GetNotificationTargetsOk returns a tuple with the NotificationTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTargets

`func (o *ContactTargets) SetNotificationTargets(v ContactTargetsNotificationTargets)`

SetNotificationTargets sets NotificationTargets field to given value.

### HasNotificationTargets

`func (o *ContactTargets) HasNotificationTargets() bool`

HasNotificationTargets returns a boolean if a field has been set.

### GetTeams

`func (o *ContactTargets) GetTeams() []ContactTargetsTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ContactTargets) GetTeamsOk() (*[]ContactTargetsTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ContactTargets) SetTeams(v []ContactTargetsTeamsInner)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ContactTargets) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


