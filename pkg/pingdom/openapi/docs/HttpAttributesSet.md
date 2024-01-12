# HttpAttributesSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to **string** | Username and password, colon separated. | [optional] 
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

### NewHttpAttributesSet

`func NewHttpAttributesSet() *HttpAttributesSet`

NewHttpAttributesSet instantiates a new HttpAttributesSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpAttributesSetWithDefaults

`func NewHttpAttributesSetWithDefaults() *HttpAttributesSet`

NewHttpAttributesSetWithDefaults instantiates a new HttpAttributesSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *HttpAttributesSet) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *HttpAttributesSet) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *HttpAttributesSet) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *HttpAttributesSet) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetUrl

`func (o *HttpAttributesSet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HttpAttributesSet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HttpAttributesSet) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HttpAttributesSet) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEncryption

`func (o *HttpAttributesSet) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *HttpAttributesSet) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *HttpAttributesSet) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *HttpAttributesSet) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetPort

`func (o *HttpAttributesSet) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HttpAttributesSet) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HttpAttributesSet) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HttpAttributesSet) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetShouldcontain

`func (o *HttpAttributesSet) GetShouldcontain() string`

GetShouldcontain returns the Shouldcontain field if non-nil, zero value otherwise.

### GetShouldcontainOk

`func (o *HttpAttributesSet) GetShouldcontainOk() (*string, bool)`

GetShouldcontainOk returns a tuple with the Shouldcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldcontain

`func (o *HttpAttributesSet) SetShouldcontain(v string)`

SetShouldcontain sets Shouldcontain field to given value.

### HasShouldcontain

`func (o *HttpAttributesSet) HasShouldcontain() bool`

HasShouldcontain returns a boolean if a field has been set.

### GetShouldnotcontain

`func (o *HttpAttributesSet) GetShouldnotcontain() string`

GetShouldnotcontain returns the Shouldnotcontain field if non-nil, zero value otherwise.

### GetShouldnotcontainOk

`func (o *HttpAttributesSet) GetShouldnotcontainOk() (*string, bool)`

GetShouldnotcontainOk returns a tuple with the Shouldnotcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldnotcontain

`func (o *HttpAttributesSet) SetShouldnotcontain(v string)`

SetShouldnotcontain sets Shouldnotcontain field to given value.

### HasShouldnotcontain

`func (o *HttpAttributesSet) HasShouldnotcontain() bool`

HasShouldnotcontain returns a boolean if a field has been set.

### GetPostdata

`func (o *HttpAttributesSet) GetPostdata() string`

GetPostdata returns the Postdata field if non-nil, zero value otherwise.

### GetPostdataOk

`func (o *HttpAttributesSet) GetPostdataOk() (*string, bool)`

GetPostdataOk returns a tuple with the Postdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostdata

`func (o *HttpAttributesSet) SetPostdata(v string)`

SetPostdata sets Postdata field to given value.

### HasPostdata

`func (o *HttpAttributesSet) HasPostdata() bool`

HasPostdata returns a boolean if a field has been set.

### GetRequestheaders

`func (o *HttpAttributesSet) GetRequestheaders() []string`

GetRequestheaders returns the Requestheaders field if non-nil, zero value otherwise.

### GetRequestheadersOk

`func (o *HttpAttributesSet) GetRequestheadersOk() (*[]string, bool)`

GetRequestheadersOk returns a tuple with the Requestheaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestheaders

`func (o *HttpAttributesSet) SetRequestheaders(v []string)`

SetRequestheaders sets Requestheaders field to given value.

### HasRequestheaders

`func (o *HttpAttributesSet) HasRequestheaders() bool`

HasRequestheaders returns a boolean if a field has been set.

### GetVerifyCertificate

`func (o *HttpAttributesSet) GetVerifyCertificate() bool`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *HttpAttributesSet) GetVerifyCertificateOk() (*bool, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *HttpAttributesSet) SetVerifyCertificate(v bool)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *HttpAttributesSet) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

### GetSslDownDaysBefore

`func (o *HttpAttributesSet) GetSslDownDaysBefore() int32`

GetSslDownDaysBefore returns the SslDownDaysBefore field if non-nil, zero value otherwise.

### GetSslDownDaysBeforeOk

`func (o *HttpAttributesSet) GetSslDownDaysBeforeOk() (*int32, bool)`

GetSslDownDaysBeforeOk returns a tuple with the SslDownDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDownDaysBefore

`func (o *HttpAttributesSet) SetSslDownDaysBefore(v int32)`

SetSslDownDaysBefore sets SslDownDaysBefore field to given value.

### HasSslDownDaysBefore

`func (o *HttpAttributesSet) HasSslDownDaysBefore() bool`

HasSslDownDaysBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


