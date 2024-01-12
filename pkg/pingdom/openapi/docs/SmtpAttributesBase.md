# SmtpAttributesBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | Target port | [optional] 
**Encryption** | Pointer to **bool** | Connection encryption | [optional] 
**Stringtoexpect** | Pointer to **string** | String to expect in response | [optional] 

## Methods

### NewSmtpAttributesBase

`func NewSmtpAttributesBase() *SmtpAttributesBase`

NewSmtpAttributesBase instantiates a new SmtpAttributesBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpAttributesBaseWithDefaults

`func NewSmtpAttributesBaseWithDefaults() *SmtpAttributesBase`

NewSmtpAttributesBaseWithDefaults instantiates a new SmtpAttributesBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *SmtpAttributesBase) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SmtpAttributesBase) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SmtpAttributesBase) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SmtpAttributesBase) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncryption

`func (o *SmtpAttributesBase) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *SmtpAttributesBase) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *SmtpAttributesBase) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *SmtpAttributesBase) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetStringtoexpect

`func (o *SmtpAttributesBase) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *SmtpAttributesBase) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *SmtpAttributesBase) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.

### HasStringtoexpect

`func (o *SmtpAttributesBase) HasStringtoexpect() bool`

HasStringtoexpect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


