# SmtpAttributesGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username for target SMTP authentication | [optional] 
**Password** | Pointer to **string** | Password for target SMTP authentication | [optional] 
**Port** | Pointer to **int32** | Target port | [optional] 
**Encryption** | Pointer to **bool** | Connection encryption | [optional] 
**Stringtoexpect** | Pointer to **string** | String to expect in response | [optional] 

## Methods

### NewSmtpAttributesGet

`func NewSmtpAttributesGet() *SmtpAttributesGet`

NewSmtpAttributesGet instantiates a new SmtpAttributesGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpAttributesGetWithDefaults

`func NewSmtpAttributesGetWithDefaults() *SmtpAttributesGet`

NewSmtpAttributesGetWithDefaults instantiates a new SmtpAttributesGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *SmtpAttributesGet) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SmtpAttributesGet) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SmtpAttributesGet) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SmtpAttributesGet) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *SmtpAttributesGet) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SmtpAttributesGet) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SmtpAttributesGet) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SmtpAttributesGet) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *SmtpAttributesGet) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SmtpAttributesGet) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SmtpAttributesGet) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SmtpAttributesGet) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncryption

`func (o *SmtpAttributesGet) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *SmtpAttributesGet) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *SmtpAttributesGet) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *SmtpAttributesGet) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetStringtoexpect

`func (o *SmtpAttributesGet) GetStringtoexpect() string`

GetStringtoexpect returns the Stringtoexpect field if non-nil, zero value otherwise.

### GetStringtoexpectOk

`func (o *SmtpAttributesGet) GetStringtoexpectOk() (*string, bool)`

GetStringtoexpectOk returns a tuple with the Stringtoexpect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringtoexpect

`func (o *SmtpAttributesGet) SetStringtoexpect(v string)`

SetStringtoexpect sets Stringtoexpect field to given value.

### HasStringtoexpect

`func (o *SmtpAttributesGet) HasStringtoexpect() bool`

HasStringtoexpect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


