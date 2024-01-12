# HttpAttributesGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username for target HTTP authentication | [optional] 
**Password** | Pointer to **string** | Password for target HTTP authentication | [optional] 
**Url** | Pointer to **string** | Path to target on server | [optional] 
**Encryption** | Pointer to **bool** | Connection encryption | [optional] 
**Port** | Pointer to **int32** | Target port | [optional] 
**Shouldcontain** | Pointer to **string** | Target site should contain this string. Note! This parameter cannot be used together with the parameter “shouldnotcontain”, use only one of them in your request. | [optional] 
**Shouldnotcontain** | Pointer to **string** | Target site should NOT contain this string. Note! This parameter cannot be used together with the parameter “shouldcontain”, use only one of them in your request. | [optional] 
**Postdata** | Pointer to **string** | Data that should be posted to the web page, for example submission data for a sign-up or login form. The data needs to be formatted in the same way as a web browser would send it to the web server | [optional] 
**Requestheaders** | Pointer to **[]string** | Custom HTTP header. The entry value should contain a one-element string array. The element should contain &#x60;headerName&#x60; and &#x60;headerValue&#x60; colon-separated. To add more than one header send other parameters named &#x60;requestheaders{number}&#x60;. | [optional] 
**VerifyCertificate** | Pointer to **bool** | Treat target site as down if an invalid/unverifiable certificate is found. | [optional] [default to true]
**SslDownDaysBefore** | Pointer to **int32** | Treat the target site as down if a certificate expires within the given number of days. This parameter will be ignored if &#x60;verify_certificate&#x60; is set to &#x60;false&#x60;. It will appear provided &#x60;verify_certificate&#x60; is true and &#x60;ssl_down_days_before&#x60; value is greater than or equals 1. | [optional] [default to 0]

## Methods

### NewHttpAttributesGet

`func NewHttpAttributesGet() *HttpAttributesGet`

NewHttpAttributesGet instantiates a new HttpAttributesGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpAttributesGetWithDefaults

`func NewHttpAttributesGetWithDefaults() *HttpAttributesGet`

NewHttpAttributesGetWithDefaults instantiates a new HttpAttributesGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *HttpAttributesGet) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *HttpAttributesGet) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *HttpAttributesGet) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *HttpAttributesGet) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *HttpAttributesGet) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HttpAttributesGet) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HttpAttributesGet) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *HttpAttributesGet) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrl

`func (o *HttpAttributesGet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpAttributesGet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpAttributesGet) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpAttributesGet) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEncryption

`func (o *HttpAttributesGet) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *HttpAttributesGet) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *HttpAttributesGet) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *HttpAttributesGet) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetPort

`func (o *HttpAttributesGet) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HttpAttributesGet) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HttpAttributesGet) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HttpAttributesGet) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetShouldcontain

`func (o *HttpAttributesGet) GetShouldcontain() string`

GetShouldcontain returns the Shouldcontain field if non-nil, zero value otherwise.

### GetShouldcontainOk

`func (o *HttpAttributesGet) GetShouldcontainOk() (*string, bool)`

GetShouldcontainOk returns a tuple with the Shouldcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldcontain

`func (o *HttpAttributesGet) SetShouldcontain(v string)`

SetShouldcontain sets Shouldcontain field to given value.

### HasShouldcontain

`func (o *HttpAttributesGet) HasShouldcontain() bool`

HasShouldcontain returns a boolean if a field has been set.

### GetShouldnotcontain

`func (o *HttpAttributesGet) GetShouldnotcontain() string`

GetShouldnotcontain returns the Shouldnotcontain field if non-nil, zero value otherwise.

### GetShouldnotcontainOk

`func (o *HttpAttributesGet) GetShouldnotcontainOk() (*string, bool)`

GetShouldnotcontainOk returns a tuple with the Shouldnotcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldnotcontain

`func (o *HttpAttributesGet) SetShouldnotcontain(v string)`

SetShouldnotcontain sets Shouldnotcontain field to given value.

### HasShouldnotcontain

`func (o *HttpAttributesGet) HasShouldnotcontain() bool`

HasShouldnotcontain returns a boolean if a field has been set.

### GetPostdata

`func (o *HttpAttributesGet) GetPostdata() string`

GetPostdata returns the Postdata field if non-nil, zero value otherwise.

### GetPostdataOk

`func (o *HttpAttributesGet) GetPostdataOk() (*string, bool)`

GetPostdataOk returns a tuple with the Postdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostdata

`func (o *HttpAttributesGet) SetPostdata(v string)`

SetPostdata sets Postdata field to given value.

### HasPostdata

`func (o *HttpAttributesGet) HasPostdata() bool`

HasPostdata returns a boolean if a field has been set.

### GetRequestheaders

`func (o *HttpAttributesGet) GetRequestheaders() []string`

GetRequestheaders returns the Requestheaders field if non-nil, zero value otherwise.

### GetRequestheadersOk

`func (o *HttpAttributesGet) GetRequestheadersOk() (*[]string, bool)`

GetRequestheadersOk returns a tuple with the Requestheaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestheaders

`func (o *HttpAttributesGet) SetRequestheaders(v []string)`

SetRequestheaders sets Requestheaders field to given value.

### HasRequestheaders

`func (o *HttpAttributesGet) HasRequestheaders() bool`

HasRequestheaders returns a boolean if a field has been set.

### GetVerifyCertificate

`func (o *HttpAttributesGet) GetVerifyCertificate() bool`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *HttpAttributesGet) GetVerifyCertificateOk() (*bool, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *HttpAttributesGet) SetVerifyCertificate(v bool)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *HttpAttributesGet) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

### GetSslDownDaysBefore

`func (o *HttpAttributesGet) GetSslDownDaysBefore() int32`

GetSslDownDaysBefore returns the SslDownDaysBefore field if non-nil, zero value otherwise.

### GetSslDownDaysBeforeOk

`func (o *HttpAttributesGet) GetSslDownDaysBeforeOk() (*int32, bool)`

GetSslDownDaysBeforeOk returns a tuple with the SslDownDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDownDaysBefore

`func (o *HttpAttributesGet) SetSslDownDaysBefore(v int32)`

SetSslDownDaysBefore sets SslDownDaysBefore field to given value.

### HasSslDownDaysBefore

`func (o *HttpAttributesGet) HasSslDownDaysBefore() bool`

HasSslDownDaysBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


