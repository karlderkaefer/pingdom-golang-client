# HTTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Target host | 
**Type** | **string** |  | 
**Probeid** | Pointer to **int32** | Probe identifier | [optional] 
**ProbeFilters** | Pointer to **int32** | Filters used for probe selections. Comma separated key:value pairs. Currently only region is supported. Possible values are &#39;EU&#39;, &#39;NA&#39;, &#39;APAC&#39; and &#39;LATAM&#39;. | [optional] 
**Ipv6** | Pointer to **bool** | Use ipv6 instead of ipv4 | [optional] 
**ResponsetimeThreshold** | Pointer to **int32** | Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.) | [optional] [default to 30000]
**Url** | Pointer to **string** | (http specific) Target path on server | [optional] [default to "/"]
**Encryption** | Pointer to **bool** | (http specific) Connection encryption | [optional] [default to false]
**Port** | Pointer to **int32** | (http specific) Target port | [optional] [default to 80]
**Auth** | Pointer to **string** | (http specific) Username and password for target HTTP authentication. | [optional] 
**Shouldcontain** | Pointer to **string** | (http specific) Target site should contain this string | [optional] 
**Shouldnotcontain** | Pointer to **string** | (http specific) Target site should NOT contain this string | [optional] 
**Postdata** | Pointer to **string** | (http specific) Data that should be posted to the web page, for example submission data for a sign-up or login form. The data needs to be formatted in the same way as a web browser would send it to the web server | [optional] 
**RequestheaderX** | Pointer to **string** | (http specific) Custom HTTP header name. Replace {X} with a number unique for each header argument. | [optional] 
**VerifyCertificate** | Pointer to **bool** | (http specific) Treat target site as down if an invalid/unverifiable certificate is found. | [optional] [default to true]
**SslDownDaysBefore** | Pointer to **int32** | (http specific) Treat the target site as down if a certificate expires within the given number of days. This parameter will be ignored if &#x60;verify_certificate&#x60; is set to &#x60;false&#x60;. | [optional] [default to 0]

## Methods

### NewHTTP

`func NewHTTP(host string, type_ string, ) *HTTP`

NewHTTP instantiates a new HTTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPWithDefaults

`func NewHTTPWithDefaults() *HTTP`

NewHTTPWithDefaults instantiates a new HTTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *HTTP) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HTTP) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HTTP) SetHost(v string)`

SetHost sets Host field to given value.


### GetType

`func (o *HTTP) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HTTP) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HTTP) SetType(v string)`

SetType sets Type field to given value.


### GetProbeid

`func (o *HTTP) GetProbeid() int32`

GetProbeid returns the Probeid field if non-nil, zero value otherwise.

### GetProbeidOk

`func (o *HTTP) GetProbeidOk() (*int32, bool)`

GetProbeidOk returns a tuple with the Probeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeid

`func (o *HTTP) SetProbeid(v int32)`

SetProbeid sets Probeid field to given value.

### HasProbeid

`func (o *HTTP) HasProbeid() bool`

HasProbeid returns a boolean if a field has been set.

### GetProbeFilters

`func (o *HTTP) GetProbeFilters() int32`

GetProbeFilters returns the ProbeFilters field if non-nil, zero value otherwise.

### GetProbeFiltersOk

`func (o *HTTP) GetProbeFiltersOk() (*int32, bool)`

GetProbeFiltersOk returns a tuple with the ProbeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeFilters

`func (o *HTTP) SetProbeFilters(v int32)`

SetProbeFilters sets ProbeFilters field to given value.

### HasProbeFilters

`func (o *HTTP) HasProbeFilters() bool`

HasProbeFilters returns a boolean if a field has been set.

### GetIpv6

`func (o *HTTP) GetIpv6() bool`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *HTTP) GetIpv6Ok() (*bool, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *HTTP) SetIpv6(v bool)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *HTTP) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetResponsetimeThreshold

`func (o *HTTP) GetResponsetimeThreshold() int32`

GetResponsetimeThreshold returns the ResponsetimeThreshold field if non-nil, zero value otherwise.

### GetResponsetimeThresholdOk

`func (o *HTTP) GetResponsetimeThresholdOk() (*int32, bool)`

GetResponsetimeThresholdOk returns a tuple with the ResponsetimeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsetimeThreshold

`func (o *HTTP) SetResponsetimeThreshold(v int32)`

SetResponsetimeThreshold sets ResponsetimeThreshold field to given value.

### HasResponsetimeThreshold

`func (o *HTTP) HasResponsetimeThreshold() bool`

HasResponsetimeThreshold returns a boolean if a field has been set.

### GetUrl

`func (o *HTTP) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HTTP) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HTTP) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *HTTP) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEncryption

