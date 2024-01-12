# Members

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Contact identifier | [optional] 
**Name** | Pointer to **string** | The team member’s name | [optional] 
**Type** | Pointer to **string** | Type defines whether the member is a user (login user) or a contact only | [optional] 

## Methods

### NewMembers

`func NewMembers() *Members`

NewMembers instantiates a new Members object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembersWithDefaults

`func NewMembersWithDefaults() *Members`

NewMembersWithDefaults instantiates a new Members object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Members) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Members) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Members) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Members) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Members) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Members) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Members) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Members) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Members) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Members) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Members) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Members) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


