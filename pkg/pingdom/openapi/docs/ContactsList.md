# ContactsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contacts** | Pointer to [**[]ContactTargets**](ContactTargets.md) | List of all contacts targets | [optional] 

## Methods

### NewContactsList

`func NewContactsList() *ContactsList`

NewContactsList instantiates a new ContactsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactsListWithDefaults

`func NewContactsListWithDefaults() *ContactsList`

NewContactsListWithDefaults instantiates a new ContactsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContacts

`func (o *ContactsList) GetContacts() []ContactTargets`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ContactsList) GetContactsOk() (*[]ContactTargets, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ContactsList) SetContacts(v []ContactTargets)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ContactsList) HasContacts() bool`

HasContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


