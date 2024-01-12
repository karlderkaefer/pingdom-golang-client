# ImapAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | Target port | [optional] 
**Stringtoexpect** | Pointer to **string** | String to expect in response | [optional] 

## Methods

### NewImapAttributes

`func NewImapAttributes() *ImapAttributes`

NewImapAttributes instantiates a new ImapAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImapAttributesWithDefaults

`func NewImapAttributesWithDefaults() *ImapAttributes`

NewImapAttributesWithDefaults instantiates a new ImapAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ImapAttributes) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ImapAttributes) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ImapAttributes) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ImapAttributes) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStringtoexpect

`func (o *ImapAttributes) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *ImapAttributes) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *ImapAttributes) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.

### HasStringtoexpect

`func (o *ImapAttributes) HasStringtoexpect() bool`

HasStringtoexpect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