`func (o *HTTP) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *HTTP) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *HTTP) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *HTTP) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetPort

`func (o *HTTP) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HTTP) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HTTP) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HTTP) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAuth

`func (o *HTTP) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *HTTP) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *HTTP) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *HTTP) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetShouldcontain

`func (o *HTTP) GetShouldcontain() string`

GetShouldcontain returns the Shouldcontain field if non-nil, zero value otherwise.

### GetShouldcontainOk

`func (o *HTTP) GetShouldcontainOk() (*string, bool)`

GetShouldcontainOk returns a tuple with the Shouldcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldcontain

`func (o *HTTP) SetShouldcontain(v string)`

SetShouldcontain sets Shouldcontain field to given value.

### HasShouldcontain

`func (o *HTTP) HasShouldcontain() bool`

HasShouldcontain returns a boolean if a field has been set.

### GetShouldnotcontain

`func (o *HTTP) GetShouldnotcontain() string`

GetShouldnotcontain returns the Shouldnotcontain field if non-nil, zero value otherwise.

### GetShouldnotcontainOk

`func (o *HTTP) GetShouldnotcontainOk() (*string, bool)`

GetShouldnotcontainOk returns a tuple with the Shouldnotcontain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldnotcontain

`func (o *HTTP) SetShouldnotcontain(v string)`

SetShouldnotcontain sets Shouldnotcontain field to given value.

### HasShouldnotcontain

`func (o *HTTP) HasShouldnotcontain() bool`

HasShouldnotcontain returns a boolean if a field has been set.

### GetPostdata

`func (o *HTTP) GetPostdata() string`

GetPostdata returns the Postdata field if non-nil, zero value otherwise.

### GetPostdataOk

`func (o *HTTP) GetPostdataOk() (*string, bool)`

GetPostdataOk returns a tuple with the Postdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostdata

`func (o *HTTP) SetPostdata(v string)`

SetPostdata sets Postdata field to given value.

### HasPostdata

`func (o *HTTP) HasPostdata() bool`

HasPostdata returns a boolean if a field has been set.

### GetRequestheaderX

`func (o *HTTP) GetRequestheaderX() string`

GetRequestheaderX returns the RequestheaderX field if non-nil, zero value otherwise.

### GetRequestheaderXOk

`func (o *HTTP) GetRequestheaderXOk() (*string, bool)`

GetRequestheaderXOk returns a tuple with the RequestheaderX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestheaderX

`func (o *HTTP) SetRequestheaderX(v string)`

SetRequestheaderX sets RequestheaderX field to given value.

### HasRequestheaderX

`func (o *HTTP) HasRequestheaderX() bool`

HasRequestheaderX returns a boolean if a field has been set.

### GetVerifyCertificate

`func (o *HTTP) GetVerifyCertificate() bool`

GetVerifyCertificate returns the VerifyCertificate field if non-nil, zero value otherwise.

### GetVerifyCertificateOk

`func (o *HTTP) GetVerifyCertificateOk() (*bool, bool)`

GetVerifyCertificateOk returns a tuple with the VerifyCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyCertificate

`func (o *HTTP) SetVerifyCertificate(v bool)`

SetVerifyCertificate sets VerifyCertificate field to given value.

### HasVerifyCertificate

`func (o *HTTP) HasVerifyCertificate() bool`

HasVerifyCertificate returns a boolean if a field has been set.

### GetSslDownDaysBefore

`func (o *HTTP) GetSslDownDaysBefore() int32`

GetSslDownDaysBefore returns the SslDownDaysBefore field if non-nil, zero value otherwise.

### GetSslDownDaysBeforeOk

`func (o *HTTP) GetSslDownDaysBeforeOk() (*int32, bool)`

GetSslDownDaysBeforeOk returns a tuple with the SslDownDaysBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslDownDaysBefore

`func (o *HTTP) SetSslDownDaysBefore(v int32)`

SetSslDownDaysBefore sets SslDownDaysBefore field to given value.

### HasSslDownDaysBefore

`func (o *HTTP) HasSslDownDaysBefore() bool`

HasSslDownDaysBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


