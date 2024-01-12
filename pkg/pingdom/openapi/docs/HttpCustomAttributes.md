# HttpCustomAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Path to target on server | 
**Encryption** | Pointer to **bool** | Connection encryption | [optional] 
**Port** | Pointer to **int32** | Target port | [optional] 
**Additionalurls** | Pointer to **string** | Full URL (including hostname) to target additional XML file | [optional] 
**VerifyCertificate** | Pointer to **bool** | Treat target site as down if an invalid/unverifiable certificate is found. | [optional] [default to true]
**SslDownDaysBefore** | Pointer to **int32** | Treat the target site as down if a certificate expires within the given number of days. This parameter will be ignored if &#x60;verify_certificate&#x60; is set to &#x60;false&#x60;. | [optional] [default to 0]

## Methods

### NewHttpCustomAttributes

`func NewHttpCustomAttributes(url string, ) *HttpCustomAttributes`

NewHttpCustomAttributes instantiates a new HttpCustomAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpCustomAttributesWithDefaults

`func NewHttpCustomAttributesWithDefaults() *HttpCustomAttributes`

NewHttpCustomAttributesWithDefaults instantiates a new HttpCustomAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *HttpCustomAttributes) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpCustomAttributes) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpCustomAttributes) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEncryption

`func (o *HttpCustomAttributes) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *HttpCustomAttributes) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *HttpCustomAttributes) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *HttpCustomAttributes) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetPort

`func (o *HttpCustomAttributes) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HttpCustomAttributes) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HttpCustomAttributes) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HttpCustomAttributes) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAdditionalurls

`func (o *HttpCustomAttributes) GetAdditionalurls() string`

GetAdditionalurls returns the Additionalurls field if non-nil, zero value otherwise.

### GetAdditionalurlsOk

`func (o *HttpCustomAttributes) GetAdditionalurlsOk() (*string, bool)`

GetAdditionalurlsOk returns a tuple with the Additionalurls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalurls

`func (o *HttpCustomAttributes) SetAdditionalurls(v string)`

SetAdditionalurls sets Additionalurls field to given value.

### HasAdditionalurls

`func (o *HttpCustomAttributes) HasAdditionalurls() bool`

HasAdditionalurls returns a boolean if a field has been set.

### GetVerifyCertificate

`func (o *HttpCustomAttributes) GetVerifyCertificate() bool`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *HttpCustomAttributes) GetVerifyCertificateOk() (*bool, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *HttpCustomAttributes) SetVerifyCertificate(v bool)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *HttpCustomAttributes) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

### GetSslDownDaysBefore

`func (o *HttpCustomAttributes) GetSslDownDaysBefore() int32`

GetSslDownDaysBefore returns the SslDownDaysBefore field if non-nil, zero value otherwise.

### GetSslDownDaysBeforeOk

`func (o *HttpCustomAttributes) GetSslDownDaysBeforeOk() (*int32, bool)`

GetSslDownDaysBeforeOk returns a tuple with the SslDownDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDownDaysBefore

`func (o *HttpCustomAttributes) SetSslDownDaysBefore(v int32)`

SetSslDownDaysBefore sets SslDownDaysBefore field to given value.

### HasSslDownDaysBefore

`func (o *HttpCustomAttributes) HasSslDownDaysBefore() bool`

HasSslDownDaysBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


