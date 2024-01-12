# UpdateContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Contact name | 
**Paused** | **bool** | Pause contact | 
**NotificationTargets** | [**CreateContactNotificationTargets**](CreateContactNotificationTargets.md) |  | 

## Methods

### NewUpdateContact

`func NewUpdateContact(name string, paused bool, notificationTargets CreateContactNotificationTargets, ) *UpdateContact`

NewUpdateContact instantiates a new UpdateContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContactWithDefaults

`func NewUpdateContactWithDefaults() *UpdateContact`

NewUpdateContactWithDefaults instantiates a new UpdateContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateContact) SetName(v string)`

SetName sets Name field to given value.


### GetPaused

`func (o *UpdateContact) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *UpdateContact) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *UpdateContact) SetPaused(v bool)`

SetPaused sets Paused field to given value.


### GetNotificationTargets

`func (o *UpdateContact) GetNotificationTargets() CreateContactNotificationTargets`

GetNotificationTargets returns the NotificationTargets field if non-nil, zero value otherwise.

### GetNotificationTargetsOk

`func (o *UpdateContact) GetNotificationTargetsOk() (*CreateContactNotificationTargets, bool)`

GetNotificationTargetsOk returns a tuple with the NotificationTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTargets

`func (o *UpdateContact) SetNotificationTargets(v CreateContactNotificationTargets)`

SetNotificationTargets sets NotificationTargets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


