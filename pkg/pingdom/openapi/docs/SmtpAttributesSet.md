# SmtpAttributesSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to **string** | Username and password colon separated for target SMTP authentication | [optional] 
**Port** | Pointer to **int32** | Target port | [optional] 
**Encryption** | Pointer to **bool** | Connection encryption | [optional] 
**Stringtoexpect** | Pointer to **string** | String to expect in response | [optional] 

## Methods

### NewSmtpAttributesSet

`func NewSmtpAttributesSet() *SmtpAttributesSet`

NewSmtpAttributesSet instantiates a new SmtpAttributesSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpAttributesSetWithDefaults

`func NewSmtpAttributesSetWithDefaults() *SmtpAttributesSet`

NewSmtpAttributesSetWithDefaults instantiates a new SmtpAttributesSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *SmtpAttributesSet) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *SmtpAttributesSet) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *SmtpAttributesSet) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *SmtpAttributesSet) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetPort

`func (o *SmtpAttributesSet) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SmtpAttributesSet) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SmtpAttributesSet) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SmtpAttributesSet) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncryption

`func (o *SmtpAttributesSet) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *SmtpAttributesSet) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *SmtpAttributesSet) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *SmtpAttributesSet) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetStringtoexpect

`func (o *SmtpAttributesSet) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *SmtpAttributesSet) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *SmtpAttributesSet) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.

### HasStringtoexpect

`func (o *SmtpAttributesSet) HasStringtoexpect() bool`

HasStringtoexpect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


